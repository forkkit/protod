import 'package:protod/delta.dart';
import 'package:protod/delta.pb.dart' as pb;
import 'package:protod/delta_behaviour.dart';
import 'package:protod/delta_shifters.dart';
import 'package:protod/delta_transform_generated.dart';

pb.Op transform(pb.Op t, pb.Op op, bool priority) {
  if (op == null) {
    return null;
  }
  if (t == null) {
    return op.clone();
  }
  if (op.type == pb.Op_Type.Compound) {
    var transformed = List<pb.Op>();
    op.ops.forEach((o) {
      final ox = transform(t, o, priority);
      if (ox != null) {
        transformed.add(ox);
      }
    });
    switch (transformed.length) {
      case 0:
        return null;
      case 1:
        return transformed[0];
      default:
        return pb.Op()
          ..type = pb.Op_Type.Compound
          ..ops.addAll(transformed);
    }
  }
  if (t.type == pb.Op_Type.Compound) {
    var out = op.clone();
    for (final t1 in t.ops) {
      out = transform(t1, out, priority);
      if (out == null) {
        return null;
      }
    }
    return out;
  }
  final oneofLocation = splitCommonOneofAncestor(t.location, op.location);
  if (oneofLocation != null) {
    // t and op have a common oneof ancestor, and are acting on separate values. Any operation on the descendant of
    // a oneof value will delete the entire tree under all the other oneof values.
    var priorityOp = priority ? t : op;
    var notOp = priority ? op : t;
    final Map<pb.Op_Type, bool> valid = {
      pb.Op_Type.Set: true,
      pb.Op_Type.Edit: false,
      pb.Op_Type.Insert: true,
      pb.Op_Type.Move: false,
      pb.Op_Type.Rename: false,
      pb.Op_Type.Delete: false,
    };
    if (!valid[priorityOp.type] && valid[notOp.type]) {
      // if notOp is valid and priorityOp is not, we can swap them round so a set/insert always takes priority
      // over an edit/move/rename/delete.
      var tmp = priorityOp;
      priorityOp = notOp;
      notOp = tmp;
    }
    if (valid[priorityOp.type]) {
      // nuke everything, and re-run the priority operation (if it's a set / insert).
      return pb.Op()
        ..type = pb.Op_Type.Compound
        ..ops.add(pb.Op()
          ..type = pb.Op_Type.Delete
          ..location.addAll(oneofLocation))
        ..ops.add(priorityOp.clone());
    } else {
      // nuke everything.
      return pb.Op()
        ..type = pb.Op_Type.Delete
        ..location.addAll(oneofLocation);
    }
  }
  return transformGenerated(t, op, priority);
}

pb.Op tIndependent(pb.Op t, pb.Op op) {
  // op1 and op2 are not acting on the same value, or the values don't conflict.
  var behaviour = getBehaviour(t);
  var opBehaviour = getBehaviour(op);

  if (behaviour.itemIsDeleted &&
      treeRelationship(t.location, op.location) ==
          TreeRelationshipType.ANCESTOR) {
    // Op is acting on a value that is a descendent of a value that t deleted. We should delete op.
    return null;
  }
  if (behaviour.valueIsLocation &&
      behaviour.valueIsDeleted &&
      treeRelationship(toLoc(t), op.location) ==
          TreeRelationshipType.ANCESTOR) {
    // Op is acting on a value that is a descendent of a value that t deleted. We should delete op.
    return null;
  }

  if (behaviour.indexValueShifter != null &&
      treeRelationship(parent(t), op.location) ==
          TreeRelationshipType.ANCESTOR) {
    // Op is acting on a value that is a descendent of a value that may have had it's list index shifted by t.
    // We should update the list index of the locator using the index shifter function.
    final valueShifter = behaviour.indexValueShifter(t, op);
    var out = op.clone();
    setItemIndex(out, valueShifter(itemIndex(op)));
    if (opBehaviour.valueIsLocation &&
        treeRelationship(parent(t), parent(op)) == TreeRelationshipType.EQUAL) {
      // If t and op both act on values within the same list (have the same parent), AND op has a location at
      // it's value, then we update the location using the location shifter. Example is two moves which are
      // acting on different values within a list.
      final locationShifter = behaviour.indexLocationShifter(t, op);
      setToIndex(out, locationShifter(toIndex(op)));
    }
    return out;
  }

  if (behaviour.keyShifter != null &&
      treeRelationship(parent(t), op.location) ==
          TreeRelationshipType.ANCESTOR) {
    // Op is acting on a value that is a descendent of a value that may have had it's map key shifted by t. We
    // should update the map key of the locator using the index shifter function.
    final shifter = behaviour.keyShifter(t, op);
    final index = t.location.length - 1;
    final value = op.location[index].key;
    var out = op.clone();
    out.location[index] = pb.Locator()..key = shifter(value);
    // We don't need to worry about updating the value because all possible instances where the value key would
    // need updating are handled by special cases. e.g. conflicting map moves.
    return out;
  }

  return op.clone();
}

