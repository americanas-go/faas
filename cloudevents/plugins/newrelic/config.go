package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/serverless/cloudevents"
)

const (
	root    = cloudevents.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
	txName  = root + ".txName"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic middleware")
	config.Add(txName, "changeme", "cloudevents newrelic middleware tx name")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func TxName() string {
	return config.String(txName)
}
