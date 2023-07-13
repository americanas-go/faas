package nats

import (
	"context"

	"github.com/americanas-go/faas/repository"
	ginats "github.com/americanas-go/ignite/nats-io/nats.go.v1"
)

// NewEvent returns an initialized NATS client that implements event repository.
func NewEvent(ctx context.Context) repository.Event {
	conn, _ := ginats.NewConn(ctx)
	return NewClient(conn)
}
