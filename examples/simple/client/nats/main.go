package main

import (
	"context"
	"io/ioutil"

	"github.com/americanas-go/config"
	ginats "github.com/americanas-go/ignite/nats-io/nats.go.v1"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
	"github.com/americanas-go/log"
	"github.com/nats-io/nats.go"
)

func main() {

	config.Load()

	logrus.NewLogger()

	var err error
	var conn *nats.Conn

	conn, err = ginats.NewDefaultConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	subject := "changeme"

	var b []byte
	b, err = ioutil.ReadFile("examples/simple/client/example.json")
	if err != nil {
		log.Fatal(err)
	}

	msg := &nats.Msg{
		Subject: subject,
		Data:    b,
	}

	err = conn.PublishMsg(msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("published group message on subject [%s]", subject)
}