pb.Op tEdit(pb.Op t, pb.Op op, bool priority) {
  // Used by:
  // tEditFieldEditField
  // tEditIndexEditIndex
  // tEditKeyEditKey

  if (treeRelationship(t.location, op.location) != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  }
  // Two delta edits operating on the same value - use Quill library to transform the operation.
  final tQuill = quillFromDelta(t.delta);
  final opQuill = quillFromDelta(op.delta);
  var out = op.clone();
  out.delta = deltaFromQuill(tQuill.transform(opQuill, priority));
  return out;
}

pb.Op tEditFieldEditField(pb.Op t, pb.Op op, bool priority) {
  return tEdit(t, op, priority);
}

pb.Op tEditIndexEditIndex(pb.Op t, pb.Op op, bool priority) {
  return tEdit(t, op, priority);
}

pb.Op tEditKeyEditKey(pb.Op t, pb.Op op, bool priority) {
  return tEdit(t, op, priority);
}

pb.Op tEditOverwrite(pb.Op t, pb.Op op) {
  // Used by:
  // tEditFieldSetField
  // tEditIndexSetIndex
  // tEditKeySetKey
  // tEditFieldDeleteField
  // tEditIndexDeleteIndex
  // tEditKeyDeleteKey
  // tSetIndexDeleteIndex

  if (treeRelationship(t.location, op.location) != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  }
  // Op is trying to overwrite the value that t has modified. In order to converge, op will have priority.
  return op.clone();
}

pb.Op tOverwriteEdit(pb.Op t, pb.Op op) {
  // Used by:
  // tSetFieldEditField
  // tSetIndexEditIndex
  // tSetKeyEditKey
  // tDeleteFieldEditField
  // tDeleteIndexEditIndex
  // tDeleteKeyEditKey
  // tDeleteIndexSetIndex

  if (treeRelationship(t.location, op.location) != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  }
  // op is trying to edit a value that t has already overwritten. In order to converge, t must take priority.
  return null;
}

pb.Op tEditFieldSetField(pb.Op t, pb.Op op, bool priority) {
  return tEditOverwrite(t, op);
}

pb.Op tEditFieldDeleteField(pb.Op t, pb.Op op, bool priority) {
  return tEditOverwrite(t, op);
}

pb.Op tEditIndexSetIndex(pb.Op t, pb.Op op, bool priority) {
  return tEditOverwrite(t, op);
}

pb.Op tEditIndexDeleteIndex(pb.Op t, pb.Op op, bool priority) {
  return tEditOverwrite(t, op);
}

pb.Op tEditKeySetKey(pb.Op t, pb.Op op, bool priority) {
  return tEditOverwrite(t, op);
}

pb.Op tEditKeyDeleteKey(pb.Op t, pb.Op op, bool priority) {
  return tEditOverwrite(t, op);
}

pb.Op tSetIndexDeleteIndex(pb.Op t, pb.Op op, bool priority) {
  return tEditOverwrite(t, op);
}

pb.Op tSetFieldEditField(pb.Op t, pb.Op op, bool priority) {
  return tOverwriteEdit(t, op);
}

pb.Op tDeleteFieldEditField(pb.Op t, pb.Op op, bool priority) {
  return tOverwriteEdit(t, op);
}

pb.Op tSetIndexEditIndex(pb.Op t, pb.Op op, bool priority) {
  return tOverwriteEdit(t, op);
}

pb.Op tDeleteIndexEditIndex(pb.Op t, pb.Op op, bool priority) {
  return tOverwriteEdit(t, op);
}

