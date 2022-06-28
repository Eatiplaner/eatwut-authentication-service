// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: app/grpc/pb/jwt.proto

package rpc_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ValidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidRequest) Reset() {
	*x = ValidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_pb_jwt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidRequest) ProtoMessage() {}

func (x *ValidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_pb_jwt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidRequest.ProtoReflect.Descriptor instead.
func (*ValidRequest) Descriptor() ([]byte, []int) {
	return file_app_grpc_pb_jwt_proto_rawDescGZIP(), []int{0}
}

func (x *ValidRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *ValidResponse) Reset() {
	*x = ValidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_pb_jwt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidResponse) ProtoMessage() {}

func (x *ValidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_pb_jwt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidResponse.ProtoReflect.Descriptor instead.
func (*ValidResponse) Descriptor() ([]byte, []int) {
	return file_app_grpc_pb_jwt_proto_rawDescGZIP(), []int{1}
}

func (x *ValidResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_app_grpc_pb_jwt_proto protoreflect.FileDescriptor

var file_app_grpc_pb_jwt_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x77,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x24, 0x0a,
	0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x25, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x32, 0x45, 0x0a, 0x0a, 0x4a, 0x77,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_grpc_pb_jwt_proto_rawDescOnce sync.Once
	file_app_grpc_pb_jwt_proto_rawDescData = file_app_grpc_pb_jwt_proto_rawDesc
)

func file_app_grpc_pb_jwt_proto_rawDescGZIP() []byte {
	file_app_grpc_pb_jwt_proto_rawDescOnce.Do(func() {
		file_app_grpc_pb_jwt_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_grpc_pb_jwt_proto_rawDescData)
	})
	return file_app_grpc_pb_jwt_proto_rawDescData
}

var file_app_grpc_pb_jwt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_grpc_pb_jwt_proto_goTypes = []interface{}{
	(*ValidRequest)(nil),  // 0: auth.ValidRequest
	(*ValidResponse)(nil), // 1: auth.ValidResponse
}
var file_app_grpc_pb_jwt_proto_depIdxs = []int32{
	0, // 0: auth.JwtService.ValidToken:input_type -> auth.ValidRequest
	1, // 1: auth.JwtService.ValidToken:output_type -> auth.ValidResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_grpc_pb_jwt_proto_init() }
func file_app_grpc_pb_jwt_proto_init() {
	if File_app_grpc_pb_jwt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_grpc_pb_jwt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidRequest); i {
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
		file_app_grpc_pb_jwt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidResponse); i {
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
			RawDescriptor: file_app_grpc_pb_jwt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_grpc_pb_jwt_proto_goTypes,
		DependencyIndexes: file_app_grpc_pb_jwt_proto_depIdxs,
		MessageInfos:      file_app_grpc_pb_jwt_proto_msgTypes,
	}.Build()
	File_app_grpc_pb_jwt_proto = out.File
	file_app_grpc_pb_jwt_proto_rawDesc = nil
	file_app_grpc_pb_jwt_proto_goTypes = nil
	file_app_grpc_pb_jwt_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JwtServiceClient is the client API for JwtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JwtServiceClient interface {
	ValidToken(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error)
}

type jwtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJwtServiceClient(cc grpc.ClientConnInterface) JwtServiceClient {
	return &jwtServiceClient{cc}
}

func (c *jwtServiceClient) ValidToken(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/auth.JwtService/ValidToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JwtServiceServer is the server API for JwtService service.
type JwtServiceServer interface {
	ValidToken(context.Context, *ValidRequest) (*ValidResponse, error)
}

// UnimplementedJwtServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJwtServiceServer struct {
}

func (*UnimplementedJwtServiceServer) ValidToken(context.Context, *ValidRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidToken not implemented")
}

func RegisterJwtServiceServer(s *grpc.Server, srv JwtServiceServer) {
	s.RegisterService(&_JwtService_serviceDesc, srv)
}

func _JwtService_ValidToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).ValidToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.JwtService/ValidToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).ValidToken(ctx, req.(*ValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JwtService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.JwtService",
	HandlerType: (*JwtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidToken",
			Handler:    _JwtService_ValidToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/grpc/pb/jwt.proto",
}
