package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type PublicController struct{}

func NewPublicController() *PublicController {
	return &PublicController{}
}

//Login
/**
 * showdoc
 * @catalog 后台/员工
 * @title 登录
 * @description 员工登录的接口
 * @method post
 * @url /app/login
 * @param username 必选 string 用户名
 * @param password 必选 string 密码
 * @param code 可选 string 动态验证码
 * @return {"code":200,"status":1,"message":"success","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHAiOiJqaCIsImV4cCI6MTc2MjY4NDIwMSwiaWQiOjI3LCJzaXRlX2lkIjoxLCJ1c2VybmFtZSI6Im1pY2hhZWwifQ.VUM-dM4--yCEvvJQZHdsL-eX6SbaN5MAUjq_xUy6cuI"}}
 * @return_param code int 状态码
 * @return_param data string 主要数据
 * @return_param message string 提示说明
 * @return_param status int 成功/失败状态
 * @remark 备注
 * @number 1
 */
func (c *PublicController) Login(ctx context.Context, req *backendRoute.PublicReq) (res *response.Response, err error) {
	token, err := backend.ServicePublic().LLogin(ctx, req)
	if err != nil {
		return nil, err
	}
	response.JsonOkCtx(ctx, g.Map{"token": token})
	return
}
