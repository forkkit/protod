package pstore

import (
	"context"
	"math/rand"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/dave/protod/delta"
	"github.com/dave/protod/perr"
	"github.com/dave/protod/pserver"
	"google.golang.org/protobuf/proto"
)

type StateId string
type DocumentId string

// During an edit request we transform the operations in order that client and server both converge on the same value:
//
//       A (server & client) -> o
//                             / \
//                            /   \
//      op2 (from client) -> /     \ <- op1 (retrieved from the database)
//                          /       \
//                         /         \
//          C (client) -> o           o <- B (server)
//                         \         /
//                          \       /
//  op1x (send to client) -> \     / <- op2x (stored in the database)
//                            \   /
//                             \ /
//                              o <- D (server & client)
//
// The client and server are originally at [A]. The server has performed [op1] to arrive at [B], at the same time
// as the client was performing [op2] to arrive at [C]. We calculate [op1x] and [op2x] which both converge on [D].
// [op1x] is sent to the client, and [op2x] is stored in the database.
//
// Function parameters:
// stateId:      The unique id of this state record. This is generated by the client and resubmitted on retries.
// documentType: The fully qualified protobuf type name.
// documentId:   The unique id of the document.
// state:        The document state that client operations were performed against. Use 0 along with a set operation at the document root to add a new document.
// op2:          The client operation that was performed.
func Edit(ctx context.Context, s *pserver.Server, documentType string, documentId DocumentId, stateId StateId, state int64, op2 *delta.Op) (int64, *delta.Op, error) {

	t := s.Type(documentType)
	if t == nil {
		return 0, nil, perr.Debugf("type %q not found", documentType)
	}

	ref := s.Firestore.Collection(t.CollectionName()).Doc(string(documentId))
	stateRef := ref.Collection(pserver.STATES_COLLECTION).Doc(string(stateId))

	if state == 0 {
		ops := op2.Flatten()
		if len(ops) == 0 || ops[0].Type != delta.Op_Set || ops[0].Location != nil {
			return 0, nil, perr.Debug("if state is 0, first op must be a single set operation with nil location")
		}
	}

	if op2 == nil {
		// Just request an update (no need to store state)
		var ops []*delta.Op
		newState, err := s.Changes(ctx, nil, t, ref, state, 0, func(op *delta.Op) error {
			ops = append(ops, op)
			return nil
		})
		if err != nil {
			return 0, nil, perr.Wrap(err).Debug("changes")
		}
		return newState, delta.Compound(ops...), nil
	}

	var lastState, newState int64
	var op1x, op2x *delta.Op
	var duplicate *firestore.DocumentSnapshot

	f := func(ctx context.Context, tx *firestore.Transaction) error {

		// check for documents with the same unique request id
		documentExists, err := s.DocumentExists(ctx, tx, ref)
		if err != nil {
			return perr.Wrap(err).Debug("getting existing document")
		}
		if !documentExists {

			if state != 0 {
				return perr.Debug("document didn't exist but edit operation supplied non zero state")
			}

			if t.OnAdd != nil {
				if err := t.OnAdd(ctx, s, tx, string(documentId)); err != nil {
					return perr.Wrap(err).Debug("add handler")
				}
			}

			if newState, err = createDocument(ctx, s, tx, t, ref, stateRef, op2); err != nil {
				return perr.Wrap(err).Debug("create document")
			}

			return nil
		}

		duplicate, err = s.QueryState(ctx, tx, stateRef)
		if err != nil {
			return perr.Wrap(err).Debug("querying state")
		}
		if duplicate != nil {
			// Exit from transaction and continue processing outside (no writes needed)
			return nil
		}

		if state == 0 {

			// For zero state, we don't need to transform because the operation will overwrite the entire document.

			lastState, err := s.Latest(ctx, tx, t, ref)
			if err != nil {
				return perr.Wrap(err).Debug("getting latest state")
			}

			newState = lastState + 1

			opBlob := t.NewBlob()
			if err := opBlob.MarshalBlob(ctx, op2); err != nil {
				return perr.Wrap(err).Debug("unmarshaling op blob")
			}

			stateItem := t.NewState()
			if err := stateItem.PackState(ctx, newState, opBlob); err != nil {
				return perr.Wrap(err).Debug("packing state")
			}

			if err := tx.Create(stateRef, stateItem); err != nil {
				return perr.Wrap(err).Debug("firestore create").Flag(pserver.Firestore)
			}

			return nil
		}

		lastState, op1x, op2x, err = s.Transform(ctx, tx, t, ref, op2, state, 0)
		if err != nil {
			return perr.Wrap(err).Debug("transform")
		}

		// custom code
		if t.OnEdit != nil {
			if err := t.OnEdit(ctx, s, tx, string(documentId)); err != nil {
				return perr.Wrap(err).Debug("edit handler")
			}
		}

		// Increment the state
		newState = lastState + 1

		// Store op2x in the database
		op2xBlob := t.NewBlob()
		if err := op2xBlob.MarshalBlob(ctx, op2x); err != nil {
			return perr.Wrap(err).Debug("marshaling op2x blob")
		}

		newStateItem := t.NewState()
		if err := newStateItem.PackState(ctx, newState, op2xBlob); err != nil {
			return perr.Wrap(err).Debug("packing state")
		}

		if err := tx.Create(stateRef, newStateItem); err != nil {
			return perr.Wrap(err).Debug("firestore create").Flag(pserver.Firestore)
		}

		return nil

	}
	tf := func() error {
		if err := s.Firestore.RunTransaction(s.FirestoreContext(ctx), f); err != nil {
			return perr.Wrap(err).Debug("running transaction").Flag(pserver.Firestore)
		}
		return nil
	}
	if err := pserver.Lock(ctx, ref.Path, tf); err != nil {
		return 0, nil, perr.Wrap(err).Debug("locking")
	}

	if duplicate == nil {
		return newState, op1x, nil
	}

	// This request has already been processed. We can recreate the correct response, and nothing needs to be
	// stored in the database.
	duplicateState, _, err := s.UnpackState(ctx, duplicate, t)
	if err != nil {
		return 0, nil, perr.Wrap(err).Debug("unpacking duplicate state")
	}

	// Note that the state provided by the client is the "before" state for the transform, and the state from the
	// duplicate record is the "after" state in the transform:
	_, op1x, _, err = s.Transform(ctx, nil, t, ref, op2, state, duplicateState)
	if err != nil {
		return 0, nil, perr.Wrap(err).Debug("transforming duplicate state")
	}

	return duplicateState, op1x, nil

}

