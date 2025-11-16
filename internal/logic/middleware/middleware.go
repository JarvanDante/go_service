package middleware

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"go-service/internal/dao/jh_site/dao"
	"go-service/internal/model/response"
	"go-service/internal/service/middleware"
	"go-service/utility/helpers"
	"io"
	"time"
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

	adminEntity, err := dao.GetAdminById(gconv.Uint(adminInfo["id"]))
	if err != nil {
		return
	}
	// 设置上下文，后面的接口就能直接取
	r.SetCtxVar("admin", adminEntity)
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
func genTraceId() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func (s *sMiddleware) LRequestLog(r *ghttp.Request) {
	start := time.Now()

	// ===== 读取 Request Body =====
	var reqBody []byte
	if r.Request.Body != nil {
		reqBody, _ = io.ReadAll(r.Request.Body)
		r.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	}

	// 使用多个参数，避免字符串拼接
	g.Log().Info(r.GetCtx(),
		"[Request]",
		"method:", r.Method,
		"path:", r.URL.Path,
		"query:", r.GetQueryMap(),
		"ip:", r.GetClientIp(),
	)

	// ===== 执行后续逻辑 =====
	r.Middleware.Next()

	// ===== 捕获 Response =====
	resBody := r.Response.BufferString()

	g.Log().Info(r.GetCtx(),
		"[Response]",
		"status:", r.Response.Status,
		"cost_ms:", time.Since(start).Milliseconds(),
		"response:", string(resBody),
	)

	return
}
