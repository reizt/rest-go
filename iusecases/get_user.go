package iusecases

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/reizt/rest-go/entities"
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

type GetUser func(GetUserInput) (*GetUserOutput, error)
