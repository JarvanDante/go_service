package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	IUser interface {
		LIndex(ctx context.Context, req *backendRoute.UsersReq) (res interface{}, err error)
	}
)

var localUser IUser

func ServiceUser() IUser {
	return localUser
}

func RegisterUser(p IUser) {
	localUser = p
}
