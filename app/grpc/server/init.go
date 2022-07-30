package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "Eatiplan-Auth/app/grpc/rpc_pb"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func InitGrpcServer() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	go func() {
		grpc_port := os.Getenv("GRPC_SERVER_PORT")

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpc_port))

		if err != nil {
			log.Fatalf("err while create listen %v", err)
		}

		s := grpc.NewServer()
		pb.RegisterJwtServiceServer(s, &JwtServer{})
		pb.RegisterActivationServiceServer(s, &ActivationServer{})

		// Register reflection service on gRPC server.
		reflection.Register(s)

		fmt.Println()
		fmt.Printf("gRPC Server listening on port %s", grpc_port)

		err = s.Serve(lis)

		if err != nil {
			log.Fatalf("err while serve %v", err)
		}
	}()

}

func InitGrpcClient() *grpc.ClientConn {
	grpc_host := fmt.Sprintf("%s:%s", os.Getenv("GRPC_CLIENT_DOMAIN"), os.Getenv("GRPC_CLIENT_PORT"))

	conn, err := grpc.Dial(grpc_host, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("err while call grpc %v", err)
	}

	return conn
}
