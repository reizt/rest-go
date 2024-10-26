package mailer

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/reizt/rest-go/iservices/imailer"
)

type Service struct {
	apiKey string
	from   string
}

func New(apiKey string, from string) *Service {
	return &Service{
		apiKey: apiKey,
		from:   from,
	}
}

func (s *Service) Send(input imailer.SendInput) error {
	from := mail.NewEmail("REST Go", s.from)
	to := mail.NewEmail(input.To, input.To)
	message := mail.NewSingleEmail(from, input.Subject, to, input.Text, input.Html)
	client := sendgrid.NewSendClient(s.apiKey)
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(response)
	return nil
}
