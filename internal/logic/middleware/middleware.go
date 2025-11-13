package middleware

import (
	"bytes"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"go-service/internal/model/response"
	"go-service/internal/service/middleware"
	"go-service/utility/helpers"
	"io"
	"strings"
)

type sMiddleware struct {
}

func init() {
	middleware.RegisterMiddleware(&sMiddleware{})
}

//LAuthToken
/**
 * @desc：解析token
 * @param r
 * @author : Carson
 */
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

//LErrorHandler
/**
 * @desc：拦截错误
 * @param r
 * @author : Carson
 */
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

		return
	}
}

//LRequestLog
/**
 * @desc：记录请求参数的中间件
 * @return ghttp.HandlerFunc
 * @author : Carson
 */
func (s *sMiddleware) LRequestLog(r *ghttp.Request) {

	// ====== 获取 Body（支持 JSON） ======
	var bodyBytes []byte
	if r.Request.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Request.Body)
		r.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 重要：写回去
	}

	// ====== 基础信息 ======
	clientIp := helpers.RemoteIp(r.Request)
	host := r.URL.Host
	path := r.URL.Path
	method := r.Method
	rawQuery := r.URL.RawQuery

	// ====== token 脱敏 ======
	bearer := r.Header.Get("Authorization")
	token := strings.TrimPrefix(bearer, "Bearer ")
	if len(token) > 10 {
		token = token[:10] + "****" // 简单脱敏
	}

	// ====== 记录日志 ======
	ctx := r.GetCtx()

	data := g.Map{
		"host":     host,
		"path":     path,
		"method":   method,
		"clientIp": clientIp,
	}

	if method == "POST" {
		data["body"] = string(bodyBytes)
	} else {
		data["rawQuery"] = rawQuery
	}

	g.Log().Info(ctx, "[Request]", data)

	// ====== 执行后续 ======
	r.Middleware.Next()
}
