// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: taa_engine/v1/taa_engine.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/go-kratos/kratos/v2/errors"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	ErrorReason_USER_UNSPECIFIED   ErrorReason = 0
	ErrorReason_USER_NOT_FOUND     ErrorReason = 1
	ErrorReason_USER_NOT_CREATED   ErrorReason = 2
	ErrorReason_USER_ALREADY_EXIST ErrorReason = 3
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "USER_UNSPECIFIED",
		1: "USER_NOT_FOUND",
		2: "USER_NOT_CREATED",
		3: "USER_ALREADY_EXIST",
	}
	ErrorReason_value = map[string]int32{
		"USER_UNSPECIFIED":   0,
		"USER_NOT_FOUND":     1,
		"USER_NOT_CREATED":   2,
		"USER_ALREADY_EXIST": 3,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_taa_engine_v1_taa_engine_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_taa_engine_v1_taa_engine_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_taa_engine_v1_taa_engine_proto_rawDescGZIP(), []int{0}
}

// Service info
type TaaEngineInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TaaEngineInfo) Reset() {
	*x = TaaEngineInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taa_engine_v1_taa_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaaEngineInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaaEngineInfo) ProtoMessage() {}

func (x *TaaEngineInfo) ProtoReflect() protoreflect.Message {
	mi := &file_taa_engine_v1_taa_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaaEngineInfo.ProtoReflect.Descriptor instead.
func (*TaaEngineInfo) Descriptor() ([]byte, []int) {
	return file_taa_engine_v1_taa_engine_proto_rawDescGZIP(), []int{0}
}

func (x *TaaEngineInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request message containing service request.
type CreateTaaEngineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateTaaEngineRequest) Reset() {
	*x = CreateTaaEngineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taa_engine_v1_taa_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaaEngineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaaEngineRequest) ProtoMessage() {}

func (x *CreateTaaEngineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taa_engine_v1_taa_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaaEngineRequest.ProtoReflect.Descriptor instead.
func (*CreateTaaEngineRequest) Descriptor() ([]byte, []int) {
	return file_taa_engine_v1_taa_engine_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaaEngineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing service response
type CreateTaaEngineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *TaaEngineInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateTaaEngineResponse) Reset() {
	*x = CreateTaaEngineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taa_engine_v1_taa_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaaEngineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaaEngineResponse) ProtoMessage() {}

func (x *CreateTaaEngineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_taa_engine_v1_taa_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaaEngineResponse.ProtoReflect.Descriptor instead.
func (*CreateTaaEngineResponse) Descriptor() ([]byte, []int) {
	return file_taa_engine_v1_taa_engine_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaaEngineResponse) GetInfo() *TaaEngineInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_taa_engine_v1_taa_engine_proto protoreflect.FileDescriptor

var file_taa_engine_v1_taa_engine_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x61, 0x61, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x61, 0x61, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x61, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x54, 0x61, 0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x61, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x2a, 0x83, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0xf4,
	0x03, 0x12, 0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1a, 0x0a, 0x10, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x32, 0x8e, 0x01, 0x0a, 0x09,
	0x54, 0x61, 0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x29, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x61, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x61, 0x61, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0e, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x42, 0x44, 0x0a, 0x11,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x61, 0x61, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x10, 0x54, 0x61, 0x61, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x74, 0x61, 0x61, 0x2d, 0x65, 0x78, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x61, 0x61, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_taa_engine_v1_taa_engine_proto_rawDescOnce sync.Once
	file_taa_engine_v1_taa_engine_proto_rawDescData = file_taa_engine_v1_taa_engine_proto_rawDesc
)

func file_taa_engine_v1_taa_engine_proto_rawDescGZIP() []byte {
	file_taa_engine_v1_taa_engine_proto_rawDescOnce.Do(func() {
		file_taa_engine_v1_taa_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_taa_engine_v1_taa_engine_proto_rawDescData)
	})
	return file_taa_engine_v1_taa_engine_proto_rawDescData
}

var file_taa_engine_v1_taa_engine_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_taa_engine_v1_taa_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_taa_engine_v1_taa_engine_proto_goTypes = []interface{}{
	(ErrorReason)(0),                // 0: api.taa_engine.v1.ErrorReason
	(*TaaEngineInfo)(nil),           // 1: api.taa_engine.v1.TaaEngineInfo
	(*CreateTaaEngineRequest)(nil),  // 2: api.taa_engine.v1.CreateTaaEngineRequest
	(*CreateTaaEngineResponse)(nil), // 3: api.taa_engine.v1.CreateTaaEngineResponse
}
var file_taa_engine_v1_taa_engine_proto_depIdxs = []int32{
	1, // 0: api.taa_engine.v1.CreateTaaEngineResponse.info:type_name -> api.taa_engine.v1.TaaEngineInfo
	2, // 1: api.taa_engine.v1.TaaEngine.CreateTaaEngine:input_type -> api.taa_engine.v1.CreateTaaEngineRequest
	3, // 2: api.taa_engine.v1.TaaEngine.CreateTaaEngine:output_type -> api.taa_engine.v1.CreateTaaEngineResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_taa_engine_v1_taa_engine_proto_init() }
func file_taa_engine_v1_taa_engine_proto_init() {
	if File_taa_engine_v1_taa_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_taa_engine_v1_taa_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaaEngineInfo); i {
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
		file_taa_engine_v1_taa_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaaEngineRequest); i {
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
		file_taa_engine_v1_taa_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaaEngineResponse); i {
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
			RawDescriptor: file_taa_engine_v1_taa_engine_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_taa_engine_v1_taa_engine_proto_goTypes,
		DependencyIndexes: file_taa_engine_v1_taa_engine_proto_depIdxs,
		EnumInfos:         file_taa_engine_v1_taa_engine_proto_enumTypes,
		MessageInfos:      file_taa_engine_v1_taa_engine_proto_msgTypes,
	}.Build()
	File_taa_engine_v1_taa_engine_proto = out.File
	file_taa_engine_v1_taa_engine_proto_rawDesc = nil
	file_taa_engine_v1_taa_engine_proto_goTypes = nil
	file_taa_engine_v1_taa_engine_proto_depIdxs = nil
}
