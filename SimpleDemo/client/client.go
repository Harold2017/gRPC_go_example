package main

import (
	"context"
	"fmt"
	"log"
	"os"
	// "reflect"
	"time"

	pb "gRPC_go_example/SimpleDemo/proto"
	"google.golang.org/grpc"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

const (
	address        = "localhost: 10000"
	defaultService = "list"
)

func listServices(conn *grpc.ClientConn) string {
	rc := rpb.NewServerReflectionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, _ := rc.ServerReflectionInfo(ctx)
	req := &rpb.ServerReflectionRequest{MessageRequest: &rpb.ServerReflectionRequest_ListServices{ListServices: "*"}}
	errreq := r.Send(req)
	if errreq == nil {
		resp, errresp := r.Recv()
		if errresp == nil {
			return fmt.Sprint(resp.MessageResponse)
		}
	}
	return fmt.Sprint(errreq)
}

func signup(conn *grpc.ClientConn) string {
	c := pb.NewAuthenticationClient(conn)

	email := "test@test.com"
	username := "test"
	password := "test"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Signup(ctx, &pb.SignupRequest{Email: email, Username: username, Password: password})
	if err != nil {
		log.Fatalf("can not signup: %v", err)
	}
	return fmt.Sprint(r)
}

func login(conn *grpc.ClientConn) string {
	c := pb.NewAuthenticationClient(conn)

	username := "test"
	password := "test"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Login(ctx, &pb.LoginRequest{Username: username, Password: password})
	if err != nil {
		log.Fatalf("can not login: %v", err)
	}
	return fmt.Sprint(r)
}

func query(conn *grpc.ClientConn) string {
	c := pb.NewAuthenticationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Query(ctx, &pb.QueryRequest{Api: "asd"})
	if err != nil {
		log.Fatalf("can not query: %v", err)
	}
	return fmt.Sprint(r)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	service := defaultService
	if len(os.Args) > 1 {
		service = os.Args[1]
	}

	resp := ""

	switch service {
	case defaultService:
		resp = listServices(conn)
		log.Println(resp)
	case "login":
		resp = login(conn)
		log.Println(resp)
	case "signup":
		resp = signup(conn)
		log.Println(resp)
	case "query":
		resp = query(conn)
		log.Println(resp)
	default:
		log.Fatalln("invalid services")
	}
}
