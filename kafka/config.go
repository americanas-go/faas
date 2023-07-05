package kafka

import (
	"github.com/americanas-go/config"
)

const (
	root        = "faas.kafka"
	topics      = root + ".topics"
	groupId     = root + ".groupId"
	brokers     = root + ".brokers"
	concorrency = root + ".concorrency"
)

func init() {
	config.Add(topics, []string{"changeme"}, "kafka listener topics")
	config.Add(brokers, []string{"localhost:9090"}, "kafka listener brokers")
	config.Add(groupId, "changeme", "kafka listener groupId")
	config.Add(concorrency, 5, "kafka listener concurrency")
}
