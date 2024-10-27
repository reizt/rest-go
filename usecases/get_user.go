package usecases

import (
	"context"

	"github.com/reizt/rest-go/iservices"
	i "github.com/reizt/rest-go/iusecases"
)

func GetUser(s *iservices.All, auth authenticator) i.GetUser {
	return func(input i.GetUserInput, ctx context.Context) (*i.GetUserOutput, error) {
		user, err := auth(input.LoginToken, ctx)
		if err != nil {
			return nil, err
		}

		output := i.GetUserOutput{
			User: user.ToEntity(),
		}
		return &output, nil
	}
}
