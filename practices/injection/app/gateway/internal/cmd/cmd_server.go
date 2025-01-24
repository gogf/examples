package cmd

import (
	"context"

	"main/app/gateway/internal/controller/user"
	"main/utility/injection"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"main/app/gateway/internal/controller/hello"
)

type ServerInput struct {
	g.Meta `name:"server" brief:"start service server"`
}
type ServerOutput struct{}

func (m *Main) Server(ctx context.Context, in ServerInput) (out *ServerOutput, err error) {
	injection.SetupDefaultInjector(ctx)
	defer injection.ShutdownDefaultInjector()

	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			hello.NewV1(),
			user.NewV1(),
		)
	})
	s.Run()
	return
}
