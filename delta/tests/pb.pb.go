// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: pb.proto

package tests

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// [proto:data]
type Person_Type int32

const (
	Person_Alpha   Person_Type = 0
	Person_Bravo   Person_Type = 1
	Person_Charlie Person_Type = 2
)

// Enum value maps for Person_Type.
var (
	Person_Type_name = map[int32]string{
		0: "Alpha",
		1: "Bravo",
		2: "Charlie",
	}
	Person_Type_value = map[string]int32{
		"Alpha":   0,
		"Bravo":   1,
		"Charlie": 2,
	}
)

func (x Person_Type) Enum() *Person_Type {
	p := new(Person_Type)
	*p = x
	return p
}

func (x Person_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_proto_enumTypes[0].Descriptor()
}

func (Person_Type) Type() protoreflect.EnumType {
	return &file_pb_proto_enumTypes[0]
}

func (x Person_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_Type.Descriptor instead.
func (Person_Type) EnumDescriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{0, 0}
}

// [proto:data]
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age      uint32                 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Cases    map[string]*Case       `protobuf:"bytes,4,rep,name=cases,proto3" json:"cases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Company  *Company               `protobuf:"bytes,5,opt,name=company,proto3" json:"company,omitempty"`
	Alias    []string               `protobuf:"bytes,6,rep,name=alias,proto3" json:"alias,omitempty"`
	Type     Person_Type            `protobuf:"varint,7,opt,name=type,proto3,enum=tests.Person_Type" json:"type,omitempty"`
	TypeList []Person_Type          `protobuf:"varint,8,rep,packed,name=typeList,proto3,enum=tests.Person_Type" json:"typeList,omitempty"`
	TypeMap  map[string]Person_Type `protobuf:"bytes,9,rep,name=typeMap,proto3" json:"typeMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=tests.Person_Type"`
	Embedded *Person_Embed          `protobuf:"bytes,10,opt,name=embedded,proto3" json:"embedded,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetCases() map[string]*Case {
	if x != nil {
		return x.Cases
	}
	return nil
}

func (x *Person) GetCompany() *Company {
	if x != nil {
		return x.Company
	}
	return nil
}

func (x *Person) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Person) GetType() Person_Type {
	if x != nil {
		return x.Type
	}
	return Person_Alpha
}

func (x *Person) GetTypeList() []Person_Type {
	if x != nil {
		return x.TypeList
	}
	return nil
}

func (x *Person) GetTypeMap() map[string]Person_Type {
	if x != nil {
		return x.TypeMap
	}
	return nil
}

func (x *Person) GetEmbedded() *Person_Embed {
	if x != nil {
		return x.Embedded
	}
	return nil
}

// [proto:data]
type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string           `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Revenue float32          `protobuf:"fixed32,12,opt,name=revenue,proto3" json:"revenue,omitempty"`
	Flags   map[int64]string `protobuf:"bytes,13,rep,name=flags,proto3" json:"flags,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{1}
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetRevenue() float32 {
	if x != nil {
		return x.Revenue
	}
	return 0
}

func (x *Company) GetFlags() map[int64]string {
	if x != nil {
		return x.Flags
	}
	return nil
}

// [proto:data]
type Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	Items []*Item `protobuf:"bytes,22,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Case) Reset() {
	*x = Case{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Case) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Case) ProtoMessage() {}

func (x *Case) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Case.ProtoReflect.Descriptor instead.
func (*Case) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{2}
}

func (x *Case) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Case) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// [proto:data]
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string   `protobuf:"bytes,31,opt,name=title,proto3" json:"title,omitempty"`
	Done  bool     `protobuf:"varint,34,opt,name=done,proto3" json:"done,omitempty"`
	Flags []string `protobuf:"bytes,35,rep,name=flags,proto3" json:"flags,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Item) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *Item) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

type Holder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numbers map[string]int32 `protobuf:"bytes,1,rep,name=numbers,proto3" json:"numbers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Holder) Reset() {
	*x = Holder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Holder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Holder) ProtoMessage() {}

func (x *Holder) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Holder.ProtoReflect.Descriptor instead.
func (*Holder) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{4}
}

func (x *Holder) GetNumbers() map[string]int32 {
	if x != nil {
		return x.Numbers
	}
	return nil
}

type Embedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmbeddedLevel0 string `protobuf:"bytes,1,opt,name=embeddedLevel0,proto3" json:"embeddedLevel0,omitempty"`
}

func (x *Embedded) Reset() {
	*x = Embedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embedded) ProtoMessage() {}

func (x *Embedded) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embedded.ProtoReflect.Descriptor instead.
func (*Embedded) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{5}
}

func (x *Embedded) GetEmbeddedLevel0() string {
	if x != nil {
		return x.EmbeddedLevel0
	}
	return ""
}

type ScopeLevel1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T *ScopeLevel1_Embedded `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
}

