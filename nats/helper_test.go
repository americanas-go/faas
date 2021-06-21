package nats

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/americanas-go/config"
	"github.com/americanas-go/faas/cloudevents"
	iglog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/nats-io/nats.go.v1"

	"github.com/stretchr/testify/suite"
)

type NatsHelperSuite struct {
	suite.Suite
}

func TestNatsHelperSuite(t *testing.T) {
	suite.Run(t, new(NatsHelperSuite))
}

func (s *NatsHelperSuite) SetupSuite() {
	config.Load()
	iglog.New()
}

func (s *NatsHelperSuite) TestNatsNewHelper() {

	ctx := context.Background()
	defaultOptions, _ := DefaultOptions()

	sUrl := fmt.Sprintf("nats://127.0.0.1:%d", TestPort)
	options, _ := nats.NewOptions()
	options.Url = sUrl
	subscriber, _ := nats.NewSubscriberWithOptions(ctx, options)

	type args struct {
		ctx        context.Context
		subscriber *nats.Subscriber
		options    *Options
		handler    *cloudevents.HandlerWrapper
	}
	tests := []struct {
		name string
		args args
		want *Helper
	}{
		{
			name: "success",
			args: args{
				ctx:        ctx,
				subscriber: subscriber,
				options:    defaultOptions,
				handler:    nil,
			},
			want: &Helper{nil, "changeme", []string{"changeme"}, subscriber},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewHelper(tt.args.ctx, tt.args.subscriber, tt.args.options, tt.args.handler)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewHelper() = %v, want %v")
		})
	}
}

func (s *NatsHelperSuite) TestNatsNewDefaultHelper() {

	ctx := context.Background()
	defaultOptions, _ := DefaultOptions()

	sUrl := fmt.Sprintf("nats://127.0.0.1:%d", TestPort)
	options, _ := nats.NewOptions()
	options.Url = sUrl
	subscriber, _ := nats.NewSubscriberWithOptions(ctx, options)

	type args struct {
		ctx        context.Context
		subscriber *nats.Subscriber
		options    *Options
		handler    *cloudevents.HandlerWrapper
	}
	tests := []struct {
		name string
		args args
		want *Helper
	}{
		{
			name: "success",
			args: args{
				ctx:        ctx,
				subscriber: subscriber,
				options:    defaultOptions,
				handler:    nil,
			},
			want: &Helper{nil, "changeme", []string{"changeme"}, subscriber},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewDefaultHelper(tt.args.ctx, tt.args.subscriber, tt.args.handler)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewHelper() = %v, want %v")
		})
	}
}