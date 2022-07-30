package client

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientService struct {
	conn func() *grpc.ClientConn
}

func New() ClientService {
	return ClientService{
		conn: InitGrpcClient,
	}
}

func InitGrpcClient() *grpc.ClientConn {
	grpc_host := fmt.Sprintf("%s:%s", os.Getenv("GRPC_CLIENT_DOMAIN"), os.Getenv("GRPC_CLIENT_PORT"))

	conn, err := grpc.Dial(grpc_host, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("err while call grpc %v", err)
	}

	return conn
}

var Service = New()
