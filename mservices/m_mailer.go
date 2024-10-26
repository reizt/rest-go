package mservices

import "github.com/reizt/rest-go/iservices/imailer"

type Mailer struct {
	Code_ func(input imailer.CodeInput) error
}

func (m Mailer) Code(input imailer.CodeInput) error {
	return m.Code_(input)
}
