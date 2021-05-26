package publisher

import (
	"sync"

	"github.com/americanas-go/faas/wrapper/provider"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			provider.Module(),
			fx.Provide(
				NewEventPublisher,
			),
		)
	})

	return options
}
