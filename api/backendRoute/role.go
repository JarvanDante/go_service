package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/roles" method:"get" summary:"获取站点职务列表"`
}
