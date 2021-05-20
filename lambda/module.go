package lambda

import (
	"sync"

	"github.com/americanas-go/faas/cloudevents"
	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

func HelperModule(extraOptions fx.Option) fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			extraOptions,
			cloudevents.HandlerWrapperModule(),
			fx.Provide(
				NewDefaultHelper,
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
