package lambda

import (
	"encoding/json"

	"github.com/americanas-go/errors"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
)

func fromSQS(record Record) (*event.Event, error) {
	var err error

	in := v2.NewEvent()
	body := []byte(record.Body)
	if err = json.Unmarshal(body, &in); err != nil {
		var data interface{}

		if err = json.Unmarshal(body, &data); err != nil {
			err = errors.NewNotValid(err, "could not decode SQS record")
		} else {

			if err = in.SetData(v2.ApplicationJSON, data); err != nil {
				err = errors.NewNotValid(err, "could not set data in event")
			}
		}
	}
	if in.ID() == "" {
		in.SetID(record.MessageId)
	}
	return &in, err
}
