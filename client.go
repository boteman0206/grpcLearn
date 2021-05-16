/**
客户端调用代码
*/

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcProLearn/services"
)

func main() {

	//添加参数 grpc.WithInsecure() 会报错
	conn, e := grpc.Dial(":8082", grpc.WithInsecure())
	if e != nil {
		fmt.Println("client error ", e.Error())
		return
	}

	defer conn.Close()

	//客户端调用
	client := HelloTest.NewHelloServiceClient(conn)

	//调用方法sqyHello
	response, i := client.SayHello(context.Background(),
		&HelloTest.HelloRequest{Username: "i am jack"})

	if i != nil {
		fmt.Println("client sayHello error: ", i.Error())
	}
	fmt.Println("client response :", response.GetMessage())

	// 调用demo1服务
	server1Client := HelloTest.NewTestServer1Client(conn)
	// 初始化oneOf的值
	req := HelloTest.Req{
		Name: "jack",
		Age:  19,
		Type: &HelloTest.Req_Addr{
			Addr: "上海市中心",
		},
	}
	demo1, e := server1Client.TestDemo1(context.Background(), &req)
	fmt.Println("demo1 的 error ：", e, " demo1 的 res： ", demo1)
}
