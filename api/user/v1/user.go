package v1

import "github.com/gogf/gf/v2/frame/g"

// 定义请求结构体
type LoginReq struct {
	g.Meta   `path:"/user/login" method:"post" tags:"User" summary:"用户登录"`
	Username string `json:"username" v:"required#用户名不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

// 定义响应结构体
type LoginRes struct {
	Message string `json:"message"`
}
