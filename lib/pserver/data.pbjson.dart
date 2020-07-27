///
//  Generated code. Do not modify.
//  source: pserver/data.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Payload$json = const {
  '1': 'Payload',
  '3': const [Payload_Edit$json, Payload_Add$json, Payload_Get$json],
};

const Payload_Edit$json = const {
  '1': 'Edit',
  '3': const [Payload_Edit_Request$json, Payload_Edit_Response$json],
};

const Payload_Edit_Request$json = const {
  '1': 'Request',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'request', '3': 2, '4': 1, '5': 9, '10': 'request'},
    const {'1': 'state', '3': 3, '4': 1, '5': 3, '10': 'state'},
    const {'1': 'op', '3': 4, '4': 1, '5': 11, '6': '.delta.Op', '10': 'op'},
  ],
};

const Payload_Edit_Response$json = const {
  '1': 'Response',
  '2': const [
    const {'1': 'state', '3': 1, '4': 1, '5': 3, '10': 'state'},
    const {'1': 'op', '3': 2, '4': 1, '5': 11, '6': '.delta.Op', '10': 'op'},
  ],
};

const Payload_Add$json = const {
  '1': 'Add',
  '3': const [Payload_Add_Request$json, Payload_Add_Response$json],
};

const Payload_Add_Request$json = const {
  '1': 'Request',
  '2': const [
    const {'1': 'request', '3': 1, '4': 1, '5': 9, '10': 'request'},
  ],
};

const Payload_Add_Response$json = const {
  '1': 'Response',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

const Payload_Get$json = const {
  '1': 'Get',
  '3': const [Payload_Get_Request$json, Payload_Get_Response$json],
};

const Payload_Get_Request$json = const {
  '1': 'Request',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

const Payload_Get_Response$json = const {
  '1': 'Response',
  '2': const [
    const {'1': 'state', '3': 1, '4': 1, '5': 3, '10': 'state'},
  ],
};

const Snapshot$json = const {
  '1': 'Snapshot',
  '2': const [
    const {'1': 'request', '3': 1, '4': 1, '5': 9, '10': 'request'},
    const {'1': 'state', '3': 2, '4': 1, '5': 3, '10': 'state'},
    const {'1': 'value', '3': 3, '4': 1, '5': 11, '6': '.pserver.Blob', '10': 'value'},
  ],
};

const State$json = const {
  '1': 'State',
  '2': const [
    const {'1': 'request', '3': 1, '4': 1, '5': 9, '10': 'request'},
    const {'1': 'state', '3': 2, '4': 1, '5': 3, '10': 'state'},
    const {'1': 'op', '3': 3, '4': 1, '5': 11, '6': '.pserver.Blob', '10': 'op'},
  ],
};

const Blob$json = const {
  '1': 'Blob',
  '2': const [
    const {'1': 'value', '3': 1, '4': 1, '5': 12, '10': 'value'},
  ],
};

