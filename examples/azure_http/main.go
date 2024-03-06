package main

import (
	"context"
	"github.com/americanas-go/config"
	logger "github.com/americanas-go/faas/cloudevents/plugins/contrib/americanas-go/log.v1"
	"github.com/americanas-go/faas/cmd"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	igce "github.com/americanas-go/ignite/cloudevents/sdk-go.v2"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"os"
)

func init() {
	// sets env var
	os.Setenv("FAAS_CMD_DEFAULT", "azure")
	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "TRACE")
}

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

	randData := Message{
		Random: uuid.NewString(),
	}

	e.SetExtension("partitionkey", "changeme")

	err := e.SetData("", randData)

	return &e, err
}

type Message struct {
	Random string
}
