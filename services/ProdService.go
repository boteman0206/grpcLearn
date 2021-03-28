package HelloTest

import (
	"context"
)

type ProdService struct {
}

//实现sayHello方法

func (this *ProdService) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "调用grpc成功"}, nil
}