pb.Op tSetKeyEditKey(pb.Op t, pb.Op op, bool priority) {
  return tOverwriteEdit(t, op);
}

pb.Op tDeleteKeyEditKey(pb.Op t, pb.Op op, bool priority) {
  return tOverwriteEdit(t, op);
}

pb.Op tDeleteIndexSetIndex(pb.Op t, pb.Op op, bool priority) {
  return tOverwriteEdit(t, op);
}

pb.Op tEditIndexInsertIndex(pb.Op t, pb.Op op, bool priority) {
  // Op is trying to insert a value into a list after t modified a value. Even when t and op act on the same
  // location, they are independent.
  return tIndependent(t, op);
}

pb.Op tInsertIndexEditIndex(pb.Op t, pb.Op op, bool priority) {
  // Op is editing at the same index that t has inserted. op will edit at the shifted index and operations will be
  // independent. This is handled by tIndependent.
  return tIndependent(t, op);
}

pb.Op tMoveModify(pb.Op t, pb.Op op) {
  // Shared by:
  // tMoveIndexEditIndex
  // tMoveIndexSetIndex

  if (isNullMove(t)) {
    return op.clone();
  }

  // op is trying to modify a value after t has moved values. Even when t and op act on the same location, they are
  // independent.
  return tIndependent(t, op);
}

pb.Op tModifyMove(pb.Op t, pb.Op op) {
  // Shared by:
  // tEditIndexMoveIndex
  // tSetIndexMoveIndex

  if (isNullMove(op)) {
    return null;
  }
  // Op is trying to move a value in a list after op modified a value. Even when t and op act on the same
  // location, they are independent.
  return tIndependent(t, op);
}

pb.Op tMoveIndexEditIndex(pb.Op t, pb.Op op, bool priority) {
  return tMoveModify(t, op);
}

pb.Op tMoveIndexSetIndex(pb.Op t, pb.Op op, bool priority) {
  return tMoveModify(t, op);
}

pb.Op tEditIndexMoveIndex(pb.Op t, pb.Op op, bool priority) {
  return tModifyMove(t, op);
}

pb.Op tSetIndexMoveIndex(pb.Op t, pb.Op op, bool priority) {
  return tModifyMove(t, op);
}

pb.Op tMoveIndexDeleteIndex(pb.Op t, pb.Op op, bool priority) {
  if (isNullMove(t)) {
    return op.clone();
  }
  // op is trying to delete a value after t has moved values. Even when t and op act on the same location,
  // they are independent.
  return tIndependent(t, op);
}

pb.Op tDeleteIndexMoveIndex(pb.Op t, pb.Op op, bool priority) {
  if (isNullMove(op)) {
    return null;
  }
  final from = treeRelationship(t.location, op.location);
  final to = treeRelationship(t.location, toLoc(op));
  if (from != TreeRelationshipType.EQUAL && to != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  } else if (from == TreeRelationshipType.EQUAL) {
    // op is trying to move the value that op has already deleted. In order to converge, t must take priority.
    return null;
  } else if (to == TreeRelationshipType.EQUAL) {
    // op is trying to move to the index of the value that op already deleted. Operations are independent.
    return tIndependent(t, op);
  } else {
    throw Exception("");
  }
}

pb.Op tEditKeyRenameKey(pb.Op t, pb.Op op, bool priority) {
  if (isNullMove(op)) {
    return null;
  }
  final from = treeRelationship(t.location, op.location);
  final to = treeRelationship(t.location, toLoc(op));
  if (from != TreeRelationshipType.EQUAL && to != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  } else if (from == TreeRelationshipType.EQUAL) {
    // op is moving the value that t already modified. Continue with op.
    return op.clone();
  } else if (to == TreeRelationshipType.EQUAL) {
    // op is overwriting the value that t already modified. In order to converge, op must take priority.
    return op.clone();
  } else {
    throw Exception("");
  }
}

