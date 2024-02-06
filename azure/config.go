package azure

import (
	"github.com/americanas-go/config"
)

const (
	root = "faas.azure"
	port = root + ".port"
	name = root + ".name"
)

func init() {
	config.Add(port, "7071", "define http port")
	config.Add(name, "handler", "define name")
}
