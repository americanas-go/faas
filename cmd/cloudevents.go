package cmd

import (
	"github.com/americanas-go/faas/cloudevents"
	gsfx "github.com/americanas-go/faas/fx"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

// NewCloudEvents returns CmdFunc for cloudevents command.
func NewCloudEvents() CmdFunc {
	return func(options fx.Option) *co.Command {
		return &co.Command{
			Use:   "cloudevents",
			Short: "cloudevents",
			Long:  "",
			RunE: func(cmd *co.Command, args []string) error {
				return gsfx.Run(cloudevents.HelperModule(options))
			},
		}
	}
}
