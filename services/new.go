package services

import (
	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/services/database"
	"github.com/reizt/rest-go/services/greeter"
	"github.com/reizt/rest-go/services/hasher"
	"github.com/reizt/rest-go/services/mailer"
	"github.com/reizt/rest-go/services/signer"
)

func New() *iservices.All {
	return &iservices.All{
		Greeter:  greeter.New(),
		Database: database.New(),
		Hasher:   hasher.New(),
		Mailer:   mailer.New(),
		Signer:   signer.New(),
	}
}
