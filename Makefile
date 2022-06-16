gen-proto:
	protoc app/grpc/pb/jwt.proto --go_out=plugins=grpc:.
run-grpc-server:
	go run app/grpc/main.go
run-grpc-client:
	go run client-grpc.go
