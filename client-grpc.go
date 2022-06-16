package main

import (
	"context"
	"log"

	pb "Eatiplan-Auth/app/grpc/rpc_pb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	defer cc.Close()

	client := pb.NewJwtServiceClient(cc)

	log.Printf("service client %f", client)

	_, err = client.ValidToken(context.Background(), &pb.ValidRequest{
		Token: "token",
	})

	if err != nil {
		log.Fatalf("call valid request err %v", err)
	}

	log.Println("Token is valid")
}
