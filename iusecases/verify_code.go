package iusecases

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type VerifyCodeInput struct {
	CodeId string
	Code   string
}

func (i VerifyCodeInput) Validate() error {
	return v.ValidateStruct(&i,
		v.Field(&i.CodeId, v.Required),
		v.Field(&i.Code, v.Required),
	)
}

type VerifyCodeOutput struct {
	Token string
}

type VerifyCode func(VerifyCodeInput) (*VerifyCodeOutput, error)
