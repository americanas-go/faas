package mock

import (
	"context"

	v2 "github.com/cloudevents/sdk-go/v2"
)

// mock holds evertything needed to publish a product
type mock struct {
}

func NewMock() *mock {
	return &mock{}
}

// Publish publishes an array of products on
func (p *mock) Publish(ctx context.Context, events []*v2.Event) error {
	return nil
}
