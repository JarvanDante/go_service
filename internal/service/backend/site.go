package backend

import (
	"context"
)

type (
	ISite interface {
		LBasicSetting(ctx context.Context) (map[string]interface{}, error)
	}
)

var localSite ISite

func ServiceSite() ISite {
	return localSite
}

func RegisterSiten(p ISite) {
	localSite = p
}
