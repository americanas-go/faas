package repository

import (
	"github.com/americanas-go/config"
)

const (
	root          = "serverless.repository"
	eventRoot     = root + ".event"
	eventProvider = eventRoot + ".provider"
)

func init() {
	config.Add(eventProvider, "mock", "event provider")
}

func EventProviderValue() string {
	return config.String(eventProvider)
}
