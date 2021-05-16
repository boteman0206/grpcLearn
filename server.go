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

	// 注册prod.proto服务
	HelloTest.RegisterHelloServiceServer(server, new(HelloTest.ProdService))
	// 注册demo1.proto服务
	HelloTest.RegisterTestServer1Server(server, new(HelloTest.Demo1Test))

	listener, e := net.Listen("tcp", ":8082")

	if e != nil {
		fmt.Println("is list error", e.Error())
		return
	}

	server.Serve(listener)
}
