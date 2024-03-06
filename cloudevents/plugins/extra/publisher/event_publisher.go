package publisher

import (
	"context"
	"time"

	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/faas/repository"
	"github.com/americanas-go/faas/wrapper/provider"
	"github.com/americanas-go/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

// EventPublisher represents an event publisher middleware.
type EventPublisher struct {
	cloudevents.UnimplementedMiddleware
	events  repository.Event
	options *Options
}

// NewEventPublisher creates an event publisher middleware.
func NewEventPublisher(options *Options, events *provider.EventWrapperProvider) cloudevents.Middleware {
	if !options.Enabled {
		return nil
	}
	return &EventPublisher{events: events, options: options}
}

// AfterAll publishes all output events after processing all handlers.
func (p *EventPublisher) AfterAll(ctx context.Context, inouts []*cloudevents.InOut) (context.Context, error) {

	logger := log.FromContext(ctx).WithTypeOf(*p)

	var inErr []*v2.Event

	for _, inout := range inouts {
		if inout.Err != nil {
			inout.In.SetExtension("error", inout.Err.Error())
			inout.In.SetSubject(p.options.Error.Topic)
			inErr = append(inErr, inout.In)
			logger.Warn("the messages could not be published because one or more messages contain errors")
			return ctx, nil
		}
	}

	if p.options.Error.Enabled {
		if er := p.events.Publish(ctx, inErr); er != nil {
			return ctx, er
		}
	}

	if !p.options.Success.Enabled {
		return ctx, nil
	}

	var outs []*v2.Event

	for _, inout := range inouts {

		out := inout.Out
		in := inout.In

		if out != nil {

			id := uuid.New()

			out.SetID(id.String())
			out.SetExtension("parentId", in.ID())
			out.SetTime(time.Now())

			for key, value := range in.Extensions() {
				out.SetExtension(key, value)
			}

			outs = append(outs, out)
		}

	}

	if er := p.events.Publish(ctx, outs); er != nil {
		return ctx, er
	}

	logger.Info("published events")

	return ctx, nil
}
