package services

import (
	"fmt"

	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/services/database"
	"github.com/reizt/rest-go/services/greeter"
	"github.com/reizt/rest-go/services/hasher"
	"github.com/reizt/rest-go/services/mailer"
	"github.com/reizt/rest-go/services/signer"
)

func New() (*iservices.All, error) {
	database, err := database.New()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to init database")
	}

	mailer, err := mailer.New()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to init mailer")
	}

	signer, err := signer.New()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to init signer")
	}

	return &iservices.All{
		Greeter:  greeter.New(),
		Database: *database,
		Hasher:   hasher.New(),
		Mailer:   mailer,
		Signer:   signer,
	}, nil
}
