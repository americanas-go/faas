package pubsub

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/americanas-go/errors"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/log"
	"github.com/cloudevents/sdk-go/v2/event"
)

// Helper assists in creating event handlers.
type Helper struct {
	handler *cloudevents.HandlerWrapper
	options *Options
}

// NewHelper returns a new Helper with options.
func NewHelper(ctx context.Context, options *Options,
	handler *cloudevents.HandlerWrapper) *Helper {

	return &Helper{
		handler: handler,
		options: options,
	}
}

// NewDefaultHelper returns a new Helper with default options.
func NewDefaultHelper(ctx context.Context, handler *cloudevents.HandlerWrapper) *Helper {

	opt, err := DefaultOptions()
	if err != nil {
		log.Fatal(err.Error())
	}

	return NewHelper(ctx, opt, handler)
}

func (h *Helper) Start() {
	ctx := context.Background()
	logger := log.FromContext(ctx)
	pubsubClient, err := pubsub.NewClient(ctx, h.options.ProjectId)
	if err != nil {
		logger.Errorf("pubsub.NewClient: %s", err.Error())
		//TODO handle error
	}

	go h.run(ctx, pubsubClient)

	c := make(chan struct{})
	<-c
}

func (h *Helper) run(ctx context.Context, pubsubClient *pubsub.Client) {
	logger := log.FromContext(ctx)
	sub := pubsubClient.Subscription(h.options.Subscription)
	sub.ReceiveSettings.Synchronous = false
	sub.ReceiveSettings.NumGoroutines = h.options.NumGoroutines
	sub.ReceiveSettings.MaxOutstandingMessages = h.options.MaxOutstandingMessages
	err := sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {

		go func(ctx context.Context, m pubsub.Message) {
			h.handle(ctx, m)
		}(ctx, *m)

		m.Ack()
	})
	if err != nil {
		logger.Errorf("pubsub read error: %s", err.Error())
	}

}

func (h *Helper) handle(ctx context.Context, msg pubsub.Message) {

	logger := log.FromContext(ctx).WithTypeOf(*h)

	in := event.New()
	if err := json.Unmarshal(msg.Data, &in); err != nil {
		var data interface{}
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			logger.Errorf("could not decode pubsub record. %s", err.Error())
			return
		}

		err := in.SetData("", data)
		if err != nil {
			logger.Errorf("could set data from pubsub record. %s", err.Error())
			return
		}
	}

	var inouts []*cloudevents.InOut

	inouts = append(inouts, &cloudevents.InOut{In: &in})

	if err := h.handler.Process(ctx, inouts); err != nil {
		logger.Error(errors.ErrorStack(err))
	}

}
