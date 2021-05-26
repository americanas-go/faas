package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/faas/cloudevents/plugins/contrib/americanas-go/log.v1"
	"github.com/americanas-go/faas/cmd"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	gice "github.com/americanas-go/ignite/cloudevents/sdk-go.v2"
	v2 "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/fx"
)

func main() {

	config.Load()

	ilog.New()

	options := fx.Options(
		fx.Provide(
			func() gice.Handler {
				return Handle
			},
			func() []cloudevents.Middleware {
				return []cloudevents.Middleware{
					log.NewLogger(),
				}
			},
		),
	)

	// sets env var
	os.Setenv("FAAS_CMD_DEFAULT", "nats")

	// go run main.go help
	err := cmd.Run(options,

		// go run main.go nats
		// or
		// FAAS_CMD_DEFAULT=nats go run main.go
		cmd.NewNats(),

		// go run main.go cloudevents
		// or
		// FAAS_CMD_DEFAULT=cloudevents go run main.go
		cmd.NewCloudEvents(),

		// go run main.go lambda
		// or
		// FAAS_CMD_DEFAULT=lambda go run main.go
		cmd.NewLambda(),
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
