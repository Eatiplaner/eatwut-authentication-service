package client

import (
	"Eatiplan-Auth/app/grpc/rpc_pb"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

func (c *ClientService) CheckActivation(params *rpc_pb.CheckActivationReq) (*rpc_pb.CheckActivationResp, error) {
	client := rpc_pb.NewConfirmationServiceClient(c.conn())

	return client.CheckActivation(context.Background(), params)
}

func (c *ClientService) ActiveUser(params *rpc_pb.ActiveUserReq) (*empty.Empty, error) {
	client := rpc_pb.NewConfirmationServiceClient(c.conn())

	return client.ActiveUser(context.Background(), params)
}

func (c *ClientService) FindUserInfoByEmail(params *rpc_pb.FindUserByEmailReq) (*rpc_pb.FindUserByEmailResp, error) {
	client := rpc_pb.NewConfirmationServiceClient(c.conn())

	return client.FindUserInfoByEmail(context.Background(), params)
}
