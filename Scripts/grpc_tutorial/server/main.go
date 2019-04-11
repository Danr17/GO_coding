package main

import (
	"net"
	"context"

	"github.com/Danr17/GO_coding/Scripts/grpc_tutorial/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listnerer, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listnerer); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error){
	a, b := request.GetA(), request.GetB()

	result := a + b
	
	return &proto.Response{Resut: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Resut: result}, nil
}
