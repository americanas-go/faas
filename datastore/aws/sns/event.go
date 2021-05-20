package sns

import (
	"github.com/americanas-go/faas/repository"
	"github.com/americanas-go/ignite/aws/aws-sdk-go.v2/client/sns"
)

// NewEvent returns a initialized client
func NewEvent(c sns.Client) repository.Event {
	return NewClient(c)
}
