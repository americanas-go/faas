package sns

import (
	"sync"

	"github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			context.Module(),
			fx.Provide(
				NewEvent,
			),
		)
	})

	return options
}
