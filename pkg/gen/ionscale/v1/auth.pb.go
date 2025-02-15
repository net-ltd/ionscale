// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: ionscale/v1/auth.proto

package ionscalev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthenticateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	mi := &file_ionscale_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_auth_proto_rawDescGZIP(), []int{0}
}

type AuthenticateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AuthUrl       string                 `protobuf:"bytes,1,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TailnetId     *uint64                `protobuf:"varint,3,opt,name=tailnet_id,json=tailnetId,proto3,oneof" json:"tailnet_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	mi := &file_ionscale_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateResponse) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *AuthenticateResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthenticateResponse) GetTailnetId() uint64 {
	if x != nil && x.TailnetId != nil {
		return *x.TailnetId
	}
	return 0
}

var File_ionscale_v1_auth_proto protoreflect.FileDescriptor

var file_ionscale_v1_auth_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7a, 0x0a, 0x14,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x09, 0x74, 0x61, 0x69,
	0x6c, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x61,
	0x69, 0x6c, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x69, 0x65, 0x62, 0x65, 0x6e, 0x73, 0x2f,
	0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6f, 0x6e,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_ionscale_v1_auth_proto_rawDescOnce sync.Once
	file_ionscale_v1_auth_proto_rawDescData []byte
)

func file_ionscale_v1_auth_proto_rawDescGZIP() []byte {
	file_ionscale_v1_auth_proto_rawDescOnce.Do(func() {
		file_ionscale_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ionscale_v1_auth_proto_rawDesc), len(file_ionscale_v1_auth_proto_rawDesc)))
	})
	return file_ionscale_v1_auth_proto_rawDescData
}

var file_ionscale_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ionscale_v1_auth_proto_goTypes = []any{
	(*AuthenticateRequest)(nil),  // 0: ionscale.v1.AuthenticateRequest
	(*AuthenticateResponse)(nil), // 1: ionscale.v1.AuthenticateResponse
}
var file_ionscale_v1_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ionscale_v1_auth_proto_init() }
func file_ionscale_v1_auth_proto_init() {
	if File_ionscale_v1_auth_proto != nil {
		return
	}
	file_ionscale_v1_auth_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ionscale_v1_auth_proto_rawDesc), len(file_ionscale_v1_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ionscale_v1_auth_proto_goTypes,
		DependencyIndexes: file_ionscale_v1_auth_proto_depIdxs,
		MessageInfos:      file_ionscale_v1_auth_proto_msgTypes,
	}.Build()
	File_ionscale_v1_auth_proto = out.File
	file_ionscale_v1_auth_proto_goTypes = nil
	file_ionscale_v1_auth_proto_depIdxs = nil
}
