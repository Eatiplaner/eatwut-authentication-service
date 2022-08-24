package client

import (
	"Eatwut-Auth/app/grpc/rpc_pb"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/metadata"
)

func (c *ClientService) CheckActivation(ctx context.Context, token string) (*rpc_pb.CheckActivationResp, error) {
	client := rpc_pb.NewConfirmationServiceClient(c.conn())
	ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", token)

	return client.CheckActivation(ctx, &empty.Empty{})
}

func (c *ClientService) ActiveUser(ctx context.Context, token string) (*empty.Empty, error) {
	client := rpc_pb.NewConfirmationServiceClient(c.conn())
	ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", token)

	return client.ActiveUser(ctx, &empty.Empty{})
}

func (c *ClientService) FindUserInfoByEmail(params *rpc_pb.FindUserByEmailReq) (*rpc_pb.FindUserByEmailResp, error) {
	client := rpc_pb.NewConfirmationServiceClient(c.conn())

	return client.FindUserInfoByEmail(context.Background(), params)
}
