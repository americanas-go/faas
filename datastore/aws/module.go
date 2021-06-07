package aws

import (
	"sync"

	igaws "github.com/americanas-go/ignite/aws/aws-sdk-go.v2"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {

	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				igaws.NewConfig,
			),
		)
	})

	return options
}
