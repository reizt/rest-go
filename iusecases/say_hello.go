package iusecases

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type SayHelloInput struct {
	Name string `json:"name"`
}

func (i SayHelloInput) Validate() error {
	return v.ValidateStruct(&i,
		v.Field(&i.Name, v.Required, v.Length(1, 100)),
	)
}

type SayHelloOutput struct {
	Message string `json:"message"`
}

type SayHello func(SayHelloInput) (SayHelloOutput, error)