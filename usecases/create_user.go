package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/iservices/idatabase"
	i "github.com/reizt/rest-go/iusecases"
	"github.com/reizt/rest-go/usecases/id"
	"github.com/reizt/rest-go/usecases/token"
)

func CreateUser(s *iservices.All) i.CreateUser {
	return func(input i.CreateUserInput, ctx context.Context) (*i.CreateUserOutput, error) {
		// Verify token
		payload, err := s.Signer.Verify(input.OTPToken)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("invalid token")
		}
		var otpPayload token.OTPTokenPayload
		if err := json.Unmarshal([]byte(payload), &otpPayload); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("invalid token")
		}

		// Create user
		userId := id.GenerateId()
		passwordHash, err := s.Hasher.Hash(input.Password)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to create user")
		}
		user := idatabase.User{
			Id:           userId,
			Email:        otpPayload.Email,
			Name:         input.Name,
			PasswordHash: passwordHash,
		}
		if err := s.Database.User.Create(user, ctx); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to create user")
		}

		// Issue login token
		loginTokenPayload := token.LoginTokenPayload{
			UserId: user.Id,
		}
		loginTokenPayloadJson, err := json.Marshal(loginTokenPayload)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to create user")
		}
		loginToken, err := s.Signer.Sign(string(loginTokenPayloadJson), time.Hour*24*7)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to create user")
		}

		// Return
		output := i.CreateUserOutput{
			LoginToken: loginToken,
		}
		return &output, nil
	}
}
