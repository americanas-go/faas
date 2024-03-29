package nats

import (
	"sync"

	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	ginatsfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/nats-io/nats.go.v1"
	"go.uber.org/fx"
)

var once sync.Once

// HelperModule returns fx module for initialization of helper to start NATS client for handlers.
//
// The module is only loaded once.
func HelperModule(extraOptions fx.Option) fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			context.Module(),
			extraOptions,
			ginatsfx.Module(),
			cloudevents.HandlerWrapperModule(),
			fx.Provide(
				DefaultOptions,
				NewHelper,
			),
			fx.Invoke(
				func(helper *Helper) {
					helper.Start()
				},
			),
		)
	})

	return options
}
