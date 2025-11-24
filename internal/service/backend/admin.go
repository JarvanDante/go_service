package backend

import (
	"context"
	"go-service/api/backendRoute"
	"go-service/internal/model/response"
)

type (
	IAdmin interface {
		LGetInfo(ctx context.Context, req *backendRoute.GetInfoReq) (res *response.GetInfoRes, err error)
		LAdmins(ctx context.Context, req *backendRoute.AdminsReq) (res interface{}, err error)
		LCreateAdmin(ctx context.Context, req *backendRoute.CreateAdminReq) (err error)
		LUpdateAdmin(ctx context.Context, req *backendRoute.UpdateAdminReq) (err error)
		LDeleteAdmin(ctx context.Context, req *backendRoute.DeleteAdminReq) (err error)
	}
)

var localAdmin IAdmin

func ServiceAdmin() IAdmin {
	return localAdmin
}

func RegisterAdmin(p IAdmin) {
	localAdmin = p
}
