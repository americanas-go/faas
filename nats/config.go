package nats

import (
	"github.com/americanas-go/config"
)

const (
	root     = "faas.nats"
	subjects = root + ".subjects"
	queue    = root + ".queue"
)

func init() {
	config.Add(subjects, []string{"changeme"}, "nats listener subjects")
	config.Add(queue, "changeme", "nats listener queue")
}

func SubjectsValue() []string {
	return config.Strings(subjects)
}

func QueueValue() string {
	return config.String(queue)
}
