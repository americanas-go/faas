package lambda

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/americanas-go/config"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/faas/cloudevents/plugins/logger"
	"github.com/americanas-go/faas/mocks"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
	"github.com/americanas-go/log"
	"github.com/aws/aws-lambda-go/lambdacontext"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/stretchr/testify/mock"
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
	logrus.NewLogger()
}

func (s *HandlerSuite) TestHandler_Handle() {

	handler := new(mocks.Handler)

	lc := new(lambdacontext.LambdaContext)
	ctx := lambdacontext.NewContext(context.Background(), lc)

	var kinesisEvent1 Event
	b, _ := ioutil.ReadFile("testdata/kinesis_success.json")
	json.Unmarshal(b, &kinesisEvent1)

	var middlewares []cloudevents.Middleware

	middlewares = append(middlewares, logger.NewLogger())

	options, _ := DefaultOptions()

	type fields struct {
		handler     *mocks.Handler
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
		mock    func(handler *mocks.Handler)
	}{
		{
			name: "on kinesis success event",
			fields: fields{
				handler:     handler,
				middlewares: middlewares,
				options:     options,
			},
			args: args{
				ctx:   ctx,
				event: kinesisEvent1,
			},
			wantErr: false,
			mock: func(handler *mocks.Handler) {

				e := v2.NewEvent()
				e.SetSubject("changeme")
				e.SetSource("changeme")
				e.SetType("changeme")
				e.SetData("", "changeme")

				handler.On("Handle", mock.Anything, mock.Anything).Times(1).
					Return(&e, nil)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {

			tt.mock(tt.fields.handler)

			hwOptions, _ := cloudevents.DefaultHandlerWrapperOptions()
			hw := cloudevents.NewHandlerWrapper(tt.fields.handler, hwOptions, tt.fields.middlewares...)
			h := NewHandler(hw, tt.fields.options)

			err := h.Handle(tt.args.ctx, tt.args.event)
			if err != nil {
				log.Error(err)
			}

			s.Assert().True((err != nil) == tt.wantErr, "Handle() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func (s *HandlerSuite) TestNewHandler() {

	type args struct {
		handler     cloudevents.Handler
		middlewares []cloudevents.Middleware
		options     *Options
	}

	handler := new(mocks.Handler)
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
				handler:     handler,
				middlewares: nil,
				options:     options,
			},
			want: NewHandler(hw, options),
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {

			hw := cloudevents.NewHandlerWrapper(tt.args.handler, hwOptions, tt.args.middlewares...)
			got := NewHandler(hw, tt.args.options)

			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewHandler() = %v, want %v")

		})
	}
}
