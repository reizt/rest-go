package usecases

import (
	"context"

	"reij.uno/iservices"
	i "reij.uno/iusecases"
)

func getUser(_ *iservices.All, auth authenticator) i.GetUser {
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
