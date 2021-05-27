/**
客户端调用代码
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/metadata"
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

	data := []string{"jack", "bob", "lucy"}
	marshal, e := json.Marshal(data)

	//todo、 添加过滤器的设置
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", string(marshal))

	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

	test, e := client.TimeOutTest(ctx, &HelloTest.HelloRequest{
		Username: "vivo",
	})

	fmt.Println("超时测试test： ", test, e)
}
