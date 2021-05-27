package main

/**
实现服务端文件
*/
import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpcProLearn/services"
	"net"
	"runtime/debug"
)

// todo 添加过滤器的设置
func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("攔截器開始run。。。。。。 ")
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}
	jwtToken, ok := md["authorization"]
	var data []string
	err = json.Unmarshal([]byte(jwtToken[0]), &data)
	fmt.Println(" lanjie  error ", err)
	for _, v := range data {
		fmt.Println("包含了所需的用户名称 : ", v)
	}

	fmt.Println("所有的请求都会到这里来进行拦截： ", jwtToken, " is ok : ", ok)

	return handler(ctx, req)
}

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(filter))

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
