package cmd

import (
	"github.com/americanas-go/config"
)

const (
	root = "gs.cmd"
	def  = root + ".default"
)

func init() {
	config.Add(def, "", "default cmd")
}

func Default() string {
	return config.String(def)
}
