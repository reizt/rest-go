package usecases

import (
	"context"
	"fmt"

	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	i "reij.uno/iusecases"
)

func updateUser(s *iservices.All, auth authenticator) i.UpdateUser {
	return func(input i.UpdateUserInput, ctx context.Context) (*i.UpdateUserOutput, error) {
		user, err := auth(input.LoginToken, ctx)
		if err != nil {
			return nil, err
		}

		// Update user
		newUser := idatabase.User{
			Id:           user.Id,
			Email:        user.Email,
			Name:         input.Data.Name,
			PasswordHash: user.PasswordHash,
		}
		if err := s.Database.User.Update(newUser, ctx); err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}

		output := i.UpdateUserOutput{}
		return &output, nil
	}
}
