package cloudevents

import (
	"context"

	cloudevents "github.com/americanas-go/ignite/cloudevents/sdk-go.v2"
)

type Helper struct {
	ctx    context.Context
	client *cloudevents.Client
}

func NewHelper(ctx context.Context, handler *HandlerWrapper) *Helper {

	client := cloudevents.NewDefaultClient(ctx, NewHandler(handler).Handle)

	return &Helper{
		client: client,
		ctx:    ctx,
	}
}

func (h *Helper) Start() {
	h.client.Start(h.ctx)
}
