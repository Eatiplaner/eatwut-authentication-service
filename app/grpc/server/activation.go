package server

import (
	"Eatiplan-Auth/app/grpc/client"
	"Eatiplan-Auth/app/grpc/rpc_pb"
	pb "Eatiplan-Auth/app/grpc/rpc_pb"
	"Eatiplan-Auth/app/kafka/procedures"
	"context"
	"errors"
	"log"

	"Eatiplan-Auth/app/services"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ActivationServer struct {
	pb.ActivationServiceServer
}

func (*ActivationServer) ActivateUser(ctx context.Context, req *pb.ActivateUserReq) (*emptypb.Empty, error) {
	log.Println("Active User is processing...")
	log.Printf("Token: %s", req.GetToken())

	token := req.GetToken()
	error := services.TokenValid(token)

	if error == nil {
		_, err := client.Service.ActiveUser(&rpc_pb.ActiveUserReq{
			Token: token,
		})

		if err != nil {
			error = err
		}
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

	check_activation_resp, check_activation_err := client.Service.CheckActivation(&rpc_pb.CheckActivationReq{
		Token: token.AccessToken,
	})

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
