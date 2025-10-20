package logic

import (
	"context"
	"fmt"
	"go-service/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(&sUser{})
}

func (s *sUser) Login(ctx context.Context, username, password string) (string, error) {
	// 模拟验证逻辑
	if username == "admin" && password == "123456" {
		return fmt.Sprintf("欢迎回来，%s！", username), nil
	}
	return "用户名或密码错误", nil
}
