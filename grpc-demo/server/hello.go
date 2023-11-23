package server

import (
	"cafeteria/grpc-demo/proto"
	"context"
	"fmt"
)

type HelloServer struct {
	proto.UnimplementedGreeterServer
}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}

func (s *HelloServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: fmt.Sprintf("hello %s", req.Name)}, nil
}
