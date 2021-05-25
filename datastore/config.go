package datastore

import (
	"github.com/americanas-go/config"
)

const (
	root          = "faas.datastore"
	eventProvider = root + ".event.provider"
)

func init() {
	config.Add(eventProvider, "nats", "event provider")
}

func EventProviderValue() string {
	return config.String(eventProvider)
}
