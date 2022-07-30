package server

import (
	"Eatiplan-Auth/app/grpc/client"
	"Eatiplan-Auth/app/grpc/rpc_pb"
	pb "Eatiplan-Auth/app/grpc/rpc_pb"
	"context"
	"errors"
	"log"

	"Eatiplan-Auth/app/services"
)

type JwtServer struct {
	pb.JwtServiceServer
}

func (*JwtServer) ValidToken(ctx context.Context, req *pb.ValidRequest) (*pb.ValidResponse, error) {
	log.Println("Valid Token is processing...")
	log.Printf("Token: %s", req.GetToken())

	token := req.GetToken()
	error := services.TokenValid(token)

	// Check User is Active or not after valid token
	if error == nil {
		user_id, _ := services.ExtractUserIdFromToken(token)

		resp, err := client.Service.CheckActivation(&rpc_pb.CheckActivationReq{
			UserId: user_id,
		})

		if err != nil {
			error = err
		}

		if !resp.IsActive {
			error = errors.New("User has not been activate yet")
		}
	}

	resp := &pb.ValidResponse{
		Valid: error == nil,
	}

	log.Println("Valid Token processed Done")
	return resp, error
}
