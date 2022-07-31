// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: app/grpc/pb/confirmation.proto

package rpc_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FindUserByEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FindUserByEmailReq) Reset() {
	*x = FindUserByEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_pb_confirmation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByEmailReq) ProtoMessage() {}

func (x *FindUserByEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_pb_confirmation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByEmailReq.ProtoReflect.Descriptor instead.
func (*FindUserByEmailReq) Descriptor() ([]byte, []int) {
	return file_app_grpc_pb_confirmation_proto_rawDescGZIP(), []int{0}
}

func (x *FindUserByEmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FindUserByEmailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
}

func (x *FindUserByEmailResp) Reset() {
	*x = FindUserByEmailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_pb_confirmation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByEmailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByEmailResp) ProtoMessage() {}

func (x *FindUserByEmailResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_pb_confirmation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByEmailResp.ProtoReflect.Descriptor instead.
func (*FindUserByEmailResp) Descriptor() ([]byte, []int) {
	return file_app_grpc_pb_confirmation_proto_rawDescGZIP(), []int{1}
}

func (x *FindUserByEmailResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindUserByEmailResp) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

type CheckActivationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
}

func (x *CheckActivationReq) Reset() {
	*x = CheckActivationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_pb_confirmation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckActivationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckActivationReq) ProtoMessage() {}

func (x *CheckActivationReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_pb_confirmation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckActivationReq.ProtoReflect.Descriptor instead.
func (*CheckActivationReq) Descriptor() ([]byte, []int) {
	return file_app_grpc_pb_confirmation_proto_rawDescGZIP(), []int{2}
}

func (x *CheckActivationReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckActivationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsActive bool `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *CheckActivationResp) Reset() {
	*x = CheckActivationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_pb_confirmation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckActivationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckActivationResp) ProtoMessage() {}

func (x *CheckActivationResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_pb_confirmation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckActivationResp.ProtoReflect.Descriptor instead.
func (*CheckActivationResp) Descriptor() ([]byte, []int) {
	return file_app_grpc_pb_confirmation_proto_rawDescGZIP(), []int{3}
}

func (x *CheckActivationResp) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type ActiveUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ActiveUserReq) Reset() {
	*x = ActiveUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_pb_confirmation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveUserReq) ProtoMessage() {}

func (x *ActiveUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_pb_confirmation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveUserReq.ProtoReflect.Descriptor instead.
func (*ActiveUserReq) Descriptor() ([]byte, []int) {
	return file_app_grpc_pb_confirmation_proto_rawDescGZIP(), []int{4}
}

func (x *ActiveUserReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_app_grpc_pb_confirmation_proto protoreflect.FileDescriptor

var file_app_grpc_pb_confirmation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x42, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xea,
	0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_grpc_pb_confirmation_proto_rawDescOnce sync.Once
	file_app_grpc_pb_confirmation_proto_rawDescData = file_app_grpc_pb_confirmation_proto_rawDesc
)

func file_app_grpc_pb_confirmation_proto_rawDescGZIP() []byte {
	file_app_grpc_pb_confirmation_proto_rawDescOnce.Do(func() {
		file_app_grpc_pb_confirmation_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_grpc_pb_confirmation_proto_rawDescData)
	})
	return file_app_grpc_pb_confirmation_proto_rawDescData
}

var file_app_grpc_pb_confirmation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_grpc_pb_confirmation_proto_goTypes = []interface{}{
	(*FindUserByEmailReq)(nil),  // 0: user.FindUserByEmailReq
	(*FindUserByEmailResp)(nil), // 1: user.FindUserByEmailResp
	(*CheckActivationReq)(nil),  // 2: user.CheckActivationReq
	(*CheckActivationResp)(nil), // 3: user.CheckActivationResp
	(*ActiveUserReq)(nil),       // 4: user.ActiveUserReq
	(*emptypb.Empty)(nil),       // 5: google.protobuf.Empty
}
var file_app_grpc_pb_confirmation_proto_depIdxs = []int32{
	0, // 0: user.ConfirmationService.FindUserInfoByEmail:input_type -> user.FindUserByEmailReq
	2, // 1: user.ConfirmationService.CheckActivation:input_type -> user.CheckActivationReq
	4, // 2: user.ConfirmationService.ActiveUser:input_type -> user.ActiveUserReq
	1, // 3: user.ConfirmationService.FindUserInfoByEmail:output_type -> user.FindUserByEmailResp
	3, // 4: user.ConfirmationService.CheckActivation:output_type -> user.CheckActivationResp
	5, // 5: user.ConfirmationService.ActiveUser:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_grpc_pb_confirmation_proto_init() }
func file_app_grpc_pb_confirmation_proto_init() {
	if File_app_grpc_pb_confirmation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_grpc_pb_confirmation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByEmailReq); i {
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
		file_app_grpc_pb_confirmation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByEmailResp); i {
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
		file_app_grpc_pb_confirmation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckActivationReq); i {
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
		file_app_grpc_pb_confirmation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckActivationResp); i {
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
		file_app_grpc_pb_confirmation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveUserReq); i {
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
			RawDescriptor: file_app_grpc_pb_confirmation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_grpc_pb_confirmation_proto_goTypes,
		DependencyIndexes: file_app_grpc_pb_confirmation_proto_depIdxs,
		MessageInfos:      file_app_grpc_pb_confirmation_proto_msgTypes,
	}.Build()
	File_app_grpc_pb_confirmation_proto = out.File
	file_app_grpc_pb_confirmation_proto_rawDesc = nil
	file_app_grpc_pb_confirmation_proto_goTypes = nil
	file_app_grpc_pb_confirmation_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConfirmationServiceClient is the client API for ConfirmationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfirmationServiceClient interface {
	FindUserInfoByEmail(ctx context.Context, in *FindUserByEmailReq, opts ...grpc.CallOption) (*FindUserByEmailResp, error)
	CheckActivation(ctx context.Context, in *CheckActivationReq, opts ...grpc.CallOption) (*CheckActivationResp, error)
	ActiveUser(ctx context.Context, in *ActiveUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type confirmationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfirmationServiceClient(cc grpc.ClientConnInterface) ConfirmationServiceClient {
	return &confirmationServiceClient{cc}
}

func (c *confirmationServiceClient) FindUserInfoByEmail(ctx context.Context, in *FindUserByEmailReq, opts ...grpc.CallOption) (*FindUserByEmailResp, error) {
	out := new(FindUserByEmailResp)
	err := c.cc.Invoke(ctx, "/user.ConfirmationService/FindUserInfoByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confirmationServiceClient) CheckActivation(ctx context.Context, in *CheckActivationReq, opts ...grpc.CallOption) (*CheckActivationResp, error) {
	out := new(CheckActivationResp)
	err := c.cc.Invoke(ctx, "/user.ConfirmationService/CheckActivation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confirmationServiceClient) ActiveUser(ctx context.Context, in *ActiveUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.ConfirmationService/ActiveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfirmationServiceServer is the server API for ConfirmationService service.
type ConfirmationServiceServer interface {
	FindUserInfoByEmail(context.Context, *FindUserByEmailReq) (*FindUserByEmailResp, error)
	CheckActivation(context.Context, *CheckActivationReq) (*CheckActivationResp, error)
	ActiveUser(context.Context, *ActiveUserReq) (*emptypb.Empty, error)
}

// UnimplementedConfirmationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConfirmationServiceServer struct {
}

func (*UnimplementedConfirmationServiceServer) FindUserInfoByEmail(context.Context, *FindUserByEmailReq) (*FindUserByEmailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserInfoByEmail not implemented")
}
func (*UnimplementedConfirmationServiceServer) CheckActivation(context.Context, *CheckActivationReq) (*CheckActivationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckActivation not implemented")
}
func (*UnimplementedConfirmationServiceServer) ActiveUser(context.Context, *ActiveUserReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActiveUser not implemented")
}

func RegisterConfirmationServiceServer(s *grpc.Server, srv ConfirmationServiceServer) {
	s.RegisterService(&_ConfirmationService_serviceDesc, srv)
}

func _ConfirmationService_FindUserInfoByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfirmationServiceServer).FindUserInfoByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ConfirmationService/FindUserInfoByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfirmationServiceServer).FindUserInfoByEmail(ctx, req.(*FindUserByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfirmationService_CheckActivation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckActivationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfirmationServiceServer).CheckActivation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ConfirmationService/CheckActivation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfirmationServiceServer).CheckActivation(ctx, req.(*CheckActivationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfirmationService_ActiveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActiveUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfirmationServiceServer).ActiveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ConfirmationService/ActiveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfirmationServiceServer).ActiveUser(ctx, req.(*ActiveUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfirmationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.ConfirmationService",
	HandlerType: (*ConfirmationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUserInfoByEmail",
			Handler:    _ConfirmationService_FindUserInfoByEmail_Handler,
		},
		{
			MethodName: "CheckActivation",
			Handler:    _ConfirmationService_CheckActivation_Handler,
		},
		{
			MethodName: "ActiveUser",
			Handler:    _ConfirmationService_ActiveUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/grpc/pb/confirmation.proto",
}
