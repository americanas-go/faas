package cmd

import (
	"github.com/americanas-go/faas/azure"
	gsfx "github.com/americanas-go/faas/fx"
	co "github.com/spf13/cobra"
	"go.uber.org/fx"
)

// NewAzure returns CmdFunc for azure functions command.
func NewAzure() CmdFunc {
	return func(options fx.Option) *co.Command {
		return &co.Command{
			Use:   "azure",
			Short: "azure",
			Long:  "",
			RunE: func(CmdFunc *co.Command, args []string) error {
				return gsfx.Run(azure.HelperModule(options))
			},
		}
	}
}
