package cmd

import (
	gsfx "github.com/americanas-go/faas/fx"
	"github.com/americanas-go/faas/nats"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

// NewNats returns CmdFunc for nats command.
func NewNats() CmdFunc {
	return func(options fx.Option) *co.Command {
		return &co.Command{
			Use:   "nats",
			Short: "nats",
			Long:  "",
			RunE: func(CmdFunc *co.Command, args []string) error {
				return gsfx.Run(nats.HelperModule(options))
			},
		}
	}
}
