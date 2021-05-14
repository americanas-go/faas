package lambda

import (
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type Record struct {
	EventVersion           string                                `json:"eventVersion"`
	EventSubscriptionArn   string                                `json:"eventSubscriptionArn"`
	EventSource            string                                `json:"eventSource"`
	EventName              string                                `json:"eventName"`
	EventID                string                                `json:"eventID"`
	SNS                    events.SNSEntity                      `json:"sns"`
	S3                     events.S3EventRecord                  `json:"s3"`
	Kinesis                events.KinesisRecord                  `json:"kinesis"`
	MessageId              string                                `json:"messageId"`
	ReceiptHandle          string                                `json:"receiptHandle"`
	Body                   string                                `json:"body"`
	Md5OfBody              string                                `json:"md5OfBody"`
	Md5OfMessageAttributes string                                `json:"md5OfMessageAttributes"`
	Attributes             map[string]string                     `json:"attributes"`
	MessageAttributes      map[string]events.SQSMessageAttribute `json:"messageAttributes"`
	EventSourceARN         string                                `json:"eventSourceARN"`
	AWSRegion              string                                `json:"awsRegion"`
}

type Event struct {
	ID         string    `json:"id"`
	Source     string    `json:"source"`
	Region     string    `json:"region"`
	DetailType string    `json:"detail-type"`
	Time       time.Time `json:"time"`
	Account    string    `json:"account"`
	Resources  []string  `json:"resources"`
	Records    []Record  `json:"Records"`
}