func Add(ctx context.Context, s *pserver.Server, documentType string, documentId DocumentId, stateId StateId, document proto.Message) (int64, error) {
	state, _, err := Edit(ctx, s, documentType, documentId, stateId, 0, delta.Root(document))
	if err != nil {
		return 0, perr.Wrap(err).Debug("editing")
	}
	return state, nil
}

func createDocument(ctx context.Context, s *pserver.Server, tx *firestore.Transaction, t *pserver.DocumentType, ref, stateRef *firestore.DocumentRef, op2 *delta.Op) (int64, error) {

	if stateRef == nil {
		stateRef = ref.Collection(pserver.STATES_COLLECTION).Doc(string(NewStateID()))
	}

	document := t.NewDocument()
	documentBlob := t.NewBlob()
	if err := documentBlob.MarshalBlob(ctx, document); err != nil {
		return 0, perr.Wrap(err).Debug("marchaling document blob")
	}
	snapshotItem := t.NewSnapshot()
	if err := snapshotItem.PackSnapshot(ctx, 0, documentBlob, nil, document); err != nil {
		return 0, perr.Wrap(err).Debug("packing snapshot")
	}
	if err := tx.Create(ref, snapshotItem); err != nil {
		return 0, perr.Wrap(err).Debug("firestore create snapshot").Flag(pserver.Firestore)
	}

	opBlob := t.NewBlob()
	if err := opBlob.MarshalBlob(ctx, op2); err != nil {
		return 0, perr.Wrap(err).Debug("marshaling op blob")
	}
	newState := int64(1)
	stateItem := t.NewState()
	if err := stateItem.PackState(ctx, newState, opBlob); err != nil {
		return 0, perr.Wrap(err).Debug("packing state")
	}

	if err := tx.Create(stateRef, stateItem); err != nil {
		return 0, perr.Wrap(err).Debug("firestore create state").Flag(pserver.Firestore)
	}

	return newState, nil
}

