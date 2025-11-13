package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"go-service/internal/model/response"
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
	// 从 context 中获取 admin 信息
	//	adminValue := g.RequestFromCtx(ctx).GetCtxVar("admin")
	//	if adminValue.IsNil() {
	//		return "", gerror.New("未找到用户信息")
	//	}
	//	adminInfo := adminValue.Map()

	//前置中间件
	r.Middleware.Next()

}

func (s *sMiddleware) LErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	ctx := r.GetCtx()
	if err := r.GetError(); err != nil {

		exception := gerror.Current(err)
		g.Log("panic").Critical(ctx, err)

		message := "系统繁忙，请稍后再试"
		//错误提示拦截
		if exception != nil {
			if msg := exception.Error(); msg != "" {
				message = msg
			}
		}
		response.JsonErrCtx(ctx, message)
		//r.Response.WriteJson(v1.DefaultHandlerResponse{
		//	Code: http.StatusInternalServerError, //请求成功  情况：1 异常 2 系统报错 都为code=500
		//	Msg:  message,
		//	Data: "",
		//})
		return
	}
}
