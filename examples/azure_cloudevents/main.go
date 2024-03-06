package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	logger "github.com/americanas-go/faas/cloudevents/plugins/contrib/americanas-go/log.v1"
	"github.com/americanas-go/faas/cmd"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	igce "github.com/americanas-go/ignite/cloudevents/sdk-go.v2"
	"github.com/americanas-go/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/fx"
)

func main() {

	config.Load()
	ilog.New()

	options := fx.Options(
		logger.Module(),
		fx.Provide(
			func() igce.Handler {
				return Handle
			},
		),
	)

	// sets env var
	os.Setenv("FAAS_CMD_DEFAULT", "cloudevents")

	// go run main.go help
	err := cmd.Run(options,
		cmd.NewCloudEvents(),
	)

	if err != nil {
		panic(err)
	}
}

func Handle(ctx context.Context, in v2.Event) (*v2.Event, error) {

	log.Info(in.Data())

	return nil, nil
}
