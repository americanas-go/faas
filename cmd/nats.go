package cmd

import (
	gsfx "github.com/americanas-go/serverless/fx"
	"github.com/americanas-go/serverless/nats"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewNats() CmdFunc {
	return func(options fx.Option) *co.Command {
		return &co.Command{
			Use:   "nats",
			Short: "nats",
			Long:  "",
			Run: func(CmdFunc *co.Command, args []string) {
				gsfx.Run(nats.HelperModule(options))
			},
		}
	}
}
