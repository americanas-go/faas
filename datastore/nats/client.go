package nats

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/faas/util"
	"github.com/americanas-go/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
)

// Client represents a NATS client that implements the event repository.
type Client struct {
	conn *nats.Conn
}

// NewClient creates a new NATS client.
func NewClient(conn *nats.Conn) *Client {
	return &Client{conn: conn}
}

// Publish publishes an event slice.
func (p *Client) Publish(ctx context.Context, outs []*v2.Event) (err error) {

	logger := log.FromContext(ctx).WithTypeOf(*p)

	logger.Info("publishing to nats")

	for _, out := range outs {

		logger = logger.
			WithField("subject", out.Subject()).
			WithField("id", out.ID())

		var rawMessage []byte

		exts := out.Extensions()

		source, ok := exts["target"]

		if ok {

			s := source.(string)

			if s == "data" {
				var data interface{}

				err = out.DataAs(&data)
				if err != nil {
					return errors.Wrap(err, errors.Internalf("error on data as. %s", err.Error()))
				}

				rawMessage, err = json.Marshal(data)

			} else {
				rawMessage, err = util.JSONBytes(*out)
			}

		} else {
			rawMessage, err = util.JSONBytes(*out)
		}

		if err != nil {
			err = errors.Wrap(err, errors.Internalf("error when transforming json into bytes"))
			logger.Error(errors.ErrorStack(err))
			continue
		}

		logger.Info(string(rawMessage))

		msg := &nats.Msg{
			Subject: out.Subject(),
			Data:    rawMessage,
		}

		err = p.conn.PublishMsg(msg)
		if err != nil {
			err = errors.Wrap(err, errors.Internalf("unable to publish to nats"))
			logger.Error(errors.ErrorStack(err))
		}

	}

	return nil
}
