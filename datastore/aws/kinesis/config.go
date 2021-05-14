package kinesis

import (
	"github.com/americanas-go/config"
)

const (
	root               = "gostack.provider.kinesis"
	randomPartitionKey = root + ".randompartitionkey"
)

func init() {
	config.Add(randomPartitionKey, false, "ramdomize partition key")
}

func RandomPartitionKeyValue() bool {
	return config.Bool(randomPartitionKey)
}
