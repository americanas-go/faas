package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/log"
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
	subscription := h.options.Subscription
	pubsubClient, err := pubsub.NewClient(ctx, h.options.ProjectId)
	if err != nil {
		logger.Errorf("pubsub.NewClient: %w", err)
		//TODO handle error
	}

	go h.run(ctx, pubsubClient, subscription)

	c := make(chan struct{})
	<-c
}

func (h *Helper) run(ctx context.Context, pubsubClient *pubsub.Client, subscriptionName string) {
	logger := log.FromContext(ctx)
	sub := pubsubClient.Subscription(subscriptionName)
	err := sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		m.Ack()
	})
	if err != nil {
		logger.Errorf("pubsub read error: %w", err)
	}
}
