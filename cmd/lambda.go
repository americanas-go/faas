package cmd

import (
	gsfx "github.com/americanas-go/serverless/fx"
	"github.com/americanas-go/serverless/lambda"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewLambda() CmdFunc {

	return func(options fx.Option) *co.Command {
		return &co.Command{
			Use:   "lambda",
			Short: "lambda",
			Long:  "",
			Run: func(CmdFunc *co.Command, args []string) {
				gsfx.Run(lambda.HelperModule(options))
			},
		}
	}
}
