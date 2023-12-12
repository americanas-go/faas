package pubsub

import (
	"github.com/americanas-go/config"
)

const (
	root                   = "faas.pubsub"
	projectid              = root + ".projectid"
	subscription           = root + ".subscription"
	numgoroutines          = root + ".numgoroutines"
	maxoutstandingmessages = root + ".maxoutstandingmessages"
)

func init() {
	config.Add(projectid, "changeme", "pubsub project id")
	config.Add(subscription, "changeme", "pubsub listener topics")
	config.Add(numgoroutines, 16, "pubsub num go routines")
	config.Add(maxoutstandingmessages, 8, "is the maximum number of unprocessed messages")

}
