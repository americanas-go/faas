package logger

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/log"
	v2 "github.com/cloudevents/sdk-go/v2"
)

type Logger struct {
	cloudevents.UnimplementedMiddleware
	level string
}

func NewLogger() cloudevents.Middleware {
	if !IsEnabled() {
		return nil
	}
	return &Logger{level: Level()}
}

func (m *Logger) Before(ctx context.Context, in *v2.Event) (context.Context, error) {
	logger := log.FromContext(ctx).WithTypeOf(*m)

	lm := m.logger(logger)

	lm("received event")

	j, err := json.Marshal(in)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
	} else {
		lm(string(j))
	}

	return ctx, nil
}

func (m *Logger) After(ctx context.Context, in v2.Event, out *v2.Event, err error) (context.Context, error) {

	logger := log.FromContext(ctx).WithTypeOf(*m)

	if out != nil && err == nil {

		lm := m.logger(logger)

		lm("returning event")

		j, err := json.Marshal(out)
		if err != nil {
			logger.Error(errors.ErrorStack(err))
		} else {
			lm(string(j))
		}

	}

	if err != nil {
		logger.Error(errors.ErrorStack(err))
	}

	return ctx, nil
}

func (m *Logger) logger(logger log.Logger) func(format string, args ...interface{}) {

	var method func(format string, args ...interface{})

	switch m.level {
	case "TRACE":
		method = logger.Tracef
	case "DEBUG":
		method = logger.Debugf
	default:
		method = logger.Infof
	}

	return method
}
