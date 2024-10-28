package iservices

import (
	"reij.uno/iservices/idatabase"
	"reij.uno/iservices/igreeter"
	"reij.uno/iservices/ihasher"
	"reij.uno/iservices/imailer"
	"reij.uno/iservices/isigner"
)

type All struct {
	Greeter  igreeter.Service
	Database idatabase.Service
	Hasher   ihasher.Service
	Mailer   imailer.Service
	Signer   isigner.Service
}
