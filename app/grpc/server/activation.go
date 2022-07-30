package server

import (
	"Eatiplan-Auth/app/grpc/client"
	pb "Eatiplan-Auth/app/grpc/rpc_pb"
	"context"
	"log"

	"Eatiplan-Auth/app/services"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ActivationServer struct {
	pb.ActivationServiceServer
}

func (*ActivationServer) ActiveUser(ctx context.Context, req *pb.ActiveUserReq) (*emptypb.Empty, error) {
	log.Println("Active User is processing...")
	log.Printf("Token: %s", req.GetToken())

	token := req.GetToken()
	error := services.TokenValid(token)

	if error == nil {
		_, err := client.Service.UpdateProfile(&pb.UpdateProfileRequest{
			Data: &pb.UpdateProfileData{IsActiveOneof: &pb.UpdateProfileData_IsActive{IsActive: true}},
		})

		if err != nil {
			error = err
		}
	}

	resp := &emptypb.Empty{}

	return resp, error
}
