package publisher

import (
	"github.com/americanas-go/ignite"
)

type Options struct {
	Enabled bool
	Success struct {
		Enabled bool
	}
	Error struct {
		Enabled bool
		Topic   string
	}
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}
