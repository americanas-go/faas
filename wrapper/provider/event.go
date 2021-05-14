package provider

import (
	"context"
	"reflect"

	"github.com/americanas-go/serverless/repository"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type EventWrapperProvider struct {
	events repository.Event
	pkg    string
	impl   string
}

func NewEventWrapperProvider(events repository.Event) *EventWrapperProvider {
	impl := reflect.TypeOf(events).Kind().String()
	pkg := reflect.TypeOf(events).PkgPath()

	return &EventWrapperProvider{events: events, impl: impl, pkg: pkg}
}

func (e *EventWrapperProvider) Publish(ctx context.Context, events []*v2.Event) error {
	txn := newrelic.FromContext(ctx)

	seg := &newrelic.MessageProducerSegment{
		StartTime:            txn.StartSegmentNow(),
		Library:              e.impl,
		DestinationType:      "",
		DestinationName:      "",
		DestinationTemporary: false,
	}
	defer seg.End()

	return e.events.Publish(ctx, events)
}
