package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	IUserLevel interface {
		LIndex(ctx context.Context, req *backendRoute.UserLevelsReq) (res interface{}, err error)
		LCreateUserLevel(ctx context.Context, req *backendRoute.CreateUserLevelReq) (err error)
		LGetUpdateUserLevel(ctx context.Context, req *backendRoute.GetUpdateUserLevelReq) (res interface{}, err error)
		LUpdateUserLevel(ctx context.Context, req *backendRoute.UpdateUserLevelReq) (err error)
		LDeleteUserLevel(ctx context.Context, req *backendRoute.DeleteUserLevelReq) (err error)
	}
)

var localUserLevel IUserLevel

func ServiceUserLevel() IUserLevel {
	return localUserLevel
}

func RegisterUserLevel(p IUserLevel) {
	localUserLevel = p
}
