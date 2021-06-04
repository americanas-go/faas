module github.com/americanas-go/faas

go 1.16

require (
	github.com/americanas-go/config v1.2.0
	github.com/americanas-go/errors v1.0.0
	github.com/americanas-go/ignite v1.0.0-alpha-16
	github.com/americanas-go/log v1.3.0
	github.com/americanas-go/utils v1.0.0
	github.com/aws/aws-lambda-go v1.13.3
	github.com/aws/aws-sdk-go-v2 v1.4.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.2.3
	github.com/aws/aws-sdk-go-v2/service/sns v1.3.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.3.2
	github.com/cloudevents/sdk-go/v2 v2.4.1
	github.com/google/uuid v1.1.2
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2
	github.com/nats-io/gnatsd v1.4.1
	github.com/nats-io/go-nats v1.7.2 // indirect
	github.com/nats-io/nats-server v1.4.1
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/nats-io/nats.go v1.11.0
	github.com/newrelic/go-agent/v3 v3.12.0
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	go.uber.org/fx v1.13.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

// replace github.com/americanas-go/ignite => ../ignite
