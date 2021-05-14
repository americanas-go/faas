package publisher

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/serverless/cloudevents"
)

const (
	root    = cloudevents.ExtRoot + ".publisher"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable publisher middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
