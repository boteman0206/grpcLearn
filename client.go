/**
客户端调用代码
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/metadata"

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

	// 测试参数int32和int64区别
	reqmy := HelloTest.ReqMy{
		//N1: int64(91),
		//N2: int32(12),
	}

	// rpc 无法传递数据貌似
	background := context.WithValue(context.Background(), "name", "jack-data")
	test, e := client.MyIntTest(background, &reqmy)
	fmt.Println("test : ", test.Message, e)

	//grpc设置参数
	fmt.Println("+==============grpc设置参数=============")
	//grpcContext := struct {
	//	Name string
	//	Age int
	//}{
	//	Name: "jack",
	//	Age: 12,
	//}
	//marshal, e := json.Marshal(grpcContext)
	//
	//ctx := metadata.AppendToOutgoingContext(context.Background(), "grpc_context", string(marshal))

	// later, add some more metadata to the context (e.g. in an interceptor)
	//md, _ := metadata.FromOutgoingContext(ctx)
	//newMD := metadata.Pairs("k3", "v3")
	//ctx = metadata.NewContext(ctx, metadata.Join(metadata.New(send), newMD))

	/**
	todo 这种方式添加参数是可以通过  metadata.FromIncomingContext(ctx)来获取context里面的参数的
		会将 k1的value合并 获取的参数的数据信息： k1 = [v1 v2],  k2 = [v3]
	*/
	md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	test1, e := client.MyIntTest(ctx, &reqmy)
	fmt.Println("二次测试： ", test1.Message, e)

	fmt.Println("+=============测试结构体参数的传递==============")

	grpcContext := struct {
		Name string
		Age  int
	}{
		Name: "jack",
		Age:  12,
	}
	marshal, e := json.Marshal(grpcContext)

	ctx1 := metadata.AppendToOutgoingContext(context.Background(), "grpc_context", string(marshal))
	intTest, e := client.MyIntTest(ctx1, &reqmy)

	fmt.Println("++++++++++++++++", intTest, e)
}
