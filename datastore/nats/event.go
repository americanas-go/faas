package nats

import (
	"context"

	"github.com/americanas-go/faas/repository"
	ginats "github.com/americanas-go/ignite/nats-io/nats.go.v1"
)

// NewEvent returns a initialized client
func NewEvent(ctx context.Context) repository.Event {
	publisher, _ := ginats.NewDefaultPublisher(ctx)
	return NewClient(publisher)
}
