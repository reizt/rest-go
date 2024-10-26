package services

import (
	"fmt"
	"os"

	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/services/database"
	"github.com/reizt/rest-go/services/greeter"
	"github.com/reizt/rest-go/services/hasher"
	"github.com/reizt/rest-go/services/mailer"
	"github.com/reizt/rest-go/services/signer"
)

func New() (*iservices.All, error) {
	privateKey := os.Getenv("PRIVATE_KEY")
	publicKey := os.Getenv("PUBLIC_KEY")

	if privateKey == "" || publicKey == "" {
		return nil, fmt.Errorf("private key or public key is not set")
	}

	return &iservices.All{
		Greeter:  greeter.New(),
		Database: database.New(),
		Hasher:   hasher.New(),
		Mailer:   mailer.New(),
		Signer:   signer.New(privateKey, publicKey),
	}, nil
}
