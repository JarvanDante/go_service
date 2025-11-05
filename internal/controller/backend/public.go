package backend

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
	"go-service/internal/service/backend"
)

type PublicController struct{}

func NewPublicController() *PublicController {
	return &PublicController{}
}

func (c *PublicController) Login(ctx context.Context, req *backendRoute.PublicReq) (res *response.Response, err error) {
	token, err := backend.ServicePublic().LLogin(ctx, req)
	fmt.Println(token)
	if err != nil {
		return nil, err
	}
	response.JsonOkCtx(ctx, g.Map{"token": token})
	return
}
