// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: pkg/outoushuugou/messages.proto

package outoushuugou

import (
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

type Status int32

const (
	Status_Unknown Status = 0
	Status_Success Status = 1
	Status_Failed  Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Unknown",
		1: "Success",
		2: "Failed",
	}
	Status_value = map[string]int32{
		"Unknown": 0,
		"Success": 1,
		"Failed":  2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_outoushuugou_messages_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pkg_outoushuugou_messages_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_outoushuugou_messages_proto_rawDescGZIP(), []int{0}
}

// Activator-Response -> Katyusha-Outou
type ResponseFeedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint32                           `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SourceIP string                           `protobuf:"bytes,2,opt,name=SourceIP,proto3" json:"SourceIP,omitempty"`
	Domain   string                           `protobuf:"bytes,3,opt,name=Domain,proto3" json:"Domain,omitempty"`
	URI      string                           `protobuf:"bytes,4,opt,name=URI,proto3" json:"URI,omitempty"`
	Method   string                           `protobuf:"bytes,5,opt,name=Method,proto3" json:"Method,omitempty"`
	Headers  []*ResponseFeedback_HeaderSchema `protobuf:"bytes,6,rep,name=Headers,proto3" json:"Headers,omitempty"`
}

func (x *ResponseFeedback) Reset() {
	*x = ResponseFeedback{}
	mi := &file_pkg_outoushuugou_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFeedback) ProtoMessage() {}

func (x *ResponseFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_outoushuugou_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFeedback.ProtoReflect.Descriptor instead.
func (*ResponseFeedback) Descriptor() ([]byte, []int) {
	return file_pkg_outoushuugou_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseFeedback) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ResponseFeedback) GetSourceIP() string {
	if x != nil {
		return x.SourceIP
	}
	return ""
}

func (x *ResponseFeedback) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ResponseFeedback) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

func (x *ResponseFeedback) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ResponseFeedback) GetHeaders() []*ResponseFeedback_HeaderSchema {
	if x != nil {
		return x.Headers
	}
	return nil
}

type ResponseConfirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SymbolizeResponse Status `protobuf:"varint,1,opt,name=SymbolizeResponse,proto3,enum=outoushuugou.Status" json:"SymbolizeResponse,omitempty"`
}

func (x *ResponseConfirm) Reset() {
	*x = ResponseConfirm{}
	mi := &file_pkg_outoushuugou_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseConfirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseConfirm) ProtoMessage() {}

func (x *ResponseConfirm) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_outoushuugou_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseConfirm.ProtoReflect.Descriptor instead.
func (*ResponseConfirm) Descriptor() ([]byte, []int) {
	return file_pkg_outoushuugou_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseConfirm) GetSymbolizeResponse() Status {
	if x != nil {
		return x.SymbolizeResponse
	}
	return Status_Unknown
}

type ResponseFeedback_HeaderSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ResponseFeedback_HeaderSchema) Reset() {
	*x = ResponseFeedback_HeaderSchema{}
	mi := &file_pkg_outoushuugou_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseFeedback_HeaderSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFeedback_HeaderSchema) ProtoMessage() {}

func (x *ResponseFeedback_HeaderSchema) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_outoushuugou_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFeedback_HeaderSchema.ProtoReflect.Descriptor instead.
func (*ResponseFeedback_HeaderSchema) Descriptor() ([]byte, []int) {
	return file_pkg_outoushuugou_messages_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ResponseFeedback_HeaderSchema) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ResponseFeedback_HeaderSchema) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pkg_outoushuugou_messages_proto protoreflect.FileDescriptor

var file_pkg_outoushuugou_messages_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x75, 0x74, 0x6f, 0x75, 0x73, 0x68, 0x75, 0x75, 0x67,
	0x6f, 0x75, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6f, 0x75, 0x74, 0x6f, 0x75, 0x73, 0x68, 0x75, 0x75, 0x67, 0x6f, 0x75, 0x22,
	0x83, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x50,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x50,
	0x12, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x49, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x49, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x45, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x75, 0x73, 0x68, 0x75, 0x75, 0x67,
	0x6f, 0x75, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x52, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x55, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x42, 0x0a, 0x11, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x75, 0x73, 0x68, 0x75, 0x75, 0x67,
	0x6f, 0x75, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x11, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x2e, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x42, 0x12, 0x5a, 0x10,
	0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x75, 0x74, 0x6f, 0x75, 0x73, 0x68, 0x75, 0x75, 0x67, 0x6f, 0x75,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_outoushuugou_messages_proto_rawDescOnce sync.Once
	file_pkg_outoushuugou_messages_proto_rawDescData = file_pkg_outoushuugou_messages_proto_rawDesc
)

func file_pkg_outoushuugou_messages_proto_rawDescGZIP() []byte {
	file_pkg_outoushuugou_messages_proto_rawDescOnce.Do(func() {
		file_pkg_outoushuugou_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_outoushuugou_messages_proto_rawDescData)
	})
	return file_pkg_outoushuugou_messages_proto_rawDescData
}

var file_pkg_outoushuugou_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_outoushuugou_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_outoushuugou_messages_proto_goTypes = []any{
	(Status)(0),                           // 0: outoushuugou.Status
	(*ResponseFeedback)(nil),              // 1: outoushuugou.ResponseFeedback
	(*ResponseConfirm)(nil),               // 2: outoushuugou.ResponseConfirm
	(*ResponseFeedback_HeaderSchema)(nil), // 3: outoushuugou.ResponseFeedback.HeaderSchema
}
var file_pkg_outoushuugou_messages_proto_depIdxs = []int32{
	3, // 0: outoushuugou.ResponseFeedback.Headers:type_name -> outoushuugou.ResponseFeedback.HeaderSchema
	0, // 1: outoushuugou.ResponseConfirm.SymbolizeResponse:type_name -> outoushuugou.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_outoushuugou_messages_proto_init() }
func file_pkg_outoushuugou_messages_proto_init() {
	if File_pkg_outoushuugou_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_outoushuugou_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_outoushuugou_messages_proto_goTypes,
		DependencyIndexes: file_pkg_outoushuugou_messages_proto_depIdxs,
		EnumInfos:         file_pkg_outoushuugou_messages_proto_enumTypes,
		MessageInfos:      file_pkg_outoushuugou_messages_proto_msgTypes,
	}.Build()
	File_pkg_outoushuugou_messages_proto = out.File
	file_pkg_outoushuugou_messages_proto_rawDesc = nil
	file_pkg_outoushuugou_messages_proto_goTypes = nil
	file_pkg_outoushuugou_messages_proto_depIdxs = nil
}