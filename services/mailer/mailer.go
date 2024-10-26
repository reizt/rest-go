package mailer

import "github.com/reizt/rest-go/iservices/imailer"

func New() imailer.Service {
	return Service{}
}

type Service struct{}

func (m Service) Send(input imailer.SendInput) error {
	return nil
}
