package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/iservices/idatabase"
	i "github.com/reizt/rest-go/iusecases"
	"github.com/reizt/rest-go/usecases/token"
)

func UpdateUser(s *iservices.All) i.UpdateUser {
	return func(input i.UpdateUserInput) (*i.UpdateUserOutput, error) {
		// Verify token
		payload, err := s.Signer.Verify(input.LoginToken)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("invalid token")
		}
		var loginTokenPayload token.LoginTokenPayload
		if err := json.Unmarshal([]byte(payload), &loginTokenPayload); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("invalid token")
		}

		// Get user
		user, err := s.Database.User.GetById(loginTokenPayload.UserId)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("user not found")
		}

		// Update user
		newUser := idatabase.User{
			Id:           user.Id,
			Email:        user.Email,
			Name:         input.Data.Name,
			PasswordHash: user.PasswordHash,
		}
		if err := s.Database.User.Update(newUser); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to update user")
		}

		// Return
		output := i.UpdateUserOutput{}
		return &output, nil
	}
}
