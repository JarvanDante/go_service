package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
 * @desc：获取后台管理员信息
 * @param ctx
 * @param req
 * @return res
 * @return err
 * @author : Carson
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
 * @desc：退出登录
 * @param ctx
 * @param req
 * @return res
 * @return err
 * @author : Carson
 */
func (c *AdminController) Logout(ctx context.Context, req *backendRoute.LogoutReq) (res *response.Response, err error) {
	response.JsonOkCtx(ctx, g.Map{}, "退出成功")

	return
}
