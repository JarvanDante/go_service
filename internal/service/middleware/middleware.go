package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		LAuthToken(r *ghttp.Request)
		LErrorHandler(r *ghttp.Request)
	}
)

var localRole IMiddleware

func ServiceMiddleware() IMiddleware {
	return localRole
}

func RegisterMiddleware(p IMiddleware) {
	localRole = p
}
