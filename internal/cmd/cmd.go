package cmd

import (
	"go-service/internal/controller/backend"
	"go-service/internal/controller/frontend"
	"go-service/internal/service/middleware"
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
				group.Middleware(
					ghttp.MiddlewareCORS,
					ghttp.MiddlewareHandlerResponse,
				)
				group.Group("/frontend", func(groupFront *ghttp.RouterGroup) {
					groupFront.Bind(
						frontend.NewDemo(),
					)
				})

				group.Group("/backend", func(groupBackend *ghttp.RouterGroup) {
					groupBackend.Bind(
						backend.NewPublicController(),
					)
				})
				group.Group("/backend", func(groupBackend *ghttp.RouterGroup) {
					groupBackend.Middleware(
						ghttp.MiddlewareCORS, // ← 关键
						middleware.ServiceMiddleware().LRequestLog,
						middleware.ServiceMiddleware().LAuthToken,
						middleware.ServiceMiddleware().LErrorHandler,
					)
					groupBackend.Bind(
						backend.NewRoleController(),
						backend.NewAdminController(),
						backend.NewSiteController(),
						backend.NewGameController(),
						backend.NewLogController(),
						backend.NewUserController(),
						backend.NewUserGradeController(),
						backend.NewUserLevelController(),
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
