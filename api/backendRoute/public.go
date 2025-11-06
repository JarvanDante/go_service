package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type PublicReq struct {
	g.Meta   `path:"/login" method:"post" summary:"后台用户登录"`
	Username string `json:"username" v:"required#用户名必填" dc:"用户名"`
	Password string `json:"password" v:"required#密码必填" dc:"密码"`
	Code     string `json:"code" v:"required#动态验证码必填" dc:"动态验证码"`
}
