package server

import (
	pb "Eatiplan-Auth/app/grpc/rpc_pb"
	"context"
	"log"

	"Eatiplan-Auth/app/services"
)

type JwtServer struct {
	pb.JwtServiceServer
}

func (*JwtServer) ValidToken(ctx context.Context, req *pb.ValidRequest) (*pb.ValidResponse, error) {
	log.Println("Valid Token is processing...")
	log.Printf("Token: %s", req.GetToken())
	error := services.TokenValid(req.GetToken())
	resp := &pb.ValidResponse{
		Valid: error == nil,
	}

	log.Println("Valid Token processed Done")
	return resp, error
}