pb.Op tRenameKeyEditKey(pb.Op t, pb.Op op, bool priority) {
  if (isNullMove(t)) {
    return op.clone();
  }
  final from = treeRelationship(t.location, op.location);
  final to = treeRelationship(toLoc(t), op.location);
  if (from != TreeRelationshipType.EQUAL && to != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  } else if (from == TreeRelationshipType.EQUAL) {
    // op is trying to modify the value that t moved. Op can continue with shifted key, which is correctly handled
    // by tIndependent.
    return tIndependent(t, op);
  } else if (to == TreeRelationshipType.EQUAL) {
    // op is trying to modify the value that t has already overwritten. In order to converge, the move must take
    // priority.
    return null;
  } else {
    throw Exception("");
  }
}

pb.Op tOverwrite(pb.Op t, pb.Op op, bool priority) {
  // Used by:
  // tSetFieldSetField
  // tSetIndexSetIndex
  // tSetKeySetKey
  // tSetFieldDeleteField
  // tSetKeyDeleteKey
  // tDeleteFieldSetField
  // tDeleteKeySetKey

  if (treeRelationship(t.location, op.location) != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  }
  // Op and t are both overwriting the same value. Use priority to determine the outcome.
  if (priority) {
    // When t has priority, remove op.
    return null;
  }
  return op.clone();
}

pb.Op tSetFieldSetField(pb.Op t, pb.Op op, bool priority) {
  return tOverwrite(t, op, priority);
}

pb.Op tSetIndexSetIndex(pb.Op t, pb.Op op, bool priority) {
  return tOverwrite(t, op, priority);
}

pb.Op tSetKeySetKey(pb.Op t, pb.Op op, bool priority) {
  return tOverwrite(t, op, priority);
}

pb.Op tSetFieldDeleteField(pb.Op t, pb.Op op, bool priority) {
  return tOverwrite(t, op, priority);
}

pb.Op tSetKeyDeleteKey(pb.Op t, pb.Op op, bool priority) {
  return tOverwrite(t, op, priority);
}

pb.Op tDeleteFieldSetField(pb.Op t, pb.Op op, bool priority) {
  return tOverwrite(t, op, priority);
}

pb.Op tDeleteKeySetKey(pb.Op t, pb.Op op, bool priority) {
  return tOverwrite(t, op, priority);
}

pb.Op tSetIndexInsertIndex(pb.Op t, pb.Op op, bool priority) {
  if (treeRelationship(t.location, op.location) != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  }
  // op is inserting at the same list index that t has already replaced. This inserts a new value, so is independent.
  return op.clone();
}

pb.Op tInsertIndexSetIndex(pb.Op t, pb.Op op, bool priority) {
  // op is replacing at the same index that t has inserted. op will replace at the shifted index and will be
  // independent. This is handled by tIndependent.
  return tIndependent(t, op);
}

pb.Op tSetKeyRenameKey(pb.Op t, pb.Op op, bool priority) {
  if (isNullMove(op)) {
    return null;
  }
  final from = treeRelationship(t.location, op.location);
  final to = treeRelationship(t.location, toLoc(op));
  if (from != TreeRelationshipType.EQUAL && to != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  } else if (from == TreeRelationshipType.EQUAL) {
    // op is attempting to move the value that t already overwrote. We can continue with op.
    return op.clone();
  } else if (to == TreeRelationshipType.EQUAL) {
    // op is attempting to overwrite the same value that t already overwrote. We can use priority to determine
    // which operation wins. However, if we remove op (the move operation), we must replace it with one that
    // deletes the "from" value (thus simulating the move operation running before the replace).
    if (priority) {
      return pb.Op()
        ..type = pb.Op_Type.Delete
        ..location.addAll(op.clone().location);
    }
    return op.clone();
  } else {
    throw Exception("");
  }
}

pb.Op tRenameKeySetKey(pb.Op t, pb.Op op, bool priority) {
  if (isNullMove(t)) {
    return op.clone();
  }
  final from = treeRelationship(t.location, op.location);
  final to = treeRelationship(toLoc(t), op.location);
  if (from != TreeRelationshipType.EQUAL && to != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  } else if (from == TreeRelationshipType.EQUAL) {
    // op is trying to replace the value that t moved. Op can continue with shifted key, which is correctly handled
    // by tIndependent.
    return tIndependent(t, op);
  } else if (to == TreeRelationshipType.EQUAL) {
    // op is trying to replace the value that t has overwritten. We can use priority to determine the winner.
    if (priority) {
      // if t has priority, we can just remove op
      return null;
    }
    // if op has priority, it can continue with a shifted key, which is correctly handled by tIndependent.
    return tIndependent(t, op);
  } else {
    throw Exception("");
  }
}

