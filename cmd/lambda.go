package cmd

import (
	gsfx "github.com/americanas-go/faas/fx"
	"github.com/americanas-go/faas/lambda"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

// NewLambda returns CmdFunc for lambda command.
func NewLambda() CmdFunc {

	return func(options fx.Option) *co.Command {
		return &co.Command{
			Use:   "lambda",
			Short: "lambda",
			Long:  "",
			RunE: func(CmdFunc *co.Command, args []string) error {
				return gsfx.Run(lambda.HelperModule(options))
			},
		}
	}
}
