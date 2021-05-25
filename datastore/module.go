package datastore

import (
	"sync"

	"github.com/americanas-go/faas/datastore/nats"
	"go.uber.org/fx"
)

var eventOnce sync.Once

func EventModule() fx.Option {

	options := fx.Options()

	eventOnce.Do(func() {

		value := EventProviderValue()

		switch value {
		default:
			options = nats.Module()
		}

	})

	return options
}
