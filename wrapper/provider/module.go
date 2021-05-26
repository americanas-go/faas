package provider

import (
	"sync"

	"github.com/americanas-go/faas/datastore"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			datastore.EventModule(),
			fx.Provide(
				NewEventWrapperProvider,
			),
		)
	})

	return options
}
