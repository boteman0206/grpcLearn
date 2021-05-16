package HelloTest

import (
	"context"
	"fmt"
)

type Demo1Test struct {
}

func (c *Demo1Test) TestDemo1(ctx context.Context, req *Req) (*Response, error) {

	age := req.Age
	name := req.GetName()
	req_type := req.GetPhone()
	req_addr := req.GetAddr()

	fmt.Println("demo1 :", age, name, req_type, req_addr)
	return &Response{Msg: "testDemo1的方法调用成功."}, nil
}
