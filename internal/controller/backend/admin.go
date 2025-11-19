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