func Get(ctx context.Context, s *pserver.Server, documentType string, documentId DocumentId, create bool) (int64, proto.Message, error) {

	t := s.Type(documentType)
	if t == nil {
		return 0, nil, perr.Debugf("type %q not found", documentType)
	}

	// custom code
	if t.OnGet != nil {
		if err := t.OnGet(ctx, s, string(documentId)); err != nil {
			return 0, nil, perr.Wrap(err).Debug("get handler")
		}
	}

	ref := s.Firestore.Collection(t.CollectionName()).Doc(string(documentId))

	var document proto.Message
	var added bool
	var state int64
	var snapshotState int64
	if !create {

		// if not create on missing, no need for transaction:
		var err error
		if snapshotState, document, _, err = s.UnpackSnapshot(ctx, nil, t, ref); err != nil {
			return 0, nil, perr.Wrap(err).Debug("unpacking snapshot")
		}

	} else {

		f := func(ctx context.Context, tx *firestore.Transaction) error {

			var err error
			snapshotState, document, _, err = s.UnpackSnapshot(ctx, tx, t, ref)
			switch {
			case perr.Any(err, pserver.IsNotFound):
				added = true
				document = t.NewDocument()
				state, err = createDocument(ctx, s, tx, t, ref, nil, delta.Root(document))
				if err != nil {
					return perr.Wrap(err).Debug("creating document")
				}
			case err != nil:
				return perr.Wrap(err).Debug("unpacking snapshot")
			}

			return nil
		}
		tf := func() error {
			if err := s.Firestore.RunTransaction(s.FirestoreContext(ctx), f); err != nil {
				return perr.Wrap(err).Debug("running transaction").Flag(pserver.Firestore)
			}
			return nil
		}
		if err := pserver.Lock(ctx, ref.Path, tf); err != nil {
			return 0, nil, perr.Wrap(err).Debug("locking")
		}
	}

	if !added {
		// this does not need to be inside the transaction, and is not needed if we
		// added a new record.
		var err error
		state, err = s.Changes(ctx, nil, t, ref, snapshotState, 0, func(op *delta.Op) error {
			if err := delta.Apply(op, document); err != nil {
				return perr.Wrap(err).Debug("applying")
			}
			return nil
		})
		if err != nil {
			return 0, nil, perr.Wrap(err).Debug("running changes")
		}
	}

	return state, document, nil
}

func Refresh(ctx context.Context, s *pserver.Server, documentType string, documentId DocumentId) error {

	t := s.Type(documentType)
	if t == nil {
		return perr.Debugf("type %q not found", documentType)
	}

	ref := s.Firestore.Collection(t.CollectionName()).Doc(string(documentId))

	// Update the value snapshot. this doesn't need to be inside a transaction, because if the
	// snapshot is slightly out of date it doesn't matter.
	snapshotState, document, oldSnapshot, err := s.UnpackSnapshot(ctx, nil, t, ref)
	if err != nil {
		return perr.Wrap(err).Debug("unpacking snapshot")
	}

	state, err := s.Changes(ctx, nil, t, ref, snapshotState, 0, func(op *delta.Op) error {
		if err := delta.Apply(op, document); err != nil {
			return perr.Wrap(err).Debug("applying")
		}
		return nil
	})
	if err != nil {
		return perr.Wrap(err).Debug("running changes")
	}
	if state == snapshotState {
		return nil
	}
	documentBlob := t.NewBlob()
	if err := documentBlob.MarshalBlob(ctx, document); err != nil {
		return perr.Wrap(err).Debug("marshaling document blob")
	}

	newSnapshot := t.NewSnapshot()
	if err := newSnapshot.PackSnapshot(ctx, state, documentBlob, oldSnapshot, document); err != nil {
		return perr.Wrap(err).Debug("packing snapshot")
	}
	if _, err := ref.Set(s.FirestoreContext(ctx), newSnapshot); err != nil {
		return perr.Wrap(err).Debug("firestore set").Flag(pserver.Firestore)
	}
	return nil
}

func NewStateID() StateId {
	return StateId(uniqueID())
}

func NewDocumentID() DocumentId {
	return DocumentId(uniqueID())
}

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func uniqueID() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const length = 20
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
