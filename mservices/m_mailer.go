package mservices

import "reij.uno/iservices/imailer"

type Mailer struct {
	Code_ func(input imailer.CodeInput) error
}

func (m Mailer) Code(input imailer.CodeInput) error {
	return m.Code_(input)
}
