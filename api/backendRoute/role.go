package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type RolesReq struct {
	g.Meta `path:"/roles" method:"get" summary:"获取站点职务列表"`
}

type PermissionsReq struct {
	g.Meta `path:"/permissions" method:"get" summary:"获取站点职务权限列表"`
}

type CreateReq struct {
	g.Meta `path:"/create-role" method:"post" summary:"添加职务"`
	Name   string `json:"Name" v:"required#角色名必填" dc:"角色名"`
}
