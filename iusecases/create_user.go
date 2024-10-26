package iusecases

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateUserInput struct {
	OTPToken string
	Name     string
	Password string
}

func (i CreateUserInput) Validate() error {
	return v.ValidateStruct(&i,
		v.Field(&i.OTPToken, v.Required, v.Length(1, 10000)),
		v.Field(&i.Name, v.Required, v.Length(1, 100)),
		v.Field(&i.Password, v.Required, v.Length(8, 100)),
	)
}

type CreateUserOutput struct {
	LoginToken string
}

type CreateUser func(CreateUserInput) (*CreateUserOutput, error)