func (x *ScopeLevel1) Reset() {
	*x = ScopeLevel1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeLevel1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeLevel1) ProtoMessage() {}

func (x *ScopeLevel1) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeLevel1.ProtoReflect.Descriptor instead.
func (*ScopeLevel1) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{6}
}

func (x *ScopeLevel1) GetT() *ScopeLevel1_Embedded {
	if x != nil {
		return x.T
	}
	return nil
}

// [proto:data]
type Person_Embed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Person_Embed) Reset() {
	*x = Person_Embed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Embed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Embed) ProtoMessage() {}

func (x *Person_Embed) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Embed.ProtoReflect.Descriptor instead.
func (*Person_Embed) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Person_Embed) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ScopeLevel1_Embedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmbeddedLevel1 string `protobuf:"bytes,1,opt,name=embeddedLevel1,proto3" json:"embeddedLevel1,omitempty"`
}

func (x *ScopeLevel1_Embedded) Reset() {
	*x = ScopeLevel1_Embedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeLevel1_Embedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeLevel1_Embedded) ProtoMessage() {}

func (x *ScopeLevel1_Embedded) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeLevel1_Embedded.ProtoReflect.Descriptor instead.
func (*ScopeLevel1_Embedded) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ScopeLevel1_Embedded) GetEmbeddedLevel1() string {
	if x != nil {
		return x.EmbeddedLevel1
	}
	return ""
}

type ScopeLevel1_ScopeLevel2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T *ScopeLevel1_ScopeLevel2_Embedded `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
}

func (x *ScopeLevel1_ScopeLevel2) Reset() {
	*x = ScopeLevel1_ScopeLevel2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeLevel1_ScopeLevel2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeLevel1_ScopeLevel2) ProtoMessage() {}

func (x *ScopeLevel1_ScopeLevel2) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeLevel1_ScopeLevel2.ProtoReflect.Descriptor instead.
func (*ScopeLevel1_ScopeLevel2) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{6, 1}
}

func (x *ScopeLevel1_ScopeLevel2) GetT() *ScopeLevel1_ScopeLevel2_Embedded {
	if x != nil {
		return x.T
	}
	return nil
}

type ScopeLevel1_ScopeLevel2_Embedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmbeddedLevel2 string `protobuf:"bytes,1,opt,name=embeddedLevel2,proto3" json:"embeddedLevel2,omitempty"`
}

func (x *ScopeLevel1_ScopeLevel2_Embedded) Reset() {
	*x = ScopeLevel1_ScopeLevel2_Embedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeLevel1_ScopeLevel2_Embedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeLevel1_ScopeLevel2_Embedded) ProtoMessage() {}

func (x *ScopeLevel1_ScopeLevel2_Embedded) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeLevel1_ScopeLevel2_Embedded.ProtoReflect.Descriptor instead.
func (*ScopeLevel1_ScopeLevel2_Embedded) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{6, 1, 0}
}

func (x *ScopeLevel1_ScopeLevel2_Embedded) GetEmbeddedLevel2() string {
	if x != nil {
		return x.EmbeddedLevel2
	}
	return ""
}

type ScopeLevel1_ScopeLevel2_ScopeLevel3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T *ScopeLevel1_ScopeLevel2_Embedded `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
}

func (x *ScopeLevel1_ScopeLevel2_ScopeLevel3) Reset() {
	*x = ScopeLevel1_ScopeLevel2_ScopeLevel3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeLevel1_ScopeLevel2_ScopeLevel3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeLevel1_ScopeLevel2_ScopeLevel3) ProtoMessage() {}

func (x *ScopeLevel1_ScopeLevel2_ScopeLevel3) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeLevel1_ScopeLevel2_ScopeLevel3.ProtoReflect.Descriptor instead.
func (*ScopeLevel1_ScopeLevel2_ScopeLevel3) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{6, 1, 1}
}

func (x *ScopeLevel1_ScopeLevel2_ScopeLevel3) GetT() *ScopeLevel1_ScopeLevel2_Embedded {
	if x != nil {
		return x.T
	}
	return nil
}

var File_pb_proto protoreflect.FileDescriptor