pb.Op tInsertIndexInsertIndex(pb.Op t, pb.Op op, bool priority) {
  if (treeRelationship(t.location, op.location) != TreeRelationshipType.EQUAL) {
    return tIndependent(t, op);
  }
  // Both operations are inserting at the same index. We should use priority to determine which is shifted.
  final shifter = insertLocationShifter(itemIndex(t), priority, true);
  var out = op.clone();
  setItemIndex(out, shifter(itemIndex(op)));
  return out;
}

//func tInsertIndexInsertIndex(t, op *Op, priority bool) *Op {
//	if treeRelationship(t.location, op.location) != TreeRelationshipType.EQUAL {
//		return tIndependent(t, op)
//	}
//	// Both operations are inserting at the same index. We should use priority to determine which is shifted.
//	shifter := insertLocationShifter(itemIndex(t), priority, true)
//	out := op.clone()
//	out.SetItemIndex(shifter(itemIndex(op)))
//	return out
//}
//
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tInsertIndexMoveIndex(t, op *Op, priority bool) *Op {
//	if (isNullMove(op)) {
//		return null
//	}
//	final from = treeRelationship(t.location, op.location)
//	final to = treeRelationship(t.location, toLoc(op))
//	switch {
//	case from != TreeRelationshipType.EQUAL && to != TreeRelationshipType.EQUAL:
//		return tIndependent(t, op)
//	case from == TreeRelationshipType.EQUAL:
//		// op is trying to move the value that was at the same index as t inserted at. We can just apply the shifter
//		// function to both from and to locations and the operations are independent. This is handled correctly by
//		// tIndependent.
//		return tIndependent(t, op)
//	case to == TreeRelationshipType.EQUAL:
//		// op is trying to move a value to the same index that t inserted at. We can use priority to determine
//		// which is shifted, but the destination index should be shifted without taking account of priority.
//		locationShifter := insertLocationShifter(itemIndex(t), priority, true)
//		valueShifter := insertValueShifter(itemIndex(t))
//		out := op.clone()
//		out.SetItemIndex(valueShifter(itemIndex(op)))
//		out.SetToIndex(locationShifter(op.ToIndex()))
//		return out
//	} else {
//		throw Exception("");
//	}
//}
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tMoveIndexInsertIndex(t, op *Op, priority bool) *Op {
//	if (isNullMove(t)) {
//		return op.clone()
//	}
//	final from = treeRelationship(t.location, op.location)
//	final to = treeRelationship(toLoc(t), op.location)
//	switch {
//	case from != TreeRelationshipType.EQUAL && to != TreeRelationshipType.EQUAL:
//		return tIndependent(t, op)
//	case from == TreeRelationshipType.EQUAL:
//		// op is inserting a new value at the same location that t moved from. Operations are independent, but we don't
//		// want to use tIndependent because that would make the insert move with the moved value (it uses the
//		// moveValueShifter shifter variant). We manually use moveLocationShifter to shift the location:
//
//		// Note: priority doesn't matter here because it is only used when (i == to) || (from < to && i == to+1)
//		// Since we know i == from, it cannot be used. Set to false to demonstrate this:
//		shifter := moveLocationShifter(itemIndex(t), t.ToIndex(), false, false)
//		out := op.clone()
//		out.SetItemIndex(shifter(itemIndex(op)))
//		return out
//	case to == TreeRelationshipType.EQUAL:
//		// op is inserting a new value at the same location that t moved to. We can priority to determine which
//		// operation is shifted.
//		shifter := moveLocationShifter(itemIndex(t), t.ToIndex(), priority, true)
//		out := op.clone()
//		out.SetItemIndex(shifter(itemIndex(op)))
//		return out
//	} else {
//		throw Exception("");
//	}
//}
//
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tInsertIndexDeleteIndex(t, op *Op, priority bool) *Op {
//	// op is deleting at the same index that t has inserted. op will delete at the shifted index and will be
//	// independent. This is handled by tIndependent.
//	return tIndependent(t, op)
//}
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tDeleteIndexInsertIndex(t, op *Op, priority bool) *Op {
//	// op is trying to insert at the same index as t deleted from. Operations are independent.
//	return tIndependent(t, op)
//}
//
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tMoveIndexMoveIndex(t, op *Op, priority bool) *Op {
//	if (isNullMove(op)) {
//		return null
//	}
//	if (isNullMove(t)) {
//		return op.clone()
//	}
//	fromFrom := treeRelationship(t.location, op.location)
//	fromTo := treeRelationship(t.location, toLoc(op))
//	toFrom := treeRelationship(toLoc(t), op.location)
//	toTo := treeRelationship(toLoc(t), toLoc(op))
//	switch {
//	case fromFrom == TreeRelationshipType.EQUAL && toTo == TreeRelationshipType.EQUAL:
//		// Op is trying to move the value that t already moved, and the "to" locations are the same. Operations are
//		// identical so we can simply delete op.
//		return null
//	case fromFrom == TreeRelationshipType.EQUAL:
//		// Op is trying to move the value that t already moved, and the "to" locations are different. So we can use
//		// priority ot determine which operation should win. If the transformer has priority, we can just remove op.
//		// If not, we use the index shifter to update the from and to location so that op moves the correct value to
//		// the intended location.
//		if priority {
//			return null
//		}
//
//		// Here all we need to do is modify the from location of op so that is uses the value at the to location of t.
//		// We don't need to use the shifter.
//		locationShifter := moveLocationShifter(itemIndex(t), t.ToIndex(), priority, true)
//		out := op.clone()
//		// must use location locationShifter here because that was used to shift t.To
//		out.SetItemIndex(locationShifter(t.ToIndex()))
//		out.SetToIndex(locationShifter(op.ToIndex()))
//		return out
//	case toTo == TreeRelationshipType.EQUAL:
//		// Op is trying to move another value to the same index that t just moved a value to. We can use priority to
//		// determine which value is shifted.
//		locationShifter := moveLocationShifter(itemIndex(t), t.ToIndex(), priority, true)
//		valueShifter := moveValueShifter(itemIndex(t), t.ToIndex())
//		out := op.clone()
//		out.SetItemIndex(valueShifter(itemIndex(op)))
//		out.SetToIndex(locationShifter(op.ToIndex()))
//		return out
//	case fromTo == TreeRelationshipType.EQUAL && toFrom == TreeRelationshipType.EQUAL:
//		// Op is trying to move the value at the to index of the move that t has just done, and move to the from index.
//		// Since both moves are non destructive and shift the other values around these operations are independent and
//		// can be handled by tIndependent.
//		return tIndependent(t, op)
//	case fromTo == TreeRelationshipType.EQUAL:
//		// Op is trying to move to the index that t just moved from. Operations are independent and can be handled by
//		// tIndependent.
//		return tIndependent(t, op)
//	case toFrom == TreeRelationshipType.EQUAL:
//		// Op is trying to move from the index that t just moved to. Operations are independent and can be handled by
//		// tIndependent.
//		return tIndependent(t, op)
//	} else {
//		// Operations are independent.
//		return tIndependent(t, op)
//	}
//}
//
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tRenameKeyRenameKey(t, op *Op, priority bool) *Op {
//	if (isNullMove(op)) {
//		return null
//	}
//	if (isNullMove(t)) {
//		return op.clone()
//	}
//	// t has moved the value from t.location and overwritten toLoc(t).
//	// op wants to move from op.location and overwrite toLoc(op)
//	toFrom := treeRelationship(toLoc(t), op.location)
//	toTo := treeRelationship(toLoc(t), toLoc(op))
//	fromFrom := treeRelationship(t.location, op.location)
//	fromTo := treeRelationship(t.location, toLoc(op))
//	switch {
//	case fromFrom == TreeRelationshipType.EQUAL && toTo == TreeRelationshipType.EQUAL:
//		// Op is trying to move the value that t already moved, and the "to" locations are the same. We can simply
//		// remove op.
//		return null
//	case fromFrom == TreeRelationshipType.EQUAL:
//		// Op is trying to move the value that t already moved, and the "to" locations are different. We can use
//		// priority to determine which to location is used. If t has priority, delete the op. If not, we change the
//		// From location to move the correct value. If we remove op, we must replace with an operation that deletes
//		// the "to" location.
//		if priority {
//			return &Op{
//				Type:     Op_Delete,
//				Location: op.clone().To(),
//			}
//		}
//		out := op.clone()
//		out.Location = proto.Clone(t).(*Op).To()
//		return out
//	case toTo == TreeRelationshipType.EQUAL:
//		// Op is trying to move a value and overwrite the value that t already overwrote. We can use priority to
//		// determine which value is used. If the transformer has priority, remove op. If not, continue with op. If
//		// we remove op, we must replace with an operation that deletes the "from" location.
//		if priority {
//			return &Op{
//				Type:     Op_Delete,
//				Location: op.clone().Location,
//			}
//		}
//		return op.clone()
//	case fromTo == TreeRelationshipType.EQUAL && toFrom == TreeRelationshipType.EQUAL:
//		// Op is trying to move the value that t has already overwritten, and the "to" location is the value that
//		// t moved. In order to converge, we must delete both values, so we replace op with a delete that removes at
//		// the From location.
//		return &Op{
//			Type:     Op_Delete,
//			Location: op.clone().Location,
//		}
//	case fromTo == TreeRelationshipType.EQUAL:
//		// Op is trying to overwrite the value that t already moved. We can simply run op with an updated "to"
//		// location.
//		out := op.clone()
//		out.Value.(*Op_Key).Key = proto.Clone(t).(*Op).Value.(*Op_Key).Key
//		return out
//	case toFrom == TreeRelationshipType.EQUAL:
//		// Op is trying to move the value that t already overwrote. We can continue with op.
//		return op.clone()
//	} else {
//		// independent operations
//		return op.clone()
//	}
//}
//
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tRenameKeyDeleteKey(t, op *Op, priority bool) *Op {
//	if (isNullMove(t)) {
//		return op.clone()
//	}
//	final from = treeRelationship(t.location, op.location)
//	final to = treeRelationship(toLoc(t), op.location)
//	switch {
//	case from != TreeRelationshipType.EQUAL && to != TreeRelationshipType.EQUAL:
//		return tIndependent(t, op)
//	case from == TreeRelationshipType.EQUAL:
//		// op is trying to delete the value that t moved. Op can continue with shifted key, which is correctly handled
//		// by tIndependent.
//		return tIndependent(t, op)
//	case to == TreeRelationshipType.EQUAL:
//		// op is trying to delete the value that t has already overwritten. We can simply remove op.
//		return null
//	} else {
//		throw Exception("");
//	}
//}
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tDeleteKeyRenameKey(t, op *Op, priority bool) *Op {
//	if (isNullMove(op)) {
//		return null
//	}
//	final from = treeRelationship(t.location, op.location)
//	final to = treeRelationship(t.location, toLoc(op))
//	switch {
//	case from != TreeRelationshipType.EQUAL && to != TreeRelationshipType.EQUAL:
//		return tIndependent(t, op)
//	case from == TreeRelationshipType.EQUAL:
//		// op is trying to move the value that op already deleted. In order to converge we must remove op and replace
//		// with an operation that deletes the "to" value.
//		return &Op{
//			Type:     Op_Delete,
//			Location: op.clone().To(),
//		}
//	case to == TreeRelationshipType.EQUAL:
//		// op is trying to overwrite the value that op already deleted. continue with op.
//		return op.clone()
//	} else {
//		throw Exception("");
//	}
//}
//
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tDelete(t, op *Op) *Op {
//	// Used by:
//	// tDeleteFieldDeleteField
//	// tDeleteIndexDeleteIndex
//	// tDeleteKeyDeleteKey
//
//	if treeRelationship(t.location, op.location) != TreeRelationshipType.EQUAL {
//		return tIndependent(t, op)
//	}
//	// op and t are both deleting the same value. We can remove op.
//	return null
//}
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
pb.Op t(pb.Op t, pb.Op op, bool priority) {}
//func tDeleteFieldDeleteField(t, op *Op, priority bool) *Op { return tDelete(t, op) }
//func tDeleteIndexDeleteIndex(t, op *Op, priority bool) *Op { return tDelete(t, op) }
//func tDeleteKeyDeleteKey(t, op *Op, priority bool) *Op     { return tDelete(t, op) }