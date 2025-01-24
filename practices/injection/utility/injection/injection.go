package injection

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/samber/do"
)

var defaultInjector *do.Injector

// MustInvoke invokes the function with the default injector and panics if any error occurs.
func MustInvoke[T any]() T {
	return do.MustInvoke[T](defaultInjector)
}

// Invoke invokes the function with the default injector.
func Invoke[T any]() (T, error) {
	return do.Invoke[T](defaultInjector)
}

// SetupDefaultInjector initializes the default injector with the given context.
func SetupDefaultInjector(ctx context.Context) *do.Injector {
	if defaultInjector != nil {
		return defaultInjector
	}
	injector := do.NewWithOpts(&do.InjectorOpts{})

	injectMongo(ctx, injector)
	injectRedis(ctx, injector)
	injectGrpcClients(ctx, injector)

	defaultInjector = injector
	return defaultInjector
}

// ShutdownDefaultInjector shuts down the default injector.
func ShutdownDefaultInjector() {
	if defaultInjector != nil {
		if err := defaultInjector.Shutdown(); err != nil {
			g.Log().Debugf(gctx.GetInitCtx(), "ShutdownDefaultInjector: %+v", err)
		}
		defaultInjector = nil
	}
}
