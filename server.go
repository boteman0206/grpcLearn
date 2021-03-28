package main

/**
实现服务端文件
*/
import (
	"fmt"
	"google.golang.org/grpc"
	"grpcProLearn/services"
	"net"
)

func main() {
	server := grpc.NewServer()

	HelloTest.RegisterHelloServiceServer(server, new(HelloTest.ProdService))

	listener, e := net.Listen("tcp", ":8082")

	if e != nil {
		fmt.Println("is list error", e.Error())
		return
	}

	server.Serve(listener)
}
