package services

import (
	"fmt"

	"reij.uno/iservices"
	"reij.uno/services/database"
	"reij.uno/services/greeter"
	"reij.uno/services/hasher"
	"reij.uno/services/mailer"
	"reij.uno/services/signer"
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
