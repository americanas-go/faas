package lambda

import (
	"encoding/json"
	"strings"

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
		unmarshal(&in, body)
	} else if in.Data() != nil {
		unmarshal(&in, in.Data())
	}
	if in.ID() == "" {
		in.SetID(record.MessageId)
	}
	return &in, err
}
func unmarshal(in *event.Event, body []byte) error {
	var err error
	var data interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return errors.New("not json message on SQS record")
	} else if isSNSMessage(body) {
		if err = unmarshalSNSMessage(body, &data); err != nil {
			return err
		}
	}
	if err = in.SetData(v2.ApplicationJSON, data); err != nil {
		return errors.NewNotValid(err, "could not set data in event")
	}
	return nil
}
func isSNSMessage(data []byte) bool {
	s := string(data)
	return strings.Contains(s, "\"TopicArn\":") && strings.Contains(s, "\"Message\":")
}

func unmarshalSNSMessage(body []byte, data *interface{}) (err error) {
	snsEvent := events.SNSEntity{}
	if err = json.Unmarshal(body, &snsEvent); err != nil {
		return errors.NewNotValid(err, "unrecognized message on SQS record")
	}
	if snsEvent.Message != "" {
		if er := json.Unmarshal([]byte(snsEvent.Message), data); er != nil {
			err = errors.NewNotValid(er, "SNS message is not a json")
		}
	}
	return
}
