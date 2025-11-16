package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/dao/jh_site/dao"
	"go-service/internal/service/backend"
	"go-service/utility/helpers"
)

type sPublic struct {
}

func init() {
	backend.RegisterPublic(&sPublic{})
}

//LLogin
/**
 * @desc：登录操作
 * @param ctx
 * @param req
 * @return string
 * @return error
 * @author : Carson
 */
func (s *sPublic) LLogin(ctx context.Context, req *backendRoute.PublicReq) (string, error) {
	//传参
	username := req.Username
	password := req.Password
	//验证
	admin, err := dao.GetAdminByPassword(username, password)
	if err != nil {
		return "", err
	}

	token, err := helpers.SetToken(ctx, admin)
	if err != nil {
		return "", err
	}

	return token, nil

}
