package sqs

import (
	"github.com/americanas-go/faas/repository"
	giawsclientsqs "github.com/americanas-go/ignite/aws/aws-sdk-go.v2/client/sqs"
)

// NewEvent returns a initialized client
func NewEvent(c giawsclientsqs.Client) repository.Event {
	return NewClient(c)
}
