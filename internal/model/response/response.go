package response

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"    description:"业务状态码" default:"200"`
	Status  int         `json:"status" description:"状态"`
	Message string      `json:"message" description:"返回说明"`
	Data    interface{} `json:"data" description:"返回数据"`
}

func JsonOkCtx(ctx context.Context, data interface{}, msg ...string) {
	JsonOk(g.RequestFromCtx(ctx), data, msg...)
}

func JsonOk(r *ghttp.Request, data interface{}, msg ...string) {
	m := "success"
	if len(msg) > 0 {
		m = msg[0]
	}
	r.Response.WriteJson(Response{
		Code:    http.StatusOK,
		Status:  1,
		Message: m,
		Data:    data,
	})
}

func JsonErrCtx(ctx context.Context, msg string) {
	JsonErr(g.RequestFromCtx(ctx), msg, 500)
}

func JsonErr(r *ghttp.Request, msg string, code ...int) {
	m := "default"
	if len(msg) > 0 {
		m = msg
	}
	c := 500
	if len(code) > 0 {
		c = code[0]
	}
	r.Response.ClearBuffer()
	r.Response.WriteJson(Response{
		Code:    c,
		Status:  0,
		Message: m,
		Data:    nil,
	})
}
