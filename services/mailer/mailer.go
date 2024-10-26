package mailer

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/reizt/rest-go/iservices/imailer"
)

type service struct {
	apiKey string
	from   string
}

func New() (imailer.Service, error) {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	from := os.Getenv("MAILER_FROM")
	if apiKey == "" || from == "" {
		return nil, fmt.Errorf("SENDGRID_API_KEY and MAILER_FROM must be set")
	}

	return &service{
		apiKey: apiKey,
		from:   from,
	}, nil
}

func (s *service) Send(input imailer.SendInput) error {
	from := mail.NewEmail("REST Go", s.from)
	to := mail.NewEmail(input.To, input.To)
	message := mail.NewSingleEmail(from, input.Subject, to, input.Text, input.Html)
	client := sendgrid.NewSendClient(s.apiKey)
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("SendGrid status:", response.StatusCode)
	return nil
}
