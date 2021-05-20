package nats

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/ignite/nats-io/nats.go.v1"
	"github.com/americanas-go/log"
	"github.com/cloudevents/sdk-go/v2/event"
	n "github.com/nats-io/nats.go"
)

type SubscriberListener struct {
	q       *nats.Subscriber
	handler *cloudevents.HandlerWrapper
	subject string
	queue   string
}

func NewSubscriberListener(q *nats.Subscriber, handler *cloudevents.HandlerWrapper, subject string,
	queue string) *SubscriberListener {
	return &SubscriberListener{
		q:       q,
		handler: handler,
		subject: subject,
		queue:   queue,
	}
}

func (l *SubscriberListener) Subscribe(ctx context.Context) (*n.Subscription, error) {
	return l.q.Subscribe(l.subject, l.queue, l.h)
}

func (l *SubscriberListener) h(msg *n.Msg) {

	in := event.New()
	err := json.Unmarshal(msg.Data, &in)
	if err != nil {

		var data interface{}

		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Errorf("could not decode nats record. %s", err.Error())
		} else {
			err := in.SetData("", data)
			if err != nil {
				log.Errorf("could set data from nats record. %s", err.Error())
				return
			}
		}

	}

	logger := log.WithTypeOf(*l).
		WithField("subject", l.subject).
		WithField("queue", l.queue)

	ctx := logger.ToContext(context.Background())

	var inouts []*cloudevents.InOut

	inouts = append(inouts, &cloudevents.InOut{In: &in})

	err = l.handler.Process(ctx, inouts)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
	}

}
