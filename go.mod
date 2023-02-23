module github.com/americanas-go/faas

go 1.18

require (
	github.com/americanas-go/config v1.8.0
	github.com/americanas-go/errors v1.1.0
	github.com/americanas-go/ignite v1.1.0-beta.12
	github.com/americanas-go/log v1.8.5
	github.com/americanas-go/utils v1.1.0
	github.com/aws/aws-lambda-go v1.33.0
	github.com/aws/aws-sdk-go v1.44.63
	github.com/aws/aws-sdk-go-v2 v1.16.7
	github.com/aws/aws-sdk-go-v2/credentials v1.12.9
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.15.9
	github.com/aws/aws-sdk-go-v2/service/sns v1.17.9
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.0
	github.com/cloudevents/sdk-go/v2 v2.10.1
	github.com/google/uuid v1.3.0
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2
	github.com/nats-io/gnatsd v1.4.1
	github.com/nats-io/nats-server v1.4.1
	github.com/nats-io/nats.go v1.16.0
	github.com/newrelic/go-agent/v3 v3.17.0
	github.com/spf13/cobra v1.5.0
	github.com/stretchr/testify v1.8.0
	go.uber.org/fx v1.17.1
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.15.14 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.9 // indirect
	github.com/aws/smithy-go v1.12.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gobeam/stringy v0.0.5 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf v1.4.2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/go-nats v1.7.2 // indirect
	github.com/nats-io/jwt/v2 v2.2.1-0.20220113022732-58e87895b296 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ravernkoh/cwlogsfmt v0.0.0-20180121032441-917bad983b4c // indirect
	github.com/rs/zerolog v1.27.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/dig v1.14.1 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220726230323-06994584191e // indirect
	golang.org/x/sys v0.0.0-20220727055044-e65921a090b8 // indirect
	golang.org/x/text v0.3.8 // indirect
	golang.org/x/time v0.0.0-20220609170525-579cf78fd858 // indirect
	google.golang.org/genproto v0.0.0-20220725144611-272f38e5d71b // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace github.com/americanas-go/ignite => ../ignite
