package repository

import (
	"github.com/americanas-go/config"
)

const (
	root          = "faas.repository"
	eventRoot     = root + ".event"
	eventProvider = eventRoot + ".provider"
)

func init() {
	config.Add(eventProvider, "mock", "event provider")
}

// EventProviderValue returns the event provider configured via the "faas.repository.event.provider" key.
// If not configured, the default is mock.
func EventProviderValue() string {
	return config.String(eventProvider)
}
