package mail

import "context"

type Mock struct{}

func (s Mock) SendEmail(context context.Context, source string, receiver string, subject string, html string) error {
	return nil
}
