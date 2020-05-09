package main

import (
	"context"
	"fmt"
	pb "go-im/service/user/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func (user *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("这个是 grpc 调动 login")
	return &pb.LoginResponse{Code:0, Authtoken:"xxxxxxxx"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("faild to listen: %s \n", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &UserService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild to serve %s", err)
	}
}


