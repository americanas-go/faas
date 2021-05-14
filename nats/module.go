package nats

import (
	"sync"

	"github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	ginatsfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/nats-io/nats.go.v1"
	"github.com/americanas-go/serverless/cloudevents"
	"go.uber.org/fx"
)

var once sync.Once

func HelperModule(extraOptions fx.Option) fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			context.Module(),
			extraOptions,
			ginatsfx.SubscriberModule(),
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
