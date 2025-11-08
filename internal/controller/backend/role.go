package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type RoleController struct{}

func NewRoleController() *RoleController {
	return &RoleController{}
}

func (c *RoleController) Index(ctx context.Context, req *backendRoute.RoleReq) (res *response.Response, err error) {
	token, err := backend.ServiceRole().LIndex(ctx, req)
	if err != nil {
		return nil, err
	}
	response.JsonOkCtx(ctx, g.Map{"token": token})
	return
}
