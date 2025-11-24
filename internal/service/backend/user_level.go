package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	IUserLevel interface {
		LIndex(ctx context.Context, req *backendRoute.UserLevelsReq) (res interface{}, err error)
	}
)

var localUserLevel IUserLevel

func ServiceUserLevel() IUserLevel {
	return localUserLevel
}

func RegisterUserLevel(p IUserLevel) {
	localUserLevel = p
}
