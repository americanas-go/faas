package pubsub

import (
	"cloud.google.com/go/pubsub"
	"github.com/americanas-go/faas/repository"
)

// NewEvent returns a initialized client
func NewEvent(c *pubsub.Client) repository.Event {
	return NewClient(c)
}
