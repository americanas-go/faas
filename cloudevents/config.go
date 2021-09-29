package cloudevents

import (
	"github.com/americanas-go/config"
)

const (
	root                  = "faas.cloudevents"
	PluginsRoot           = root + ".plugins"
	handleDiscardEventsID = root + ".handle.discard.ids"
)

func init() {
	config.Add(handleDiscardEventsID, "", "cloudevents events id that will not be processed, comma separated")
}

// HandleDiscardEventsIDValue returns the event IDs to be discarded comma-separated from the configuration via "faas.cloudevents.handle.discard.ids" key.
func HandleDiscardEventsIDValue() string {
	return config.String(handleDiscardEventsID)
}
