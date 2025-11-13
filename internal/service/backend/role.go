package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/dao/jh_site/model/entity"
	"go-service/internal/model/response"
)

type (
	IRole interface {
		LIndex(ctx context.Context, req *backendRoute.RolesReq) (role []*entity.AdminRole, err error)
		LPermissions(ctx context.Context, req *backendRoute.PermissionsReq) (res *response.RolePermissionsRes, err error)
	}
)

var localRole IRole

func ServiceRole() IRole {
	return localRole
}

func RegisterRole(p IRole) {
	localRole = p
}
