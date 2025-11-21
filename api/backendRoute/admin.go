package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type GetInfoReq struct {
	g.Meta `path:"/get-info" method:"get" summary:"获取后台管理员信息"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" summary:"后台退出登录"`
}

type AdminsReq struct {
	g.Meta   `path:"/admins" method:"get" summary:"获取员工列表"`
	Status   *int    `json:"status" dc:"状态"`
	Username *string `json:"username" dc:"用户名"`
	Page     *int    `json:"page" d:"1" dc:"页数"`
	Size     *int    `json:"size" d:"20" dc:"页码"`
}

type CreateAdminReq struct {
	g.Meta   `path:"/create-admin" method:"post" summary:"添加员工"`
	Username string `json:"username" v:"required#用户名必填" dc:"用户名"`
	Password string `json:"password" v:"required#密码必填" dc:"密码"`
	Nickname string `json:"nickname" v:"required#昵称必填" dc:"昵称"`
	Role     int    `json:"role" v:"required#角色必填" dc:"角色"`
	Status   int    `json:"status" v:"required#状态必填" dc:"状态"`
}

type UpdateAdminReq struct {
	g.Meta   `path:"/update-admin" method:"post" summary:"编辑员工"`
	Id       int    `json:"id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Role     int    `json:"role"`
	Status   int    `json:"status"`
}

type DeleteAdminReq struct {
	g.Meta `path:"/delete-admin" method:"post" summary:"删除员工"`
	Id     int `json:"id"`
}
