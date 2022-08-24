package server

import (
	"Eatwut-Auth/app/grpc/client"
	"Eatwut-Auth/app/grpc/rpc_pb"
	pb "Eatwut-Auth/app/grpc/rpc_pb"
	"Eatwut-Auth/app/kafka/procedures"
	"context"
	"errors"
	"log"

	"Eatwut-Auth/app/services"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ActivationServer struct {
	pb.ActivationServiceServer
}

func (*ActivationServer) ActivateUser(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	log.Println("Active User is processing...")
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

		if error == nil {
			_, err := client.Service.ActiveUser(ctx, token)

			if err != nil {
				error = err
			}
		}
	} else {
		error = errors.New("Authorization has not been set yet in metadata")
	}

	resp := &emptypb.Empty{}

	return resp, error
}

func (*ActivationServer) RegenerateConfirmationByEmail(ctx context.Context, req *pb.RegenerateConfirmationByEmailReq) (*emptypb.Empty, error) {
	log.Println("Regenerate Confirmation Email is processing...")
	log.Printf("Email: %s", req.GetEmail())

	email := req.GetEmail()

	resp, error := client.Service.FindUserInfoByEmail(&rpc_pb.FindUserByEmailReq{
		Email: email,
	})

	if error != nil {
		return &emptypb.Empty{}, error
	}

	user_id := resp.Id
	token, create_token_err := services.CreateToken(uint64(user_id))

	if create_token_err != nil {
		return &emptypb.Empty{}, create_token_err
	}

	check_activation_resp, check_activation_err := client.Service.CheckActivation(ctx, token.AccessToken)

	if check_activation_err != nil {
		return &emptypb.Empty{}, check_activation_err
	}

	if check_activation_resp.IsActive {
		error = errors.New("user has been activated before")
	} else {
		log.Printf("Sent Confirmation Notification To Email: %s", email)
		procedures.SendNotification(map[string]interface{}{
			"communication_type": "email",
			"data": map[string]string{
				"full_name":    resp.FullName,
				"to_email":     email,
				"token":        token.AccessToken,
				"template_key": "sign_up",
			},
		})
	}

	return &emptypb.Empty{}, error
}
