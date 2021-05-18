package HelloTest

import (
	"context"
	"fmt"
)

type ProdService struct {
}

//实现sayHello方法

func (this *ProdService) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "调用grpc成功"}, nil
}

func (this *ProdService) MyIntTest(ctx context.Context, req *ReqMy) (*HelloResponse, error) {

	fmt.Println("获取的myintTest的参数：", req, req.N1, req.N2)
	return &HelloResponse{Message: "测试int32和int64参数的使用。。。"}, nil
}
