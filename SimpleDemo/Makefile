stubs:
	protoc -I proto proto/Authentication.proto --go_out=plugins=grpc:proto

server:
	go run server/server.go

client:
	go run client/client.go
