package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/dao/jh_site/model/entity"
	entityJh "go-service/internal/dao/jinhuang/model/entity"
)

type (
	IRole interface {
		LIndex(ctx context.Context, req *backendRoute.RolesReq) (role []*entity.AdminRole, err error)
		LPermissions(ctx context.Context, req *backendRoute.PermissionsReq) (role []*entityJh.AdminPermission, err error)
	}
)

var localRole IRole

func ServiceRole() IRole {
	return localRole
}

func RegisterRole(p IRole) {
	localRole = p
}
