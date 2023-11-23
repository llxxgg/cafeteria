package main

import (
	"cafeteria/grpc-demo/proto"
	"cafeteria/grpc-demo/server"
	"google.golang.org/grpc"
	"net"
)

func main() {
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, server.NewHelloServer())
	//reflection.Register(s)

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
