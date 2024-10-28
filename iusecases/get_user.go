package iusecases

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"reij.uno/entities"
)

type GetUserInput struct {
	LoginToken string
}

func (i GetUserInput) Validate() error {
	return v.ValidateStruct(&i,
		v.Field(&i.LoginToken, v.Required, v.Length(1, 10000)),
	)
}

type GetUserOutput struct {
	User entities.User
}

type GetUser func(GetUserInput, context.Context) (*GetUserOutput, error)
