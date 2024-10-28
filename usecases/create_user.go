package usecases

import (
	"context"
	"encoding/json"
	"fmt"

	"reij.uno/entities"
	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	i "reij.uno/iusecases"
	"reij.uno/usecases/id"
	"reij.uno/usecases/token"
)

func CreateUser(s *iservices.All) i.CreateUser {
	return func(input i.CreateUserInput, ctx context.Context) (*i.CreateUserOutput, error) {
		// Verify token
		payload, err := s.Signer.Verify(input.OTPToken)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrInvalidToken
		}
		var otpPayload token.OTPTokenPayload
		if err := json.Unmarshal([]byte(payload), &otpPayload); err != nil {
			fmt.Println(err)
			return nil, i.ErrInvalidToken
		}
		if otpPayload.Action != entities.CodeActionCreateUser {
			return nil, i.ErrInvalidToken
		}

		// Create user
		userId := id.GenerateId()
		passwordHash, err := s.Hasher.Hash(input.Password)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}
		user := idatabase.User{
			Id:           userId,
			Email:        otpPayload.Email,
			Name:         input.Name,
			PasswordHash: passwordHash,
		}
		if err := s.Database.User.Create(user, ctx); err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}

		// Issue login token
		loginTokenPayload := token.LoginTokenPayload{
			UserId: user.Id,
		}
		loginTokenPayloadJson, err := json.Marshal(loginTokenPayload)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}
		loginToken, err := s.Signer.Sign(string(loginTokenPayloadJson), LoginTokenExpiration)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}

		// Return
		output := i.CreateUserOutput{
			LoginToken: loginToken,
		}
		return &output, nil
	}
}
