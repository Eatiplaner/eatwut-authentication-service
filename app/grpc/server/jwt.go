package server

import (
	"Eatwut-Auth/app/grpc/client"
	pb "Eatwut-Auth/app/grpc/rpc_pb"
	"context"
	"errors"
	"log"

	"Eatwut-Auth/app/services"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type JwtServer struct {
	pb.JwtServiceServer
}

func (*JwtServer) ValidActivationToken(ctx context.Context, req *emptypb.Empty) (*pb.ValidResponse, error) {
	log.Println("Valid Token is processing...")
	var values []string
	var error error

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values = md.Get("authorization")
	}

	if len(values) > 0 {
		token := values[0]
		log.Printf("Token: %s", token)
		error = services.TokenValid(token)

		// Check User is Active or not after valid token
		if error == nil {
			resp, err := client.Service.CheckActivation(ctx, token)

			if err != nil {
				error = err
			} else {
				if !resp.IsActive {
					error = errors.New("User has not been activate yet")
				}
			}
		}
	} else {
		error = errors.New("Authorization has not been set yet in metadata")
	}

	resp := &pb.ValidResponse{
		Valid: error == nil,
	}

	log.Println("Valid Token processed Done")
	return resp, error
}

func (*JwtServer) ValidToken(ctx context.Context, req *emptypb.Empty) (*pb.ValidResponse, error) {
	log.Println("Valid Token is processing...")
	var values []string
	var error error

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values = md.Get("authorization")
	}

	if len(values) > 0 {
		token := values[0]
		log.Printf("Token: %s", token)
		error = services.TokenValid(token)
	} else {
		error = errors.New("Authorization has not been set yet in metadata")
	}

	resp := &pb.ValidResponse{
		Valid: error == nil,
	}

	log.Println("Valid Token processed Done")

	return resp, error
}
