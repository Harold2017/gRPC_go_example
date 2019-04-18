stubs:
	protoc -I proto proto/Greeter.proto --go_out=plugins=grpc:proto

server:
	go run server/server.go

client:
	go run client/client.go
