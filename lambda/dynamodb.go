package lambda

import (
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
)

func fromDynamoDB(record Record) (*event.Event, error) {

	in := v2.NewEvent()
	err := in.SetData("", record.DynamoDB)
	return &in, err
}
