package sqs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/faas/util"
	"github.com/americanas-go/ignite/aws/aws-sdk-go.v2/client/sqs"
	"github.com/americanas-go/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	awssqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/matryer/try"
	"golang.org/x/sync/errgroup"
)

// Client represents a sqs client.
type Client struct {
	client sqs.Client
}

// NewClient creates a new sqs client.
func NewClient(c sqs.Client) *Client {
	return &Client{client: c}
}

// Publish publishes an event slice.
func (p *Client) Publish(ctx context.Context, events []*v2.Event) error {

	logger := log.FromContext(ctx).WithTypeOf(*p)

	logger.Info("publishing to awssqs")

	if len(events) > 0 {

		return p.send(ctx, events)

	}

	logger.Warnf("no messages were reported for posting")

	return nil
}

// GetURL gets URL based on default aws configs, resource name and service as the pattern:
// https://\<service\>.\<region\>.amazonaws.com/\<account_number\>/\<resource_name\>
func (p *Client) getAwsUrl(resourceName, service string) string {
	//FIXME reference to fix this up: https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/sqs/sqs_getqueueurl.go
	//TODO move to ignite sqs client
	region := "fixme"        //giaws.Region()
	accountNumber := "fixme" //giaws.AccountNumber()
	return fmt.Sprintf("https://%s.%s.amazonaws.com/%s/%s", service, region, accountNumber, resourceName)
}

func (p *Client) send(parentCtx context.Context, events []*v2.Event) (err error) {

	logger := log.FromContext(parentCtx).WithTypeOf(*p)

	g, gctx := errgroup.WithContext(parentCtx)
	defer gctx.Done()

	for _, e := range events {

		event := e

		g.Go(func() (err error) {

			var rawMessage []byte

			rawMessage, err = p.rawMessage(event)
			if err != nil {
				return errors.Wrap(err, errors.Internalf("error on marshal. %s", err.Error()))
			}

			input := &awssqs.SendMessageInput{
				MessageBody: aws.String(string(rawMessage)),
				QueueUrl:    aws.String(p.getAwsUrl(event.Subject(), "awssqs")),
			}

			if group, ok := event.Extensions()["group"]; ok {
				input.MessageGroupId = aws.String(fmt.Sprintf("%v", group))
			}

			logger.WithField("subject", event.Subject()).
				WithField("id", event.ID()).
				Info(string(rawMessage))

			err = try.Do(func(attempt int) (bool, error) {
				var err error
				err = p.client.Publish(gctx, input)
				if err != nil {
					return attempt < 5, errors.NewInternal(err, "could not be published in awssqs")
				}
				return false, nil
			})

			return err

		})

	}

	return g.Wait()
}

func (p *Client) rawMessage(out *v2.Event) (rawMessage []byte, err error) {
	exts := out.Extensions()

	source, ok := exts["target"]

	if ok {

		s := source.(string)

		if s == "data" {
			var data interface{}

			err = out.DataAs(&data)
			if err != nil {
				return nil, errors.Wrap(err, errors.Internalf("error on data as. %s", err.Error()))
			}

			rawMessage, err = json.Marshal(data)

		} else {
			rawMessage, err = util.JSONBytes(*out)
		}
	} else {
		rawMessage, err = util.JSONBytes(*out)
	}

	return rawMessage, err
}
