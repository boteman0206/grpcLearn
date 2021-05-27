package HelloTest

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/metadata"
	"strings"
)

type ProdService struct {
}

//实现sayHello方法

func (this *ProdService) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "调用grpc成功"}, nil
}

func (this *ProdService) MyIntTest(ctx context.Context, req *ReqMy) (*HelloResponse, error) {
	/**
	todo 这种方式获取不到，可能是传参问题
	var header, trailer metadata.MD
	grpc.Header(&header)
	grpc.Trailer(&trailer)
	fmt.Println("header: ", header)
	fmt.Println("trailer: ", trailer )
	*/

	/**
	todo 测试使用context来传递参数,也是无法获取参数数据的
	fmt.Println("获取context 的参数 ： ",ctx,  ctx.Value("name"))
	fmt.Println("获取的myintTest的参数：", req, req.N1, req.N2)
	*/

	// todo 这个就可以获取grpc请求头的参数信息
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println("获取的md： ", md, ok)
	k1 := md.Get("k1")
	k2 := md.Get("k2")
	fmt.Println("获取的参数的数据信息： ", k1, k2)

	// todo 传递过来的结构体也是可以获取的
	get := md.Get("grpc_context") // todo 返回的是数组可以直接get[0]来获取
	var data map[string]interface{}
	join := strings.Join(get, "")
	err := json.Unmarshal([]byte(join), &data)
	fmt.Println("获取的结构体的数据： ", err, data, data["Name"], data["Age"])
	return &HelloResponse{Message: "测试int32和int64参数的使用。。。"}, nil

}
