package mock

import "github.com/americanas-go/serverless/repository"

// NewEvent returns a initialized mock
func NewEvent() repository.Event {
	return NewMock()
}
