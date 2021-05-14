package repository

import "github.com/americanas-go/config"

type Options struct {
	Provider string
}

func DefaultEventOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(eventRoot, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
