package entities

import v "github.com/go-ozzo/ozzo-validation/v4"

const (
	CodeActionCreateUser    = "create-user"
	CodeActionResetPassword = "reset-password"
)

type Code struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Action    string `json:"action"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (c Code) Validate() error {
	return v.ValidateStruct(&c,
		v.Field(&c.Id, v.Required, v.Length(1, 100)),
		v.Field(&c.Email, v.Required, v.Length(1, 100)),
		v.Field(&c.Action, v.Required, v.In(CodeActionCreateUser, CodeActionResetPassword)),
		v.Field(&c.ExpiresAt, v.Required),
	)
}
