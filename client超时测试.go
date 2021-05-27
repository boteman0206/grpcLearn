/**
客户端调用代码
*/

package main

import (
	"context"
	"fmt"
	"time"

	//kit "github.com/tricobbler/rp-kit"
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

	clientDeadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	test, e := client.TimeOutTest(ctx, &HelloTest.HelloRequest{
		Username: "vivo",
	})

	fmt.Println("超时测试test： ", test, e)
}
