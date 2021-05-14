package nats

import (
	"context"
	"fmt"
	"testing"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/nats-io/nats.go.v1"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
	v2 "github.com/cloudevents/sdk-go/v2"

	"github.com/nats-io/gnatsd/server"
	natsserver "github.com/nats-io/nats-server/test"

	"github.com/stretchr/testify/assert"
)

const TestPort = 8369

type Handler struct {
}

func (h *Handler) Handle(ctx context.Context, in v2.Event) (*v2.Event, error) {
	return nil, nil
}

func runServerOnPort(port int) *server.Server {
	opts := natsserver.DefaultTestOptions
	opts.Port = port
	return runServerWithOptions(&opts)
}

func runServerWithOptions(opts *server.Options) *server.Server {
	return natsserver.RunServer(opts)
}

func TestSubscriberListenerSubscribe(t *testing.T) {

	config.Load()
	logrus.NewLogger()

	var err error
	var options *nats.Options

	s := runServerOnPort(TestPort)
	defer s.Shutdown()

	sUrl := fmt.Sprintf("nats://127.0.0.1:%d", TestPort)

	options, err = nats.DefaultOptions()
	assert.Nil(t, err)

	options.Url = sUrl

	conn, err := nats.NewConnection(context.Background(), options)
	assert.Nil(t, err)
	defer conn.Close()

	q, err := nats.NewSubscriber(context.Background(), options)
	assert.Nil(t, err)

	lis := NewSubscriberListener(q, nil, "subject", "queue")
	subscribe, err := lis.Subscribe(context.Background())
	assert.Nil(t, err)

	assert.True(t, subscribe.IsValid())

	err = subscribe.Unsubscribe()
	assert.Nil(t, err)
}
