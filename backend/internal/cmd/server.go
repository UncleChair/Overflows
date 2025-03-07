package cmd

import (
	"context"
	"overflows/api"
	"overflows/internal/controller/auth"
	"overflows/internal/controller/users"
	"overflows/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Server = gcmd.Command{
		Name:  "server",
		Brief: "start http server backend only",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			openapi := s.GetOpenApi()
			openapi.Config.CommonResponse = api.CommonRes{}
			openapi.Config.CommonResponseDataField = `Data`
			openapi.Info = goai.Info{
				Title:       "Overflows API Reference",
				Description: "Overflows",
			}
			s.Use(
				service.Middleware().Ctx,
				service.Middleware().CORS,
				service.Middleware().ResponseHandler,
			)
			s.EnableAdmin("/system")
			s.BindMiddleware("/system/*", service.Middleware().Ctx)
			s.Group("/api/v1/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.Bind(auth.NewV1())
				})
				group.Group("/users", func(group *ghttp.RouterGroup) {
					group.Bind(users.NewV1())
				})
			})
			s.SetFileServerEnabled(true)
			s.SetServerRoot("resource/public/www")
			s.BindHandler("/*", func(r *ghttp.Request) {
				path := r.URL.Path
				if len(path) >= 4 && path[0:4] != "/api" {
					r.Response.ServeFile("resource/public/www/index.html")
				}
			})
			s.Run()
			return nil
		},
	}
)
