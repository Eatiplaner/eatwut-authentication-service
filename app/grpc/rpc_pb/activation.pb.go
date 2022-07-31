// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: app/grpc/pb/activation.proto

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

type ActivateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ActivateUserReq) Reset() {
	*x = ActivateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_pb_activation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateUserReq) ProtoMessage() {}

func (x *ActivateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_pb_activation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateUserReq.ProtoReflect.Descriptor instead.
func (*ActivateUserReq) Descriptor() ([]byte, []int) {
	return file_app_grpc_pb_activation_proto_rawDescGZIP(), []int{0}
}

func (x *ActivateUserReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RegenerateConfirmationByEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RegenerateConfirmationByEmailReq) Reset() {
	*x = RegenerateConfirmationByEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_pb_activation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegenerateConfirmationByEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegenerateConfirmationByEmailReq) ProtoMessage() {}

func (x *RegenerateConfirmationByEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_pb_activation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegenerateConfirmationByEmailReq.ProtoReflect.Descriptor instead.
func (*RegenerateConfirmationByEmailReq) Descriptor() ([]byte, []int) {
	return file_app_grpc_pb_activation_proto_rawDescGZIP(), []int{1}
}

func (x *RegenerateConfirmationByEmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_app_grpc_pb_activation_proto protoreflect.FileDescriptor

var file_app_grpc_pb_activation_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x20, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x32, 0xb7, 0x01, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x1d, 0x52,
	0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x13,
	0x5a, 0x11, 0x2e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_grpc_pb_activation_proto_rawDescOnce sync.Once
	file_app_grpc_pb_activation_proto_rawDescData = file_app_grpc_pb_activation_proto_rawDesc
)

func file_app_grpc_pb_activation_proto_rawDescGZIP() []byte {
	file_app_grpc_pb_activation_proto_rawDescOnce.Do(func() {
		file_app_grpc_pb_activation_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_grpc_pb_activation_proto_rawDescData)
	})
	return file_app_grpc_pb_activation_proto_rawDescData
}

var file_app_grpc_pb_activation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_grpc_pb_activation_proto_goTypes = []interface{}{
	(*ActivateUserReq)(nil),                  // 0: auth.ActivateUserReq
	(*RegenerateConfirmationByEmailReq)(nil), // 1: auth.RegenerateConfirmationByEmailReq
	(*emptypb.Empty)(nil),                    // 2: google.protobuf.Empty
}
var file_app_grpc_pb_activation_proto_depIdxs = []int32{
	0, // 0: auth.ActivationService.ActivateUser:input_type -> auth.ActivateUserReq
	1, // 1: auth.ActivationService.RegenerateConfirmationByEmail:input_type -> auth.RegenerateConfirmationByEmailReq
	2, // 2: auth.ActivationService.ActivateUser:output_type -> google.protobuf.Empty
	2, // 3: auth.ActivationService.RegenerateConfirmationByEmail:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_grpc_pb_activation_proto_init() }
func file_app_grpc_pb_activation_proto_init() {
	if File_app_grpc_pb_activation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_grpc_pb_activation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateUserReq); i {
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
		file_app_grpc_pb_activation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegenerateConfirmationByEmailReq); i {
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
			RawDescriptor: file_app_grpc_pb_activation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_grpc_pb_activation_proto_goTypes,
		DependencyIndexes: file_app_grpc_pb_activation_proto_depIdxs,
		MessageInfos:      file_app_grpc_pb_activation_proto_msgTypes,
	}.Build()
	File_app_grpc_pb_activation_proto = out.File
	file_app_grpc_pb_activation_proto_rawDesc = nil
	file_app_grpc_pb_activation_proto_goTypes = nil
	file_app_grpc_pb_activation_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ActivationServiceClient is the client API for ActivationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivationServiceClient interface {
	ActivateUser(ctx context.Context, in *ActivateUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RegenerateConfirmationByEmail(ctx context.Context, in *RegenerateConfirmationByEmailReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type activationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivationServiceClient(cc grpc.ClientConnInterface) ActivationServiceClient {
	return &activationServiceClient{cc}
}

func (c *activationServiceClient) ActivateUser(ctx context.Context, in *ActivateUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.ActivationService/ActivateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activationServiceClient) RegenerateConfirmationByEmail(ctx context.Context, in *RegenerateConfirmationByEmailReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.ActivationService/RegenerateConfirmationByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivationServiceServer is the server API for ActivationService service.
type ActivationServiceServer interface {
	ActivateUser(context.Context, *ActivateUserReq) (*emptypb.Empty, error)
	RegenerateConfirmationByEmail(context.Context, *RegenerateConfirmationByEmailReq) (*emptypb.Empty, error)
}

// UnimplementedActivationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedActivationServiceServer struct {
}

func (*UnimplementedActivationServiceServer) ActivateUser(context.Context, *ActivateUserReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateUser not implemented")
}
func (*UnimplementedActivationServiceServer) RegenerateConfirmationByEmail(context.Context, *RegenerateConfirmationByEmailReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegenerateConfirmationByEmail not implemented")
}

func RegisterActivationServiceServer(s *grpc.Server, srv ActivationServiceServer) {
	s.RegisterService(&_ActivationService_serviceDesc, srv)
}

func _ActivationService_ActivateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivationServiceServer).ActivateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.ActivationService/ActivateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivationServiceServer).ActivateUser(ctx, req.(*ActivateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivationService_RegenerateConfirmationByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegenerateConfirmationByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivationServiceServer).RegenerateConfirmationByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.ActivationService/RegenerateConfirmationByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivationServiceServer).RegenerateConfirmationByEmail(ctx, req.(*RegenerateConfirmationByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActivationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.ActivationService",
	HandlerType: (*ActivationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActivateUser",
			Handler:    _ActivationService_ActivateUser_Handler,
		},
		{
			MethodName: "RegenerateConfirmationByEmail",
			Handler:    _ActivationService_RegenerateConfirmationByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/grpc/pb/activation.proto",
}
