package cloudevents

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/log"
	v2 "github.com/cloudevents/sdk-go/v2"
)

type Handler struct {
	handler *HandlerWrapper
}

func NewHandler(h *HandlerWrapper) *Handler {
	return &Handler{handler: h}
}

func (h *Handler) Handle(ctx context.Context, in v2.Event) (out *v2.Event, err error) {

	logger := log.FromContext(ctx).WithTypeOf(*h)

	if isSQSEvent(ctx, &in) {
		fromSQS(ctx, &in)
	}

	inouts := []*InOut{
		{
			In:  &in,
			Err: err,
		},
	}

	err = h.handler.Process(ctx, inouts)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
		return nil, err
	}

	logger.Debug("all events called")

	for _, inout := range inouts {
		if inout.Err != nil {
			err := errors.Wrap(inout.Err, errors.New("closing with errors lambda handle"))
			logger.Error(errors.ErrorStack(err))
			return nil, err
		}
	}

	return inouts[0].Out, nil
}
