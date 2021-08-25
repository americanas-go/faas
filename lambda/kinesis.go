package lambda

import (
	"encoding/json"

	"github.com/americanas-go/errors"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
)

func fromKinesis(record Record) (*event.Event, error) {
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

	return &in, err
}
