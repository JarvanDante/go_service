package frontend

import (
	"context"
	"fmt"
	frontend2 "go-service/api/frontend"
	"go-service/internal/service/frontend"
)

type DemoController struct{}

func NewDemo() *DemoController {
	return &DemoController{}
}

func (c *DemoController) Index(ctx context.Context, req *frontend2.DemoReq) (res *frontend2.DemoRes, err error) {
	fmt.Println(frontend.Demo())
	msg, err := frontend.Demo().Index(ctx)
	if err != nil {
		return &frontend2.DemoRes{Message: ""}, nil
	}
	return &frontend2.DemoRes{Message: msg}, nil
}
