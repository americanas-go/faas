package repository

import "github.com/americanas-go/config"

// Options represents loaded repository configurations.
type Options struct {
	Provider string
}

// DefaultEventOptions loads default event options.
func DefaultEventOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(eventRoot, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
