package log

import (
	"github.com/americanas-go/faas/cloudevents"
	"sync"

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
			fx.Provide(
				fx.Annotated{
					Group: "_faas_middleware_",
					Target: func() cloudevents.Middleware {
						return NewLogger()
					},
				},
			),
		)
	})

	return options
}