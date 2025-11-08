package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/service/backend"
)

type sRole struct {
}

func init() {
	backend.RegisterRole(&sRole{})
}

func (s *sRole) LIndex(ctx context.Context, req *backendRoute.RoleReq) (string, error) {
	////传参
	//username := req.Username
	//password := req.Password
	////验证
	//admin, err := dao.GetAdmin(username, password)
	//if err != nil {
	//	return "", err
	//}
	//
	//token, err := helpers.SetToken(ctx, admin)
	//if err != nil {
	//	return "", err
	//}

	return "", nil

}
