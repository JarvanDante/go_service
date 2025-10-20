package service

import "context"

type IUser interface {
	Login(ctx context.Context, username, password string) (string, error)
}

var localUser IUser

func User() IUser {
	return localUser
}

func RegisterUser(u IUser) {
	localUser = u
}
