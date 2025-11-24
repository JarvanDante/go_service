package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type AdminController struct{}

func NewAdminController() *AdminController {
	return &AdminController{}
}

//GetInfo
/**
 * showdoc
 * @catalog 后台/员工
 * @title 获取后台管理员信息
 * @description 获取后台管理员信息
 * @method get
 * @url /app/get-info
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"success","data":{"roles":["测试"],"name":"michael","avatar":"https://wpimg.wallstcn.com/577965b9-bb9e-4e02-9f0c-095b41417191","introduction":"","menus":[{"children":[{"children":null,"id":179,"name":"文档","path":"documentation/index","sort":0,"type":1}],"id":178,"name":"文档管理","path":"documentation","sort":0,"type":1},{"children":[{"children":null,"id":181,"name":"创建文章","path":"example/create","sort":0,"type":1},{"children":null,"id":182,"name":"编辑文章","path":"example/edit","sort":0,"type":1},{"children":null,"id":183,"name":"文章列表","path":"example/list","sort":0,"type":1}],"id":180,"name":"综合示例","path":"example","sort":0,"type":1}]}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @return_param data.roles array 角色组
 * @return_param data.name string 管理员名称
 * @return_param data.avatar string 管理员头像
 * @return_param data.introduction string 介绍
 * @return_param data.menus string 拥有权限
 * @remark 备注
 * @number 1
 */
func (c *AdminController) GetInfo(ctx context.Context, req *backendRoute.GetInfoReq) (res *response.Response, err error) {
	data, err := backend.ServiceAdmin().LGetInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	response.JsonOkCtx(ctx, data)

	return
}

//Logout
/**
 * showdoc
 * @catalog 后台/系统
 * @title 退出登录
 * @description 退出登录
 * @method post
 * @url /app/logout
 * @param token 必选 string 员工token
 * @return {"code":200,"status":1,"message":"退出成功","data":{}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @remark 备注
 * @number 1
 */
func (c *AdminController) Logout(ctx context.Context, req *backendRoute.LogoutReq) (res *response.Response, err error) {
	response.JsonOkCtx(ctx, nil, "退出成功")

	return
}

//Admins
/**
 * showdoc
 * @catalog 后台/员工
 * @title 员工列表
 * @description 员工列表
 * @method get
 * @url /app/admins
 * @param token 必选 string 员工token
 * @param status 必选 int 状态。1=启用；0=停用
 * @param username 必选 string 用户名
 * @param page 可选 string 页数
 * @param size 可选 string 数码
 * @return {"code":200,"status":1,"message":"获取数据成功","data":{"count":3,"list":[{"admin_role_id":1,"admin_role_name":"管理员","id":41,"last_login_ip":"","last_login_time":null,"nickname":"qwe789","status":1,"username":"qwe789"},{"admin_role_id":1,"admin_role_name":"管理员","id":29,"last_login_ip":"211.24.114.214","last_login_time":"2018-04-05 14:47:20","nickname":"pgcs","status":1,"username":"pgcs8"},{"admin_role_id":1,"admin_role_name":"管理员","id":28,"last_login_ip":"103.204.180.60","last_login_time":"2018-03-22 13:40:36","nickname":"cz2018","status":1,"username":"cz2018"}]}}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data string 主要数据
 * @return_param data.count int 总条数据
 * @return_param data.list.id int id
 * @return_param data.list.admin_role_id int 角色id
 * @return_param data.list.admin_role_name string 角色名称
 * @return_param data.list.last_login_ip string 最后登录
 * @return_param data.list.last_login_time string 最后登录时间
 * @return_param data.list.nickname string 昵称
 * @return_param data.list.status int 状态
 * @return_param data.list.username string 用户名
 * @remark 备注
 * @number 1
 */
func (c *AdminController) Admins(ctx context.Context, req *backendRoute.AdminsReq) (res *response.Response, err error) {

	data, err := backend.ServiceAdmin().LAdmins(ctx, req)
	if err != nil {
		response.JsonErrCtx(ctx, err.Error())
		return
	}
	response.JsonOkCtx(ctx, data, "获取数据成功")

	return
}

//CreateAdmin
/**
 * showdoc
 * @catalog 后台/员工
 * @title 添加员工
 * @description 添加员工
 * @method post
 * @url /app/create-admin
 * @param token 必选 string 员工token
 * @param username 必选 string 员工用户名
 * @param password 必选 string 员工密码
 * @param nickname 必选 string 员工昵称
 * @param role 必选 int 员工角色
 * @param status 必选 int 员工状态
 * @return {"code":200,"status":1,"message":"添加员工成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @remark 备注
 * @number 1
 */
func (c *AdminController) CreateAdmin(ctx context.Context, req *backendRoute.CreateAdminReq) (res *response.Response, err error) {

	err = backend.ServiceAdmin().LCreateAdmin(ctx, req)

	response.JsonOkCtx(ctx, nil, "添加员工成功")

	return
}

//UpdateAdmin
/**
 * showdoc
 * @catalog 后台/员工
 * @title 编辑员工
 * @description 编辑员工
 * @method post
 * @url /app/update-admin
 * @param token 必选 string 员工token
 * @param id 必选 int 员工ID
 * @param password 可选 string 员工密码
 * @param nickname 可选 string 员工昵称
 * @param role 可选 int 员工角色
 * @param status 可选 int 员工状态
 * @return {"code":200,"status":1,"message":"编辑员工成功","data":null}
 * @return_param code int 状态码
 * @return_param status int 成功/失败状态
 * @return_param message string 提示说明
 * @return_param data array 数组
 * @remark 备注
 * @number 1
 */
func (c *AdminController) UpdateAdmin(ctx context.Context, req *backendRoute.UpdateAdminReq) (res *response.Response, err error) {

	err = backend.ServiceAdmin().LUpdateAdmin(ctx, req)

	response.JsonOkCtx(ctx, nil, "编辑员工成功")

	return
}
