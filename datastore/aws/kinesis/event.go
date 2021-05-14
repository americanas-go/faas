package kinesis

import (
	"github.com/americanas-go/ignite/aws/aws-sdk-go.v2/client/kinesis"
	"github.com/americanas-go/serverless/repository"
)

// NewEvent returns a initialized client
func NewEvent(c kinesis.Client, options *Options) repository.Event {
	return NewClient(c, options)
}
