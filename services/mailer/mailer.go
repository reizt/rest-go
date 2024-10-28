package mailer

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"reij.uno/iservices/imailer"
)

type service struct {
	apiKey string
	from   string
}

func New() (*service, error) {
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

type sendInput struct {
	To      string
	Subject string
	Text    string
	Html    string
}

func (s service) send(input sendInput) error {
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

func (s service) Code(input imailer.CodeInput) error {
	sendInput := sendInput{
		To:      input.To,
		Subject: "Your code",
		Text:    input.Code,
		Html:    fmt.Sprintf("Your code is <code>%s</code>", input.Code),
	}
	return s.send(sendInput)
}
