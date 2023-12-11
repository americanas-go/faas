package pubsub

import (
	"github.com/americanas-go/config"
)

const (
	root         = "faas.pubsub"
	projectid    = root + ".projectid"
	subscription = root + ".subscription"
	concurrency  = root + ".concurrency"
)

func init() {
	config.Add(projectid, "changeme", "pubsub project id")
	config.Add(subscription, "changeme", "pubsub listener topics")
	config.Add(concurrency, 10, "pubsub goroutine concurrency")
}
