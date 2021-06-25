package publisher

import (
	"sync"

	"github.com/americanas-go/faas/wrapper/provider"
	"go.uber.org/fx"
)

var once sync.Once

// Module returns fx module for initialization of event publisher middleware.
// Which depends on event wrapper provider module.
//
// The module is only loaded once.
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
