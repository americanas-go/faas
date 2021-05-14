package nats

import (
	"context"

	"github.com/americanas-go/ignite/nats-io/nats.go.v1"
	"github.com/americanas-go/log"
	"github.com/americanas-go/serverless/cloudevents"
)

type Helper struct {
	handler    *cloudevents.HandlerWrapper
	queue      string
	subjects   []string
	subscriber *nats.Subscriber
}

func NewHelper(ctx context.Context, subscriber *nats.Subscriber, options *Options,
	handler *cloudevents.HandlerWrapper) *Helper {

	return &Helper{
		handler:    handler,
		queue:      options.Queue,
		subjects:   options.Subjects,
		subscriber: subscriber,
	}
}

func NewDefaultHelper(ctx context.Context, subscriber *nats.Subscriber, handler *cloudevents.HandlerWrapper) *Helper {

	opt, err := DefaultOptions()
	if err != nil {
		log.Fatal(err.Error())
	}

	return NewHelper(ctx, subscriber, opt, handler)
}

func (h *Helper) Start() {

	for i := range h.subjects {
		go h.subscribe(context.Background(), h.subjects[i])
	}

	c := make(chan struct{})
	<-c
}

func (h *Helper) subscribe(ctx context.Context, subject string) {

	logger := log.FromContext(ctx).WithTypeOf(*h)

	subscriber := NewSubscriberListener(h.subscriber, h.handler, subject, h.queue)
	subscribe, err := subscriber.Subscribe(ctx)
	if err != nil {
		logger.Error(err)
	}

	if subscribe.IsValid() {
		logger.Infof("nats: subscribed on %s with queue %s", subject, h.queue)
	} else {
		logger.Errorf("nats: not subscribed on %s with queue %s", subject, h.queue)
	}

}
