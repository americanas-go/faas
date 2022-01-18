package lambda

import (
	"encoding/json"

	"github.com/americanas-go/errors"
	"github.com/aws/aws-lambda-go/events"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
)

func fromSQS(record Record) (*event.Event, error) {
	var err error

	in := v2.NewEvent()
	body := []byte(record.Body)
	if err = json.Unmarshal(body, &in); err != nil {
		var data interface{}
		snsEvent := events.SNSEntity{}

		if err = json.Unmarshal(body, &snsEvent); err == nil {
			if snsEvent.Message != "" {
				if er := json.Unmarshal([]byte(snsEvent.Message), &data); er != nil {
					err = errors.NewNotValid(er, "could not decode SNS message from SQS record")
				} else if er := in.SetData(v2.ApplicationJSON, data); er != nil {
					return nil, errors.NewNotValid(er, "could not set data in event")
				}
			}
		}
	}
	if in.ID() == "" {
		in.SetID(record.MessageId)
	}
	return &in, err
}
