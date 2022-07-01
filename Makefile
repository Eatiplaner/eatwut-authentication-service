gen-proto:
	protoc app/grpc/pb/jwt.proto --go_out=plugins=grpc:.
run-grpc-test-client:
	go run client-grpc.go