var file_pb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x22, 0xbc, 0x04, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x74, 0x79,
	0x70, 0x65, 0x4d, 0x61, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70,
	0x12, 0x2f, 0x0a, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x52, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65,
	0x64, 0x1a, 0x1b, 0x0a, 0x05, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x45,
	0x0a, 0x0a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e, 0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x76,
	0x6f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x10, 0x02,
	0x22, 0xa2, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x46, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x22, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18,
	0x23, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x22, 0x7a, 0x0a, 0x06,
	0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x32, 0x0a, 0x08, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x30, 0x22, 0xad, 0x02, 0x0a,
	0x0b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x31, 0x12, 0x29, 0x0a, 0x01,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x31, 0x2e, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x65, 0x64, 0x52, 0x01, 0x74, 0x1a, 0x32, 0x0a, 0x08, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6d, 0x62,
	0x65, 0x64, 0x64, 0x65, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x31, 0x1a, 0xbe, 0x01, 0x0a, 0x0b,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x32, 0x12, 0x35, 0x0a, 0x01, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x32, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x52,
	0x01, 0x74, 0x1a, 0x32, 0x0a, 0x08, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x32,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x32, 0x1a, 0x44, 0x0a, 0x0b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x33, 0x12, 0x35, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x32,
	0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x52, 0x01, 0x74, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x76, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x2f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_proto_rawDescOnce sync.Once
	file_pb_proto_rawDescData = file_pb_proto_rawDesc
)

func file_pb_proto_rawDescGZIP() []byte {
	file_pb_proto_rawDescOnce.Do(func() {
		file_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_proto_rawDescData)
	})
	return file_pb_proto_rawDescData
}

var file_pb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_pb_proto_goTypes = []interface{}{
	(Person_Type)(0),                         // 0: tests.Person.Type
	(*Person)(nil),                           // 1: tests.Person
	(*Company)(nil),                          // 2: tests.Company
	(*Case)(nil),                             // 3: tests.Case
	(*Item)(nil),                             // 4: tests.Item
	(*Holder)(nil),                           // 5: tests.Holder
	(*Embedded)(nil),                         // 6: tests.Embedded
	(*ScopeLevel1)(nil),                      // 7: tests.ScopeLevel1
	(*Person_Embed)(nil),                     // 8: tests.Person.Embed
	nil,                                      // 9: tests.Person.CasesEntry
	nil,                                      // 10: tests.Person.TypeMapEntry
	nil,                                      // 11: tests.Company.FlagsEntry
	nil,                                      // 12: tests.Holder.NumbersEntry
	(*ScopeLevel1_Embedded)(nil),             // 13: tests.ScopeLevel1.Embedded
	(*ScopeLevel1_ScopeLevel2)(nil),          // 14: tests.ScopeLevel1.ScopeLevel2
	(*ScopeLevel1_ScopeLevel2_Embedded)(nil), // 15: tests.ScopeLevel1.ScopeLevel2.Embedded
	(*ScopeLevel1_ScopeLevel2_ScopeLevel3)(nil), // 16: tests.ScopeLevel1.ScopeLevel2.ScopeLevel3
}
var file_pb_proto_depIdxs = []int32{
	9,  // 0: tests.Person.cases:type_name -> tests.Person.CasesEntry
	2,  // 1: tests.Person.company:type_name -> tests.Company
	0,  // 2: tests.Person.type:type_name -> tests.Person.Type
	0,  // 3: tests.Person.typeList:type_name -> tests.Person.Type
	10, // 4: tests.Person.typeMap:type_name -> tests.Person.TypeMapEntry
	8,  // 5: tests.Person.embedded:type_name -> tests.Person.Embed
	11, // 6: tests.Company.flags:type_name -> tests.Company.FlagsEntry
	4,  // 7: tests.Case.items:type_name -> tests.Item
	12, // 8: tests.Holder.numbers:type_name -> tests.Holder.NumbersEntry
	13, // 9: tests.ScopeLevel1.t:type_name -> tests.ScopeLevel1.Embedded
	3,  // 10: tests.Person.CasesEntry.value:type_name -> tests.Case
	0,  // 11: tests.Person.TypeMapEntry.value:type_name -> tests.Person.Type
	15, // 12: tests.ScopeLevel1.ScopeLevel2.t:type_name -> tests.ScopeLevel1.ScopeLevel2.Embedded
	15, // 13: tests.ScopeLevel1.ScopeLevel2.ScopeLevel3.t:type_name -> tests.ScopeLevel1.ScopeLevel2.Embedded
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_pb_proto_init() }
func file_pb_proto_init() {
	if File_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Case); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Holder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embedded); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeLevel1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Embed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeLevel1_Embedded); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeLevel1_ScopeLevel2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeLevel1_ScopeLevel2_Embedded); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeLevel1_ScopeLevel2_ScopeLevel3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_proto_goTypes,
		DependencyIndexes: file_pb_proto_depIdxs,
		EnumInfos:         file_pb_proto_enumTypes,
		MessageInfos:      file_pb_proto_msgTypes,
	}.Build()
	File_pb_proto = out.File
	file_pb_proto_rawDesc = nil
	file_pb_proto_goTypes = nil
	file_pb_proto_depIdxs = nil
}
