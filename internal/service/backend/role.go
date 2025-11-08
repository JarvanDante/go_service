package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	IRole interface {
		LIndex(ctx context.Context, req *backendRoute.RoleReq) (string, error)
	}
)

var localRole IRole

func ServiceRole() IRole {
	return localRole
}

func RegisterRole(p IRole) {
	localRole = p
}
