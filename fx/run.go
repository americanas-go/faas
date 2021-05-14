package fx

import (
	gifx "github.com/americanas-go/ignite/go.uber.org/fx.v1"
	"go.uber.org/fx"
)

func Run(options fx.Option) {
	gifx.NewApp(options).Run()
}
