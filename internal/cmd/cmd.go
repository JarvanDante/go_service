package cmd

import (
	"context"
	"go-service/internal/controller"
	"go-service/internal/controller/frontend"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"go-service/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "by",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
					controller.NewUserV1(),
					frontend.NewDemo(),
				)
			})
			s.Run()
			return nil
		},
	}
)
