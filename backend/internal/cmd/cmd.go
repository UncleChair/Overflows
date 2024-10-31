package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"overflows/internal/controller/auth"
	"overflows/internal/controller/users"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					auth.NewV1(),
					users.NewV1(),
				)
			})
			s.SetFileServerEnabled(true)
			s.SetServerRoot("./resource/public/www")
			s.Run()
			return nil
		},
	}
)
