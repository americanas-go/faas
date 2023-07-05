package kafka

import (
	"context"
	"encoding/json"
	"github.com/americanas-go/errors"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/semaphore"
	"time"

	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/log"
)

// Helper assists in creating event handlers.
type Helper struct {
	handler *cloudevents.HandlerWrapper
	options *Options
}

// NewHelper returns a new Helper with options.
func NewHelper(ctx context.Context, options *Options,
	handler *cloudevents.HandlerWrapper) *Helper {

	return &Helper{
		handler: handler,
		options: options,
	}
}

// NewDefaultHelper returns a new Helper with default options.
func NewDefaultHelper(ctx context.Context, handler *cloudevents.HandlerWrapper) *Helper {

	opt, err := DefaultOptions()
	if err != nil {
		log.Fatal(err.Error())
	}

	return NewHelper(ctx, opt, handler)
}

func (h *Helper) Start() {

	for i := range h.options.Subjects {
		go h.subscribe(context.Background(), h.options.Subjects[i])
	}

	c := make(chan struct{})
	<-c
}

func (h *Helper) subscribe(ctx context.Context, topic string) {

	ctx = log.WithTypeOf(*h).
		WithField("topic", topic).
		WithField("groupId", h.options.GroupId).ToContext(ctx)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     h.options.Brokers,
		GroupID:     h.options.GroupId,
		Topic:       topic,
		Logger:      &Logger{},
		ErrorLogger: &ErrorLogger{},
		/*
			GroupTopics:            nil,
			Partition:              0,
			Dialer:                 nil,
			QueueCapacity:          0,
			MinBytes:               0,
			MaxBytes:               0,
			MaxWait:                0,
			ReadBatchTimeout:       0,
			ReadLagInterval:        0,
			GroupBalancers:         nil,
			HeartbeatInterval:      0,
			CommitInterval:         0,
			PartitionWatchInterval: 0,
			WatchPartitionChanges:  false,
			SessionTimeout:         0,
			RebalanceTimeout:       0,
			JoinGroupBackoff:       0,
			RetentionTime:          0,
			StartOffset:            0,
			ReadBackoffMin:         0,
			ReadBackoffMax:         0,
			Logger:                 nil,
			ErrorLogger:            nil,
			IsolationLevel:         0,
			MaxAttempts:            0,
			OffsetOutOfRangeError:  false,
		*/
	})

	var sem = semaphore.NewWeighted(int64(h.options.Concurrency))

	for {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Errorf(err.Error())
		}
		go func(ctx context.Context) {
			bctx := &contextWithoutDeadline{ctx}
			m, err := reader.ReadMessage(bctx)
			if err != nil {
				log.Errorf(err.Error())
			}
			h.handle(ctx, m)
			sem.Release(1)
		}(ctx)
	}

}

func (h *Helper) handle(ctx context.Context, msg kafka.Message) {

	logger := log.FromContext(ctx).WithTypeOf(*h)

	in := event.New()
	if err := json.Unmarshal(msg.Value, &in); err != nil {
		var data interface{}
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logger.Errorf("could not decode kafka record. %s", err.Error())
			return
		} else {
			err := in.SetData("", data)
			if err != nil {
				logger.Errorf("could set data from kafka record. %s", err.Error())
				return
			}
		}
	}

	var inouts []*cloudevents.InOut

	inouts = append(inouts, &cloudevents.InOut{In: &in})

	if err := h.handler.Process(ctx, inouts); err != nil {
		logger.Error(errors.ErrorStack(err))
	}

}

type contextWithoutDeadline struct {
	ctx context.Context
}

func (*contextWithoutDeadline) Deadline() (time.Time, bool) { return time.Time{}, false }
func (*contextWithoutDeadline) Done() <-chan struct{}       { return nil }
func (*contextWithoutDeadline) Err() error                  { return nil }

func (l *contextWithoutDeadline) Value(key interface{}) interface{} {
	return l.ctx.Value(key)
}
