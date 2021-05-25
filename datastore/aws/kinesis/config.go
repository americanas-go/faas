package kinesis

import (
	"github.com/americanas-go/config"
)

const (
	root               = "faas.provider.kinesis"
	randomPartitionKey = root + ".randomPartitionKey"
)

func init() {
	config.Add(randomPartitionKey, false, "ramdomize partition key")
}

func RandomPartitionKeyValue() bool {
	return config.Bool(randomPartitionKey)
}
