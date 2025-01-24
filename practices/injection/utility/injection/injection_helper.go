package injection

import (
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/samber/do"
)

// SetupShutdownHelper sets up a shutdown helper.
func SetupShutdownHelper[T any](injector *do.Injector, service T, onShutdown func(service T) error) {
	do.Provide(injector, func(i *do.Injector) (ShutdownHelper[T], error) {
		g.Log().Debugf(gctx.GetInitCtx(), "NewShutdownHelper: %s", reflect.TypeOf(service))
		return NewShutdownHelper(service, onShutdown), nil
	})
	do.MustInvoke[ShutdownHelper[T]](injector)
}

// SetupShutdownHelperNamed sets up a shutdown helper with a name.
func SetupShutdownHelperNamed[T any](injector *do.Injector, service T, name string, onShutdown func(service T) error) {
	name = fmt.Sprintf("ShutdownHelper:%s", name)
	do.ProvideNamed(injector, name, func(i *do.Injector) (ShutdownHelper[T], error) {
		g.Log().Debugf(
			gctx.GetInitCtx(),
			"NewShutdownHelper: %s, %s",
			reflect.TypeOf(service), name,
		)
		return NewShutdownHelperNamed(service, name, onShutdown), nil
	})
	do.MustInvokeNamed[ShutdownHelper[T]](injector, name)
}

// ShutdownHelper is a helper struct for shutdown.
type ShutdownHelper[T any] struct {
	name       string
	service    T
	onShutdown func(service T) error
}

// NewShutdownHelper creates a new ShutdownHelper.
func NewShutdownHelper[T any](service T, onShutdown func(service T) error) ShutdownHelper[T] {
	return ShutdownHelper[T]{
		service:    service,
		onShutdown: onShutdown,
	}
}

// NewShutdownHelperNamed creates a new ShutdownHelper with a name.
func NewShutdownHelperNamed[T any](service T, name string, onShutdown func(service T) error) ShutdownHelper[T] {
	return ShutdownHelper[T]{
		name:       name,
		service:    service,
		onShutdown: onShutdown,
	}
}

// Shutdown shuts down the service.
func (h ShutdownHelper[T]) Shutdown() error {
	g.Log().Debugf(
		gctx.GetInitCtx(),
		"ShutdownHelper Shutdown: %s, %s",
		reflect.TypeOf(h.service), h.name,
	)
	return h.onShutdown(h.service)
}
