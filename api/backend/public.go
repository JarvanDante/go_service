package backend

import "github.com/gogf/gf/v2/frame/g"

type PublicReq struct {
	g.Meta `path:"/login" method:"post" summary:"后台用户登录"`
}

type PublicRes struct {
}
