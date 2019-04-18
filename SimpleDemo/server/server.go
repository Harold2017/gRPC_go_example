package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"

	pb "gRPC_go_example/SimpleDemo/proto"
)

var port = flag.Int("port", 10000, "the port to serve on")

type server struct{}

func clientIP(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("[getClinetIP] invoke FromContext() failed")
	}
	if pr.Addr == net.Addr(nil) {
		return "", fmt.Errorf("[getClientIP] peer.Addr is nil")
	}
	// ipPortSlice := strings.Split(pr.Addr.String(), ":")
	// return ipPortSlice[0], nil
	return pr.Addr.String(), nil
}

func getClientIP(ctx context.Context) string {
	ip, err := clientIP(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return ip
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.Response, error) {
	ip := getClientIP(ctx)
	log.Printf("Login service called from: %s", ip)
	return &pb.Response{Token: "asdasd", Time: 0.1, Status: 200}, nil
}

func (s *server) Signup(ctx context.Context, in *pb.SignupRequest) (*pb.Response, error) {
	ip := getClientIP(ctx)
	log.Printf("Signup service called from: %s", ip)
	return &pb.Response{Token: "asdasd", Time: 0.1, Status: 200}, nil
}

func (s *server) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryResponse, error) {
	ip := getClientIP(ctx)
	log.Printf("Query service called from: %s", ip)
	return &pb.QueryResponse{Res: "/asf/asd", Time: 0.1, Status: 200}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Server is listening at %v\n", lis.Addr())

	s := grpc.NewServer()

	pb.RegisterAuthenticationServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
