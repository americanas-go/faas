faas
=======

An easy way to develop event-driven functions with modular approach using dependency injection.

The faas lets you write lightweight functions that run in different environments, including:

*   Your [local](https://github.com/nats-io/nats.go) development machine
*   [AWS Lambda Functions](https://github.com/aws/aws-lambda-go)
*   [Knative Eventing](https://github.com/knative/eventing)

Installation
------------

	go get -u github.com/americanas-go/faas

Example
--------
```go
package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/faas/cloudevents/plugins/contrib/americanas-go/log.v1"
	"github.com/americanas-go/faas/cmd"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	igce "github.com/americanas-go/ignite/cloudevents/sdk-go.v2"
	v2 "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/fx"
)

func main() {

	config.Load()

	ilog.New()

	options := fx.Options(
		fx.Provide(
			func() igce.Handler {
				return Handle
			},
			func() []cloudevents.Middleware {
				return []cloudevents.Middleware{
					log.NewLogger(),
				}
			},
		),
	)

	// sets env var
	os.Setenv("FAAS_CMD_DEFAULT", "nats")
	// os.Setenv("FAAS_CMD_DEFAULT", "cloudevents")
	// os.Setenv("FAAS_CMD_DEFAULT", "lambda")

	// go run main.go help
	err := cmd.Run(options,

		// go run main.go nats
		// or
		// FAAS_CMD_DEFAULT=nats go run main.go
		cmd.NewNats(),

		// go run main.go cloudevents
		// or
		// FAAS_CMD_DEFAULT=cloudevents go run main.go
		cmd.NewCloudEvents(),

		// go run main.go lambda
		// or
		// FAAS_CMD_DEFAULT=lambda go run main.go
		cmd.NewLambda(),
	)

	if err != nil {
		panic(err)
	}
}

func Handle(ctx context.Context, in v2.Event) (*v2.Event, error) {
	e := v2.NewEvent()
	e.SetID("changeme")
	e.SetSubject("changeme")
	e.SetSource("changeme")
	e.SetType("changeme")
	e.SetExtension("partitionkey", "changeme")
	err := e.SetData(v2.TextPlain, "changeme")

	return &e, err
}
```

Example of nats client
--------
```go
package main

import (
	"context"
	"io/ioutil"

	"github.com/americanas-go/config"
	iglog "github.com/americanas-go/ignite/americanas-go/log.v1"
	ignats "github.com/americanas-go/ignite/nats-io/nats.go.v1"
	"github.com/americanas-go/log"
	"github.com/nats-io/nats.go"
)

func main() {

	config.Load()
	iglog.New()

	var err error
	var conn *nats.Conn

	conn, err = ignats.NewConn(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	subject := "changeme"

	var b []byte
	b, err = ioutil.ReadFile("examples/simple/client/example-nats.json")
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
```

Example of sns client
--------
```go
package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/americanas-go/config"
	iglog "github.com/americanas-go/ignite/americanas-go/log.v1"
	igaws "github.com/americanas-go/ignite/aws/aws-sdk-go.v2"
	"github.com/americanas-go/log"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go/aws"
)

type Message struct {
	Default string `json:"default"`
}

func main() {

	config.Load()
	iglog.New()

	var err error

	ctx := context.Background()

	awsConfig := igaws.NewConfig(ctx)
	// if you have already set aws credentials in your system environment variables,
	// ignore the two lines below
	awsConfig.Region = "YOUR_AWS_REGION"
	awsConfig.Credentials = credentials.
		NewStaticCredentialsProvider("YOUR_AWS_ACCESS_KEY_ID", "YOUR_AWS_SECRET_ACCESS_KEY", "")

	client := sns.NewFromConfig(awsConfig)

	topic := "arn:aws:sns:us-east-1:000000000000:changeme"

	var b []byte
	b, err = ioutil.ReadFile("examples/simple/client/example-sns.json")
	if err != nil {
		log.Fatal(err)
	}

	msg := Message{
		Default: string(b),
	}
	msgBytes, _ := json.Marshal(msg)
	msgStr := string(msgBytes)

	res, err := client.Publish(ctx, &sns.PublishInput{
		Message:          aws.String(msgStr),
		MessageStructure: aws.String("json"),
		TopicArn:         aws.String(topic),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Infof("published group message on topic [%s]", topic)

	resJSON, _ := json.Marshal(res)

	log.Info(string(resJSON))
}
```

Contributing
--------
Every help is always welcome. Feel free do throw us a pull request, we'll do our best to check it out as soon as possible. But before that, let us establish some guidelines:

1. This is an open source project so please do not add any proprietary code or infringe any copyright of any sort.
2. Avoid unnecessary dependencies or messing up go.mod file.
3. Be aware of golang coding style. Use a lint to help you out.
4.  Add tests to cover your contribution.
5. Add [godoc](https://elliotchance.medium.com/godoc-tips-tricks-cda6571549b) to your code. 
6. Use meaningful [messages](https://medium.com/@menuka/writing-meaningful-git-commit-messages-a62756b65c81) to your commits.
7. Use [pull requests](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests).
8. At last, but also important, be kind and polite with the community.

Any submitted issue which disrespect one or more guidelines above, will be discarded and closed.


<hr>

Released under the [MIT License](LICENSE).
