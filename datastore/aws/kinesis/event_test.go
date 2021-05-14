package kinesis

import (
	"reflect"
	"testing"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/aws/aws-sdk-go.v2/client/kinesis/mocks"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
	"github.com/americanas-go/serverless/repository"
	"github.com/stretchr/testify/suite"
)

type EventSuite struct {
	suite.Suite
}

func (s *EventSuite) SetupSuite() {
	config.Load()
	logrus.NewLogger()
}

func (s *EventSuite) TestNewEvent() {

	client := new(mocks.Client)
	options, _ := DefaultOptions()

	type args struct {
		client  *mocks.Client
		options *Options
	}
	tests := []struct {
		name string
		args args
		want repository.Event
	}{
		{
			name: "Success",
			args: args{
				client:  client,
				options: options,
			},
			want: NewEvent(client, options),
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewEvent(tt.args.client, tt.args.options)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewEvent() = %v, want %v", got, tt.want)
		})
	}
}

func TestEventSuite(t *testing.T) {
	suite.Run(t, new(EventSuite))
}
