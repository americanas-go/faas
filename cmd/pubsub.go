package cmd

import (
	gsfx "github.com/americanas-go/faas/fx"
	"github.com/americanas-go/faas/pubsub"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

// NewPubSub returns CmdFunc for pubsub command.
func NewPubSub() CmdFunc {
	return func(options fx.Option) *co.Command {
		return &co.Command{
			Use:   "pubsub",
			Short: "pubsub",
			Long:  "",
			RunE: func(CmdFunc *co.Command, args []string) error {
				return gsfx.Run(pubsub.HelperModule(options))
			},
		}
	}
}
