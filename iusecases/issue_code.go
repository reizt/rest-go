package iusecases

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IssueCodeInput struct {
	Email  string `json:"email"`
	Action string `json:"action"`
}

func (i IssueCodeInput) Validate() error {
	return v.ValidateStruct(&i,
		v.Field(&i.Email, v.Required, is.Email),
		v.Field(&i.Action, v.Required, v.In("create-user", "reset-password")),
	)
}

type IssueCodeOutput struct {
	CodeId string `json:"codeId"`
}

type IssueCode func(IssueCodeInput) (*IssueCodeOutput, error)
