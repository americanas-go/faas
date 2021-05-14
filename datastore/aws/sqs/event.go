package sqs

import (
	giawsclientsqs "github.com/americanas-go/ignite/aws/aws-sdk-go.v2/client/sqs"
	"github.com/americanas-go/serverless/repository"
)

// NewEvent returns a initialized client
func NewEvent(c giawsclientsqs.Client) repository.Event {
	return NewClient(c)
}
