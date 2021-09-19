package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dainh-2247/grpc-gateway-demo/proto"
	"google.golang.org/grpc"
)

type server struct{}

// Echo
func (*server) Echo(ctx context.Context, req *proto.StringMessage) (*proto.StringMessage, error) {
	log.Printf("receive message %s\n", req.GetValue())
	return req, nil
}

// Sum
func (*server) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf("Sum is called...")

	resp := &proto.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("err while create listen  %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterMyServiceServer(s, &server{})

	fmt.Printf("server is running...")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
