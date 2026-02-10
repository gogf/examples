package cmd

import (
	"context"

	"practices/quick-demo/internal/controller/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"practices/quick-demo/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func:  mainFunc,
	}
)

func mainFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			hello.NewV1(),
			user.NewV1(),
		)
	})
	s.Run()
	return nil
}
