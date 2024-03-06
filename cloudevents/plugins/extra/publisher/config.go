package publisher

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/faas/cloudevents"
)

const (
	root           = cloudevents.PluginsRoot + ".publisher"
	enabled        = root + ".enabled"
	successEnabled = root + ".success.enabled"
	errorEnabled   = root + ".error.enabled"
	errorTopic     = root + ".error.topic"
)

func init() {
	config.Add(enabled, true, "enable/disable publisher middleware")
	config.Add(successEnabled, false, "enable/disable success publisher middleware")
	config.Add(errorEnabled, false, "enable/disable error publisher middleware")
	config.Add(errorTopic, "changeme", "sets error topic publisher middleware")
}
