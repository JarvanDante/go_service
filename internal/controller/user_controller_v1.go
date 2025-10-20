package controller

import (
	"context"
	"fmt"
	"go-service/api/user/v1"
	"go-service/internal/service"
)

type UserControllerV1 struct{}

func NewUserV1() *UserControllerV1 {
	return &UserControllerV1{}
}

func (c *UserControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	fmt.Printf("service.User(): %#v\n", service.User())
	msg, _ := service.User().Login(ctx, req.Username, req.Password)

	return &v1.LoginRes{Message: msg}, nil
}
