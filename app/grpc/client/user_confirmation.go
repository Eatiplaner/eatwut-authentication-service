package client

import (
	"Eatiplan-Auth/app/grpc/rpc_pb"
	"context"
)

func (c *ClientService) CheckActivation(params *rpc_pb.CheckActivationReq) (*rpc_pb.CheckActivationResp, error) {
	client := rpc_pb.NewConfirmationServiceClient(c.conn)

	return client.CheckActivation(context.Background(), params)
}

func (c *ClientService) UpdateProfile(params *rpc_pb.UpdateProfileRequest) (*rpc_pb.UpdateProfileResponse, error) {
	client := rpc_pb.NewProfileServiceClient(c.conn)

	return client.UpdateProfile(context.Background(), params)
}
