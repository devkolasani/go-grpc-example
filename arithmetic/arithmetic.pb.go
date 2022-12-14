// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: arithmetic/arithmetic.proto

package arithmetic

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

type OperationType int32

const (
	OperationType_ADD OperationType = 0
	OperationType_SUB OperationType = 1
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "ADD",
		1: "SUB",
	}
	OperationType_value = map[string]int32{
		"ADD": 0,
		"SUB": 1,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_arithmetic_arithmetic_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_arithmetic_arithmetic_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_arithmetic_arithmetic_proto_rawDescGZIP(), []int{0}
}

type OperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperandOne    int64         `protobuf:"varint,1,opt,name=operand_one,json=operandOne,proto3" json:"operand_one,omitempty"`
	OperandTwo    int64         `protobuf:"varint,2,opt,name=operand_two,json=operandTwo,proto3" json:"operand_two,omitempty"`
	OperationType OperationType `protobuf:"varint,3,opt,name=operation_type,json=operationType,proto3,enum=arithmetic.OperationType" json:"operation_type,omitempty"`
}

func (x *OperationRequest) Reset() {
	*x = OperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arithmetic_arithmetic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationRequest) ProtoMessage() {}

func (x *OperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_arithmetic_arithmetic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationRequest.ProtoReflect.Descriptor instead.
func (*OperationRequest) Descriptor() ([]byte, []int) {
	return file_arithmetic_arithmetic_proto_rawDescGZIP(), []int{0}
}

func (x *OperationRequest) GetOperandOne() int64 {
	if x != nil {
		return x.OperandOne
	}
	return 0
}

func (x *OperationRequest) GetOperandTwo() int64 {
	if x != nil {
		return x.OperandTwo
	}
	return 0
}

func (x *OperationRequest) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_ADD
}

type OperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OperationResponse) Reset() {
	*x = OperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arithmetic_arithmetic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResponse) ProtoMessage() {}

func (x *OperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arithmetic_arithmetic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResponse.ProtoReflect.Descriptor instead.
func (*OperationResponse) Descriptor() ([]byte, []int) {
	return file_arithmetic_arithmetic_proto_rawDescGZIP(), []int{1}
}

func (x *OperationResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_arithmetic_arithmetic_proto protoreflect.FileDescriptor

var file_arithmetic_arithmetic_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2f, 0x61, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x54, 0x77, 0x6f,
	0x12, 0x40, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a,
	0x21, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x55, 0x42,
	0x10, 0x01, 0x32, 0x58, 0x0a, 0x0a, 0x41, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63,
	0x12, 0x4a, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e,
	0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x6b, 0x6f,
	0x6c, 0x61, 0x73, 0x61, 0x6e, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arithmetic_arithmetic_proto_rawDescOnce sync.Once
	file_arithmetic_arithmetic_proto_rawDescData = file_arithmetic_arithmetic_proto_rawDesc
)

func file_arithmetic_arithmetic_proto_rawDescGZIP() []byte {
	file_arithmetic_arithmetic_proto_rawDescOnce.Do(func() {
		file_arithmetic_arithmetic_proto_rawDescData = protoimpl.X.CompressGZIP(file_arithmetic_arithmetic_proto_rawDescData)
	})
	return file_arithmetic_arithmetic_proto_rawDescData
}

var file_arithmetic_arithmetic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_arithmetic_arithmetic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_arithmetic_arithmetic_proto_goTypes = []interface{}{
	(OperationType)(0),        // 0: arithmetic.OperationType
	(*OperationRequest)(nil),  // 1: arithmetic.OperationRequest
	(*OperationResponse)(nil), // 2: arithmetic.OperationResponse
}
var file_arithmetic_arithmetic_proto_depIdxs = []int32{
	0, // 0: arithmetic.OperationRequest.operation_type:type_name -> arithmetic.OperationType
	1, // 1: arithmetic.Arithmetic.Operation:input_type -> arithmetic.OperationRequest
	2, // 2: arithmetic.Arithmetic.Operation:output_type -> arithmetic.OperationResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_arithmetic_arithmetic_proto_init() }
func file_arithmetic_arithmetic_proto_init() {
	if File_arithmetic_arithmetic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arithmetic_arithmetic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationRequest); i {
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
		file_arithmetic_arithmetic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationResponse); i {
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
			RawDescriptor: file_arithmetic_arithmetic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arithmetic_arithmetic_proto_goTypes,
		DependencyIndexes: file_arithmetic_arithmetic_proto_depIdxs,
		EnumInfos:         file_arithmetic_arithmetic_proto_enumTypes,
		MessageInfos:      file_arithmetic_arithmetic_proto_msgTypes,
	}.Build()
	File_arithmetic_arithmetic_proto = out.File
	file_arithmetic_arithmetic_proto_rawDesc = nil
	file_arithmetic_arithmetic_proto_goTypes = nil
	file_arithmetic_arithmetic_proto_depIdxs = nil
}
