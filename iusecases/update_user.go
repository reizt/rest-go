package iusecases

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateUserInputData struct {
	Name string
}

type UpdateUserInput struct {
	LoginToken string
	Data       UpdateUserInputData
}

func (i UpdateUserInput) Validate() error {
	return v.ValidateStruct(&i,
		v.Field(&i.LoginToken, v.Required, v.Length(1, 10000)),
		v.Field(&i.Data, v.By(func(datai interface{}) error {
			data := datai.(UpdateUserInputData)
			return v.ValidateStruct(&data,
				v.Field(&data.Name, v.Required, v.Length(1, 100)),
			)
		})),
	)
}

type UpdateUserOutput struct {
}

type UpdateUser func(UpdateUserInput, context.Context) (*UpdateUserOutput, error)
