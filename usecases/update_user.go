package usecases

import (
	"context"
	"fmt"

	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/iservices/idatabase"
	i "github.com/reizt/rest-go/iusecases"
)

func UpdateUser(s *iservices.All) i.UpdateUser {
	auth := authenticator(s)
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

		// Return
		output := i.UpdateUserOutput{}
		return &output, nil
	}
}
