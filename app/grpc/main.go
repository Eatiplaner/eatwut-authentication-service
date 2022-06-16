package main

import (
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "Eatiplan-Auth/app/grpc/rpc_pb"
	"Eatiplan-Auth/app/grpc/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterJwtServiceServer(s, &server.JwtServer{})

	fmt.Println("gRPC Server listening on port 8080")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
