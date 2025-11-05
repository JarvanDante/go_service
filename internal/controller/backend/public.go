package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/backend"
	"go-service/internal/model/response"
)

type PublicController struct{}

func NewPublicController() *PublicController {
	return &PublicController{}
}

func (c *PublicController) Login(ctx context.Context, req *backend.PublicReq) (res *response.Response, err error) {
	response.JsonOkCtx(ctx, g.Map{"token": "fsdfsdf"})
	return
}
