package nats

import (
	"context"

	ginats "github.com/americanas-go/ignite/nats-io/nats.go.v1"
	"github.com/americanas-go/serverless/repository"
)

// NewEvent returns a initialized client
func NewEvent(ctx context.Context) repository.Event {
	publisher, _ := ginats.NewDefaultPublisher(ctx)
	return NewClient(publisher)
}
