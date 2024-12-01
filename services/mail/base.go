package mail

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type MailService struct {
	ses *ses.Client
}

type MailInterface interface {

	// SendEmail
	/**
	 * Function to send an email based on the
	 * specified parameters.
	 * 	- receiver: The receiver of the email
	 *	- subject: The subject of the mail
	 *	- html: The body of the email
	 */
	SendEmail(context context.Context, source string, receiver string, subject string, html string) error
}

func New(context context.Context, region string) MailInterface {

	cfg, err := config.LoadDefaultConfig(context, config.WithRegion(region))

	if err != nil {
		panic(err)
	}

	return &MailService{
		ses: ses.NewFromConfig(cfg),
	}
}

func NewTest() MailInterface {
	return Mock{}
}

func (s *MailService) SendEmail(context context.Context, source string, receiver string, subject string, body string) error {

	_, err := s.ses.SendEmail(context, &ses.SendEmailInput{
		Source: aws.String(source),
		Destination: &types.Destination{
			ToAddresses: []string{
				receiver,
			},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
