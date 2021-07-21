package lambda

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/log"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// Handler can be used to process events.
type Handler struct {
	handler *cloudevents.HandlerWrapper
	options *Options
}

// NewHandler creates a new handler wrapped in middleware.
func NewHandler(handler *cloudevents.HandlerWrapper, options *Options) *Handler {
	return &Handler{handler: handler, options: options}
}

// Handle processes an event by calling the necessary middlewares.
func (h *Handler) Handle(ctx context.Context, event Event) error {

	logger := log.FromContext(ctx).WithTypeOf(*h)

	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return errors.Internalf("lambda context not exists")
	}

	logger = logger.WithField("awsrequestid", lc.AwsRequestID)
	ctx = logger.ToContext(ctx)

	if h.options.Skip {
		logger.Info("skipping event")
		return nil
	}

	inouts, err := h.getInOuts(ctx, event)
	if err != nil {
		return err
	}

	if len(inouts) > 0 {

		err := h.handler.Process(ctx, inouts)
		if err != nil {
			logger.Error(errors.ErrorStack(err))
			return err
		}

		logger.Debug("all events called")

		for _, inout := range inouts {
			if inout.Err != nil {
				err := errors.Wrap(inout.Err, errors.New("closing with errors lambda handle"))
				logger.Error(errors.ErrorStack(err))
				return err
			}
		}

	}

	logger.Info("closing lambda handle")

	return nil
}

func (h *Handler) getInOuts(ctx context.Context, event Event) ([]*cloudevents.InOut, error) {

	logger := log.FromContext(ctx)

	var inouts []*cloudevents.InOut

	if len(event.Records) > 0 {
		eventSource := event.Records[0].EventSource
		switch eventSource {
		case "aws:kinesis":
			inouts = fromKinesis(ctx, event)
		case "aws:sqs":
			inouts = fromSQS(ctx, event)
		case "aws:sns":
			inouts = fromSNS(ctx, event)
		case "aws:s3":
			inouts = fromS3(ctx, event)
		case "aws:dynamodb":
			inouts = fromDynamoDB(ctx, event)
		default:
			return nil, errors.NotImplementedf("the trigger received has not yet been implemented")
		}

	} else {

		if event.Source == "aws.events" {
			inouts = fromCloudWatch(ctx, event)
		} else {
			logger.Warnf("ignoring trigger")
			return nil, nil
		}

	}

	return inouts, nil
}
