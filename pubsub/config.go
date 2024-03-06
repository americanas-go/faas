package pubsub

import (
	"github.com/americanas-go/config"
)

const (
	root        = "faas.pubsub"
	topic       = root + ".topic"
	concurrency = root + ".concurrency"
)

func init() {
	config.Add(topic, "changeme", "pubsub listener topics")
	config.Add(concurrency, 10, "pubsub goroutine concurrency")
}
