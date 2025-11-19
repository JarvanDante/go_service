package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	ISite interface {
		LBasicSetting(ctx context.Context) (map[string]interface{}, error)
		LUpdateBasicSetting(ctx context.Context, req *backendRoute.UpdateBasicSettingReq) error
	}
)

var localSite ISite

func ServiceSite() ISite {
	return localSite
}

func RegisterSiten(p ISite) {
	localSite = p
}
