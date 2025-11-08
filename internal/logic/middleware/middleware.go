package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"go-service/internal/service/middleware"
	"go-service/utility/helpers"
)

type sMiddleware struct {
}

func init() {
	middleware.RegisterMiddleware(&sMiddleware{})
}

func (s *sMiddleware) LAuthToken(r *ghttp.Request) {
	//获取token
	adminInfo, err := helpers.GetAdminInfoFromToken(r)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code": 401,
			"msg":  "无效token",
			"data": "",
		})
	}

	// 设置上下文，后面的接口就能直接取
	r.SetCtxVar("admin", adminInfo)
	//admin := r.GetCtxVar("admin").Map()

	//前置中间件
	r.Middleware.Next()

}
