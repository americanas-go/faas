package lambda

import (
	"github.com/americanas-go/config"
)

const (
	root = "serverless.lambda"
	skip = root + ".skip"
)

func init() {
	config.Add(skip, false, "skip all triggers")
}
