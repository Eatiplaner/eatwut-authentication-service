package client

import (
	"Eatiplan-Auth/app/grpc"
	"Eatiplan-Auth/app/grpc/rpc_pb"
	"context"
)

func FindUserByCredential(params *rpc_pb.FindUserRequest) (*rpc_pb.UserResponse, error) {
	conn := grpc.InitGrpcClient()

	client := rpc_pb.NewLoginSignupServiceClient(conn)

	return client.FindUserByCredential(context.Background(), params)
}

func CreateUser(params *rpc_pb.CreateRequest) (*rpc_pb.UserResponse, error) {
	conn := grpc.InitGrpcClient()

	client := rpc_pb.NewLoginSignupServiceClient(conn)

	return client.CreateUser(context.Background(), params)
}
