// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: ionscale/v1/iam.proto

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

type GetIAMPolicyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TailnetId     uint64                 `protobuf:"varint,1,opt,name=tailnet_id,json=tailnetId,proto3" json:"tailnet_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIAMPolicyRequest) Reset() {
	*x = GetIAMPolicyRequest{}
	mi := &file_ionscale_v1_iam_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIAMPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIAMPolicyRequest) ProtoMessage() {}

func (x *GetIAMPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_iam_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIAMPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetIAMPolicyRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_iam_proto_rawDescGZIP(), []int{0}
}

func (x *GetIAMPolicyRequest) GetTailnetId() uint64 {
	if x != nil {
		return x.TailnetId
	}
	return 0
}

type GetIAMPolicyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Policy        string                 `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIAMPolicyResponse) Reset() {
	*x = GetIAMPolicyResponse{}
	mi := &file_ionscale_v1_iam_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIAMPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIAMPolicyResponse) ProtoMessage() {}

func (x *GetIAMPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_iam_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIAMPolicyResponse.ProtoReflect.Descriptor instead.
func (*GetIAMPolicyResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_iam_proto_rawDescGZIP(), []int{1}
}

func (x *GetIAMPolicyResponse) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

type SetIAMPolicyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TailnetId     uint64                 `protobuf:"varint,1,opt,name=tailnet_id,json=tailnetId,proto3" json:"tailnet_id,omitempty"`
	Policy        string                 `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetIAMPolicyRequest) Reset() {
	*x = SetIAMPolicyRequest{}
	mi := &file_ionscale_v1_iam_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetIAMPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIAMPolicyRequest) ProtoMessage() {}

func (x *SetIAMPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_iam_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIAMPolicyRequest.ProtoReflect.Descriptor instead.
func (*SetIAMPolicyRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_iam_proto_rawDescGZIP(), []int{2}
}

func (x *SetIAMPolicyRequest) GetTailnetId() uint64 {
	if x != nil {
		return x.TailnetId
	}
	return 0
}

func (x *SetIAMPolicyRequest) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

type SetIAMPolicyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetIAMPolicyResponse) Reset() {
	*x = SetIAMPolicyResponse{}
	mi := &file_ionscale_v1_iam_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetIAMPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIAMPolicyResponse) ProtoMessage() {}

func (x *SetIAMPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_iam_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIAMPolicyResponse.ProtoReflect.Descriptor instead.
func (*SetIAMPolicyResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_iam_proto_rawDescGZIP(), []int{3}
}

var File_ionscale_v1_iam_proto protoreflect.FileDescriptor

var file_ionscale_v1_iam_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x41, 0x4d, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x49, 0x41, 0x4d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x4c, 0x0a, 0x13, 0x53, 0x65,
	0x74, 0x49, 0x41, 0x4d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x49,
	0x41, 0x4d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x73, 0x69, 0x65, 0x62, 0x65, 0x6e, 0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_ionscale_v1_iam_proto_rawDescOnce sync.Once
	file_ionscale_v1_iam_proto_rawDescData []byte
)

func file_ionscale_v1_iam_proto_rawDescGZIP() []byte {
	file_ionscale_v1_iam_proto_rawDescOnce.Do(func() {
		file_ionscale_v1_iam_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ionscale_v1_iam_proto_rawDesc), len(file_ionscale_v1_iam_proto_rawDesc)))
	})
	return file_ionscale_v1_iam_proto_rawDescData
}

var file_ionscale_v1_iam_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ionscale_v1_iam_proto_goTypes = []any{
	(*GetIAMPolicyRequest)(nil),  // 0: ionscale.v1.GetIAMPolicyRequest
	(*GetIAMPolicyResponse)(nil), // 1: ionscale.v1.GetIAMPolicyResponse
	(*SetIAMPolicyRequest)(nil),  // 2: ionscale.v1.SetIAMPolicyRequest
	(*SetIAMPolicyResponse)(nil), // 3: ionscale.v1.SetIAMPolicyResponse
}
var file_ionscale_v1_iam_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ionscale_v1_iam_proto_init() }
func file_ionscale_v1_iam_proto_init() {
	if File_ionscale_v1_iam_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ionscale_v1_iam_proto_rawDesc), len(file_ionscale_v1_iam_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ionscale_v1_iam_proto_goTypes,
		DependencyIndexes: file_ionscale_v1_iam_proto_depIdxs,
		MessageInfos:      file_ionscale_v1_iam_proto_msgTypes,
	}.Build()
	File_ionscale_v1_iam_proto = out.File
	file_ionscale_v1_iam_proto_goTypes = nil
	file_ionscale_v1_iam_proto_depIdxs = nil
}
