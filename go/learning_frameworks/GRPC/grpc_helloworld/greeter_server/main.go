//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto
// Package main implements a server for Greeter service.

package main

import (
	"context"
	"log"
	"net"

	"github.com/Danr17/GO_coding/GRPC/grpc_helloworld/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

//server is used to implement proto.GreeterServer
type server struct{}

//SayHello implements proto.GreeterServer
func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &proto.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello again " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
