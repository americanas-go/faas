package datadog

import (
	"context"

	"github.com/americanas-go/faas/cloudevents"
	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	v2 "github.com/cloudevents/sdk-go/v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

// DataDog represents a datadog agent middleware for events.
type DataDog struct {
	cloudevents.UnimplementedMiddleware
}

// NewDataDog creates a datadog agent middleware.
func NewDataDog() cloudevents.Middleware {
	if !IsEnabled() {
		return nil
	}
	return &DataDog{}
}

// BeforeAll starts a datadog transaction before processing all input event handlers.
// The transaction started is passed via context.
func (m *DataDog) BeforeAll(ctx context.Context, inout []*cloudevents.InOut) (context.Context, error) {
	if !datadog.IsTracerEnabled() {
		datadog.StartTracer(ctx, tracer.WithLambdaMode(true))
	}
	if !datadog.IsProfilerEnabled() {
		datadog.StartProfiler(ctx)
	}
	return ctx, nil
}

// Before enables the datadog transacation for use in multiple goroutines to be used by the handler.
func (m *DataDog) Before(ctx context.Context, in *v2.Event) (context.Context, error) {

	if !datadog.IsTracerEnabled() {
		opts := []ddtrace.StartSpanOption{
			tracer.ResourceName("execution"),
			tracer.SpanType("serverless"),
			tracer.Tag("id", in.ID()),
			tracer.Tag("source", in.Source()),
			tracer.Tag("subject", in.Subject()),
		}

		_, ctx = tracer.StartSpanFromContext(ctx, "execution", opts...)
	}

	return ctx, nil
}

// After checks if the handler has returned any error and notifies via datadog agent.
func (m *DataDog) After(ctx context.Context, in v2.Event, out *v2.Event, err error) (context.Context, error) {
	if !datadog.IsTracerEnabled() {
		span, ok := tracer.SpanFromContext(ctx)
		if ok {
			span.Finish()
		}
	}

	return ctx, nil
}

// Close finishes the datadog transaction.
func (m *DataDog) Close(ctx context.Context) error {
	if !datadog.IsTracerEnabled() {
		tracer.Stop()
	}
	if !datadog.IsProfilerEnabled() {
		profiler.Stop()
	}
	return nil
}
