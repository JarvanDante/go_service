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
	Username string `json:"username" v:"required|length:4,12#请输入用户名|用户名长度4-12" dc:"用户名"`
	Password string `json:"password" v:"required|password2#请输入密码|密码包含大小写字母和数字并且长度在6~18的" dc:"密码"`
	Nickname string `json:"nickname" v:"required#昵称必填" dc:"昵称"`
	Role     int    `json:"role" v:"required#角色必填" dc:"角色"`
	Status   int    `json:"status" v:"required#状态必填" dc:"状态"`
}

type UpdateAdminReq struct {
	g.Meta   `path:"/update-admin" method:"post" summary:"编辑员工"`
	Id       int    `json:"id" v:"required#请输入员工ID" dc:"员工ID"`
	Password string `json:"password" v:"password2#密码必须由字母和数字组成并且长度在6~18" dc:"密码"`
	Nickname string `json:"nickname" v:"length:2,20#昵称长度2-20" dc:"昵称"`
	Role     int    `json:"role" dc:"角色"`
	Status   int    `json:"status" dc:"状态"`
}

type DeleteAdminReq struct {
	g.Meta `path:"/delete-admin" method:"post" summary:"删除员工"`
	Id     int `json:"id" v:"required#请输入员工ID" dc:"员工ID"`
}

type MenusReq struct {
	g.Meta `path:"/menus" method:"get" summary:"获取菜单列表"`
}
