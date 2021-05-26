package lambda

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/americanas-go/config"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/faas/cloudevents/plugins/contrib/americanas-go/log.v1"
	iglog "github.com/americanas-go/ignite/americanas-go/log.v1"
	igcloudevents "github.com/americanas-go/ignite/cloudevents/sdk-go.v2"
	"github.com/aws/aws-lambda-go/lambdacontext"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) SetupSuite() {
	config.Load()
	iglog.New()
}

func (s *HandlerSuite) TestHandler_Handle() {

	lc := new(lambdacontext.LambdaContext)
	ctx := lambdacontext.NewContext(context.Background(), lc)

	var kinesisEvent1 Event
	b, _ := ioutil.ReadFile("testdata/kinesis_success.json")
	json.Unmarshal(b, &kinesisEvent1)

	var middlewares []cloudevents.Middleware

	middlewares = append(middlewares, log.NewLogger())

	options, _ := DefaultOptions()

	type fields struct {
		handler     igcloudevents.Handler
		middlewares []cloudevents.Middleware
		options     *Options
	}

	type args struct {
		ctx   context.Context
		event Event
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "on kinesis success event",
			fields: fields{
				handler: func(ctx context.Context, in v2.Event) (*v2.Event, error) {
					e := v2.NewEvent()
					e.SetSubject("changeme")
					e.SetSource("changeme")
					e.SetType("changeme")
					e.SetData("", "changeme")
					return &e, nil
				},
				middlewares: middlewares,
				options:     options,
			},
			args: args{
				ctx:   ctx,
				event: kinesisEvent1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {

			hwOptions, _ := cloudevents.DefaultHandlerWrapperOptions()
			hw := cloudevents.NewHandlerWrapper(tt.fields.handler, hwOptions, tt.fields.middlewares...)
			h := NewHandler(hw, tt.fields.options)

			err := h.Handle(tt.args.ctx, tt.args.event)
			s.Assert().True((err != nil) == tt.wantErr, "Handle() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func (s *HandlerSuite) TestNewHandler() {

	type args struct {
		handler igcloudevents.Handler
		options *Options
	}

	handler := func(ctx context.Context, in v2.Event) (*v2.Event, error) { return nil, nil }
	options, _ := DefaultOptions()
	hwOptions, _ := cloudevents.DefaultHandlerWrapperOptions()
	hw := cloudevents.NewHandlerWrapper(handler, hwOptions)

	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "success",
			args: args{
				handler: handler,
				options: options,
			},
			want: NewHandler(hw, options),
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got := NewHandler(hw, tt.args.options)

			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewHandler() = %v, want %v", got, tt.want)

		})
	}
}
