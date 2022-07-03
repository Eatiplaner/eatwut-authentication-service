package client

import (
	"Eatiplan-Auth/app/grpc/rpc_pb"
	"context"
)

func (c *ClientService) FindUserByCredential(params *rpc_pb.FindUserRequest) (*rpc_pb.UserResponse, error) {
	client := rpc_pb.NewLoginSignupServiceClient(c.conn)

	return client.FindUserByCredential(context.Background(), params)
}

func (c *ClientService) CreateUser(params *rpc_pb.CreateRequest) (*rpc_pb.UserResponse, error) {
	client := rpc_pb.NewLoginSignupServiceClient(c.conn)

	return client.CreateUser(context.Background(), params)
}
