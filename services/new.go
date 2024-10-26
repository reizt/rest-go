package services

import (
	"os"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/services/database"
	"github.com/reizt/rest-go/services/greeter"
	"github.com/reizt/rest-go/services/hasher"
	"github.com/reizt/rest-go/services/mailer"
	"github.com/reizt/rest-go/services/signer"
)

type Env struct {
	JwtPrivateKey  string
	JwtPublicKey   string
	SendgridApiKey string
	MailerFrom     string
}

func (e Env) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.JwtPrivateKey, v.Required, v.Length(1, 10000)),
		v.Field(&e.JwtPublicKey, v.Required, v.Length(1, 10000)),
		v.Field(&e.SendgridApiKey, v.Required, v.Length(1, 10000)),
		v.Field(&e.MailerFrom, v.Required, v.Length(1, 10000)),
	)
}

func New() (*iservices.All, error) {
	env := Env{
		JwtPrivateKey:  os.Getenv("JWT_PRIVATE_KEY"),
		JwtPublicKey:   os.Getenv("JWT_PUBLIC_KEY"),
		SendgridApiKey: os.Getenv("SENDGRID_API_KEY"),
		MailerFrom:     os.Getenv("MAILER_FROM"),
	}
	if err := env.Validate(); err != nil {
		return nil, err
	}

	signer, err := signer.New(env.JwtPrivateKey, env.JwtPublicKey)
	if err != nil {
		return nil, err
	}

	return &iservices.All{
		Greeter:  greeter.New(),
		Database: database.New(),
		Hasher:   hasher.New(),
		Mailer:   mailer.New(env.SendgridApiKey, env.MailerFrom),
		Signer:   signer,
	}, nil
}
