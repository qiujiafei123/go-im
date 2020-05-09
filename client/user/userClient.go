package main

import (
	"context"
	"fmt"
	pb "go-im/service/user/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:7000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Login(ctx, &pb.LoginRequest{Name: "xiaopeng", Password:"123456"})
	if err != nil {
		log.Fatalf("err resp %s", err)
	}
	fmt.Printf("%+v \n", r)
}
