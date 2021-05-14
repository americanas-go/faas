package main

import (
	"context"

	"github.com/americanas-go/config"
	gice "github.com/americanas-go/ignite/cloudevents/sdk-go.v2"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
	"github.com/americanas-go/serverless/cloudevents"
	"github.com/americanas-go/serverless/cloudevents/plugins/logger"
	"github.com/americanas-go/serverless/cmd"
	v2 "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/fx"
)

func main() {

	config.Load()

	logrus.NewLogger()

	options := fx.Options(
		fx.Provide(
			func() gice.Handler {
				return Handle
			},
			func() []cloudevents.Middleware {
				return []cloudevents.Middleware{
					logger.NewLogger(),
				}
			},
		),
	)

	err := cmd.Run(options,
		cmd.NewNats(),
		cmd.NewCloudEvents(),
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
