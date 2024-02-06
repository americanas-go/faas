package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	"github.com/americanas-go/faas/cloudevents"
	logger "github.com/americanas-go/faas/cloudevents/plugins/contrib/americanas-go/log.v1"
	"github.com/americanas-go/faas/cmd"
	"github.com/americanas-go/faas/wrapper/provider"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	igce "github.com/americanas-go/ignite/cloudevents/sdk-go.v2"
	v2 "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/fx"
)

func main() {

	config.Load()
	ilog.New()

	options := fx.Options(
		provider.Module(),
		fx.Provide(
			func() igce.Handler {
				return Handle
			},
			func(events *provider.EventWrapperProvider) []cloudevents.Middleware {
				return []cloudevents.Middleware{
					logger.NewLogger(),
				}
			},
		),
	)

	// sets env var
	os.Setenv("FAAS_CMD_DEFAULT", "azure")

	// go run main.go help
	err := cmd.Run(options,
		cmd.NewAzure(),
	)

	if err != nil {
		panic(err)
	}
}

func Handle(ctx context.Context, in v2.Event) (*v2.Event, error) {

	e := v2.NewEvent()
	e.SetID("changeme")
	e.SetSubject("changeme")
	e.SetSource("changeme")
	e.SetType("changeme")
	e.SetExtension("partitionkey", "changeme")
	err := e.SetData(v2.TextPlain, "changeme")

	return &e, err
}
