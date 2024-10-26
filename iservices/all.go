package iservices

import (
	"github.com/reizt/rest-go/iservices/idatabase"
	"github.com/reizt/rest-go/iservices/igreeter"
	"github.com/reizt/rest-go/iservices/ihasher"
	"github.com/reizt/rest-go/iservices/imailer"
	"github.com/reizt/rest-go/iservices/isigner"
)

type All struct {
	Greeter  igreeter.Service
	Database idatabase.Service
	Hasher   ihasher.Service
	Mailer   imailer.Service
	Signer   isigner.Service
}
