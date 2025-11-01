package frontend

import (
	frontend2 "go-service/api/frontend"
	"go-service/internal/service/frontend"
	"golang.org/x/net/context"
)

type DemoController struct{}

func NewDemo() *DemoController {
	return &DemoController{}
}

func (c *DemoController) Index(ctx context.Context, req *frontend2.DemoReq) (res *frontend2.DemoRes, err error) {

	msg, err := frontend.Demo().Index(ctx)
	if err != nil {
		return &frontend2.DemoRes{Message: ""}, nil
	}
	return &frontend2.DemoRes{Message: msg}, nil
}
