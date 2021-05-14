package kinesis

import (
	"github.com/americanas-go/config"
)

type Options struct {
	RandomPartitionKey bool `config:"randompartitionkey"`
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
