package iusecases

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdatePasswordInput struct {
	OTPToken string
	Password string
}

func (i UpdatePasswordInput) Validate() error {
	return v.ValidateStruct(&i,
		v.Field(&i.OTPToken, v.Required, v.Length(1, 10000)),
		v.Field(&i.Password, v.Required, v.Length(8, 100)),
	)
}

type UpdatePasswordOutput struct {
	LoginToken string
}

type UpdatePassword func(UpdatePasswordInput, context.Context) (*UpdatePasswordOutput, error)
