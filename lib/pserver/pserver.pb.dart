///
//  Generated code. Do not modify.
//  source: pserver/pserver.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class Snapshot extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Snapshot', package: const $pb.PackageName('pserver'), createEmptyInstance: create)
    ..aInt64(1, 'state')
    ..aOM<Blob>(2, 'value', subBuilder: Blob.create)
    ..hasRequiredFields = false
  ;

  Snapshot._() : super();
  factory Snapshot() => create();
  factory Snapshot.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Snapshot.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Snapshot clone() => Snapshot()..mergeFromMessage(this);
  Snapshot copyWith(void Function(Snapshot) updates) => super.copyWith((message) => updates(message as Snapshot));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Snapshot create() => Snapshot._();
  Snapshot createEmptyInstance() => create();
  static $pb.PbList<Snapshot> createRepeated() => $pb.PbList<Snapshot>();
  @$core.pragma('dart2js:noInline')
  static Snapshot getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Snapshot>(create);
  static Snapshot _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get state => $_getI64(0);
  @$pb.TagNumber(1)
  set state($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasState() => $_has(0);
  @$pb.TagNumber(1)
  void clearState() => clearField(1);

  @$pb.TagNumber(2)
  Blob get value => $_getN(1);
  @$pb.TagNumber(2)
  set value(Blob v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => clearField(2);
  @$pb.TagNumber(2)
  Blob ensureValue() => $_ensure(1);
}

class State extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('State', package: const $pb.PackageName('pserver'), createEmptyInstance: create)
    ..aInt64(1, 'state')
    ..aOM<Blob>(2, 'op', subBuilder: Blob.create)
    ..hasRequiredFields = false
  ;

  State._() : super();
  factory State() => create();
  factory State.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory State.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  State clone() => State()..mergeFromMessage(this);
  State copyWith(void Function(State) updates) => super.copyWith((message) => updates(message as State));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static State create() => State._();
  State createEmptyInstance() => create();
  static $pb.PbList<State> createRepeated() => $pb.PbList<State>();
  @$core.pragma('dart2js:noInline')
  static State getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<State>(create);
  static State _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get state => $_getI64(0);
  @$pb.TagNumber(1)
  set state($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasState() => $_has(0);
  @$pb.TagNumber(1)
  void clearState() => clearField(1);

  @$pb.TagNumber(2)
  Blob get op => $_getN(1);
  @$pb.TagNumber(2)
  set op(Blob v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasOp() => $_has(1);
  @$pb.TagNumber(2)
  void clearOp() => clearField(2);
  @$pb.TagNumber(2)
  Blob ensureOp() => $_ensure(1);
}

class Blob extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Blob', package: const $pb.PackageName('pserver'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, 'value', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  Blob._() : super();
  factory Blob() => create();
  factory Blob.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Blob.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Blob clone() => Blob()..mergeFromMessage(this);
  Blob copyWith(void Function(Blob) updates) => super.copyWith((message) => updates(message as Blob));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Blob create() => Blob._();
  Blob createEmptyInstance() => create();
  static $pb.PbList<Blob> createRepeated() => $pb.PbList<Blob>();
  @$core.pragma('dart2js:noInline')
  static Blob getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Blob>(create);
  static Blob _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.int> get value => $_getN(0);
  @$pb.TagNumber(1)
  set value($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasValue() => $_has(0);
  @$pb.TagNumber(1)
  void clearValue() => clearField(1);
}

