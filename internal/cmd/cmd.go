package cmd

import (
	"go-service/internal/controller/frontend"
	"golang.org/x/net/context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
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
				group.Group("/frontend", func(groupFront *ghttp.RouterGroup) {
					groupFront.Bind(
						frontend.NewDemo(),
					)
				})

				group.Group("/backend", func(groupBackend *ghttp.RouterGroup) {
					groupBackend.Bind(
						frontend.NewDemo(),
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
