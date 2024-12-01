package sqssession

import (
	"context"
	"github.com/kostanios/gotemp/events"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type MessageHandler func(msg types.Message) error

type SqsSession interface {
	Publish(topic string, event events.Deserializable) error
	Subscribe(topic string, msgHandler MessageHandler)
}

type sqsSession struct {
	client *sqs.Client
}

func NewSqsSession(accessKeyID string, secretAccessKey string, region string) (SqsSession, error) {
	customResolver := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, ""))

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(customResolver),
	)

	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	return &sqsSession{client: client}, nil
}

func (ses *sqsSession) Publish(topic string, event events.Deserializable) error {
	bytes, err := event.Deserialize()
	if err != nil {
		return err
	}
	body := string(bytes)

	_, err = ses.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &topic,
		MessageBody: &body,
	})

	return err
}

func (ses *sqsSession) Subscribe(topic string, msgHandler MessageHandler) {
	for {
		result, err := ses.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl:            &topic,
			MaxNumberOfMessages: 10,
			VisibilityTimeout:   30,
			WaitTimeSeconds:     10,
		})

		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}

		for _, message := range result.Messages {
			err := msgHandler(message)
			if err != nil {
				go func() {
					ses.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
						QueueUrl:      &topic,
						ReceiptHandle: message.ReceiptHandle,
					})
				}()
			}
		}
	}
}
