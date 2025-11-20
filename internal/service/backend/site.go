package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	ISite interface {
		LBasicSetting(ctx context.Context) (map[string]interface{}, error)
		LUpdateBasicSetting(ctx context.Context, req *backendRoute.UpdateBasicSettingReq) error
		LRegisterSetting(ctx context.Context) (interface{}, error)
		LUpdateRegisterSetting(ctx context.Context, req *backendRoute.UpdateRegisterSettingReq) error
	}
)

var localSite ISite

func ServiceSite() ISite {
	return localSite
}

func RegisterSite(p ISite) {
	localSite = p
}
