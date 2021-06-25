package mock

import "github.com/americanas-go/faas/repository"

// NewEvent returns a initialized mock that implements an event repository.
func NewEvent() repository.Event {
	return NewMock()
}
