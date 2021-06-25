package sns

import (
	"sync"

	"github.com/americanas-go/faas/datastore/aws"
	igsns "github.com/americanas-go/ignite/aws/aws-sdk-go.v2/client/sns"
	"github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"go.uber.org/fx"
)

var once sync.Once

// Module loads the sns module providing an initialized client.
//
// The module is only loaded once.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			context.Module(),
			aws.Module(),
			fx.Provide(
				sns.NewFromConfig,
				igsns.NewClient,
				NewEvent,
			),
		)
	})

	return options
}
