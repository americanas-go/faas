package lambda

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/log"
	"github.com/aws/aws-lambda-go/lambdacontext"
	v2 "github.com/cloudevents/sdk-go/v2"
)

func fromKinesis(parentCtx context.Context, event Event) []*cloudevents.InOut {

	logger := log.FromContext(parentCtx)
	logger.Info("receiving Kinesis event")

	lc, _ := lambdacontext.FromContext(parentCtx)

	var inouts []*cloudevents.InOut

	for _, record := range event.Records {

		j, _ := json.Marshal(record)
		logger.Debug(string(j))

		var err error
		in := v2.NewEvent()

		if err = json.Unmarshal(record.Kinesis.Data, &in); err != nil {
			var data interface{}

			if err = json.Unmarshal(record.Kinesis.Data, &data); err != nil {
				err = errors.NewNotValid(err, "could not decode kinesis record")
			} else {

				if err = in.SetData(v2.ApplicationJSON, data); err != nil {
					err = errors.NewNotValid(err, "could not set data in event")
				}
			}
		}

		in.SetType(record.EventName)

		if in.ID() == "" {
			in.SetID(record.EventID)
		}

		in.SetSource(record.EventSource)

		in.SetExtension("awsRequestID", lc.AwsRequestID)
		in.SetExtension("invokedFunctionArn", lc.InvokedFunctionArn)

		inouts = append(inouts, &cloudevents.InOut{
			In:  &in,
			Err: err,
		})
	}

	return inouts
}
