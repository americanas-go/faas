package publisher

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/faas/cloudevents"
)

const (
	root    = cloudevents.ExtRoot + ".publisher"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable publisher middleware")
}

// IsEnabled reports whether publisher middleware is enabled in the configuration.
func IsEnabled() bool {
	return config.Bool(enabled)
}
