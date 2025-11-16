package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type GetInfoReq struct {
	g.Meta `path:"/get-info" method:"get" summary:"获取后台管理员信息"`
}
