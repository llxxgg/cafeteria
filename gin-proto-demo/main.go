package main

import (
	"cafeteria/gin-proto-demo/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	proto.RegisterGreeterHTTPServer(e, &Server{})
	e.Run()
}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: fmt.Sprintf("hello, %s", req.Name)}, nil
}
