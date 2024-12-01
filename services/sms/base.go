package sms

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

type Interface interface {
	SendSMS(context context.Context, message string, phone string) error
}

type SMSService struct {
	senderName *string
	snsClient *sns.Client
}

func New(context context.Context, region string, senderName *string) Interface {

	cfg, err := config.LoadDefaultConfig(context, config.WithRegion(region))

	if err != nil {
		panic(err)
	}

	return SMSService{
		senderName:senderName,
		snsClient:sns.NewFromConfig(cfg),
	}
}

func (s SMSService) SendSMS(context context.Context, message string, phone string) error {

	if len(message) <= 0 {
		return errors.New(ErrorNoContent)
	}

	if len(phone) <= 0 {
		return errors.New(ErrorMissingPhoneNumber)
	}

	input := sns.PublishInput{
		Message:aws.String(message),
		PhoneNumber:aws.String(phone),
	}

	if s.senderName != nil {
		input.MessageAttributes = map[string]types.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {
				StringValue:s.senderName,
				DataType:aws.String("String"),
			},
			"AWS.SNS.SMS.SMSType":{
				StringValue:aws.String("Transactional"),
				DataType:aws.String("String"),
			},
		}
	}

	_,err := s.snsClient.Publish(context, &input)

	if err != nil {
		return err
	}

	return nil
}