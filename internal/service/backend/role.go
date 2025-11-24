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
		LCreate(ctx context.Context, req *backendRoute.CreateReq) (err error)
		LUpdate(ctx context.Context, req *backendRoute.UpdateReq) (err error)
		LDelete(ctx context.Context, req *backendRoute.DeleteReq) (err error)
		LSavePermission(ctx context.Context, req *backendRoute.SavePermissionReq) (err error)
	}
)

var localRole IRole

func ServiceRole() IRole {
	return localRole
}

func RegisterRole(p IRole) {
	localRole = p
}
