package lambda

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/log"
	"github.com/aws/aws-lambda-go/lambdacontext"
	v2 "github.com/cloudevents/sdk-go/v2"
	"golang.org/x/sync/errgroup"
)

func fromS3(parentCtx context.Context, event Event) []*cloudevents.InOut {

	logger := log.FromContext(parentCtx)
	logger.Info("receiving S3 event")

	lc, _ := lambdacontext.FromContext(parentCtx)

	var inouts []*cloudevents.InOut

	g, gctx := errgroup.WithContext(parentCtx)

	for _, record := range event.Records {

		record := record

		g.Go(func() error {

			j, _ := json.Marshal(record)
			logger.Debug(string(j))

			var err error

			in := v2.NewEvent()
			in.SetID(record.S3.S3.Object.Key)
			in.SetType(record.S3.EventName)
			in.SetData("", record.S3.S3)
			in.SetSource(record.EventSource)

			in.SetExtension("awsRequestID", lc.AwsRequestID)
			in.SetExtension("invokedFunctionArn", lc.InvokedFunctionArn)

			inouts = append(inouts, &cloudevents.InOut{
				In:  &in,
				Err: err,
			})

			return nil
		})

	}

	if err := g.Wait(); err == nil {
		logger.Debug("all events converted")
	}

	gctx.Done()

	return inouts
}
