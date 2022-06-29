include .env
gen-proto:
	protoc app/grpc/pb/jwt.proto --go_out=plugins=grpc:.
sync-proto:
	curl -H "Authorization: token $(GH_PROTO_REPO_TOKEN)" \
		-H "Accept: application/vnd.github.v3.raw" \
		-o app/grpc/pb/jwt.proto \
		-L "https://api.github.com/repos/Eatiplan-Project/eatiplan-grpc-proto/contents/auth/jwt.proto?ref=main"
	protoc app/grpc/pb/jwt.proto --go_out=plugins=grpc:.
run-grpc-test-client:
	go run client-grpc.go
