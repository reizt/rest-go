package usecases

import (
	"context"
	"encoding/json"
	"fmt"

	"reij.uno/entities"
	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	i "reij.uno/iusecases"
	"reij.uno/usecases/token"
)

func updatePassword(s *iservices.All) i.UpdatePassword {
	return func(input i.UpdatePasswordInput, ctx context.Context) (*i.UpdatePasswordOutput, error) {
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
		if otpPayload.Action != entities.CodeActionResetPassword {
			return nil, i.ErrInvalidToken
		}

		// Get user
		user, err := s.Database.User.GetByEmail(otpPayload.Email, ctx)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrInvalidToken
		}

		// Update user
		passwordHash, err := s.Hasher.Hash(input.Password)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}
		newUser := idatabase.User{
			Id:           user.Id,
			Email:        user.Email,
			Name:         user.Name,
			PasswordHash: passwordHash,
		}
		if err := s.Database.User.Update(newUser, ctx); err != nil {
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

		output := i.UpdatePasswordOutput{
			LoginToken: loginToken,
		}
		return &output, nil
	}
}
