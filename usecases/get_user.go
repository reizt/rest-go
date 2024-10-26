package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/reizt/rest-go/iservices"
	i "github.com/reizt/rest-go/iusecases"
	"github.com/reizt/rest-go/usecases/token"
)

func GetUser(s *iservices.All) i.GetUser {
	return func(input i.GetUserInput) (*i.GetUserOutput, error) {
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

		// Return
		output := i.GetUserOutput{
			User: user.ToEntity(),
		}
		return &output, nil
	}
}
