package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/dao/jh_site/model/entity"
)

type (
	IRole interface {
		LIndex(ctx context.Context, req *backendRoute.RoleReq) (role []*entity.AdminRole, err error)
	}
)

var localRole IRole

func ServiceRole() IRole {
	return localRole
}

func RegisterRole(p IRole) {
	localRole = p
}
