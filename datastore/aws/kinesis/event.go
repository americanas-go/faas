package kinesis

import (
	"github.com/americanas-go/faas/repository"
	"github.com/americanas-go/ignite/aws/aws-sdk-go.v2/client/kinesis"
)

// NewEvent returns a initialized client
func NewEvent(c kinesis.Client, options *Options) repository.Event {
	return NewClient(c, options)
}
