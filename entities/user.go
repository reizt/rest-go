package entities

import v "github.com/go-ozzo/ozzo-validation/v4"

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u User) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.Id, v.Required, v.Length(1, 100)),
		v.Field(&u.Email, v.Required, v.Length(1, 100)),
		v.Field(&u.Name, v.Required, v.Length(1, 100)),
	)
}
