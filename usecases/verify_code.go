package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/reizt/rest-go/iservices"
	i "github.com/reizt/rest-go/iusecases"
	"github.com/reizt/rest-go/usecases/token"
)

func VerifyCode(s *iservices.All) i.VerifyCode {
	return func(input i.VerifyCodeInput, ctx context.Context) (*i.VerifyCodeOutput, error) {
		// Get code from database
		code, err := s.Database.Code.GetById(input.CodeId, ctx)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrCodeNotFound
		}

		// Check if code is expired
		if code.ExpiresAt < time.Now().Unix() {
			return nil, i.ErrCodeExpired
		}

		// Check if code is valid
		if err := s.Hasher.Validate(input.Code, code.ValueHash); err != nil {
			return nil, i.ErrInvalidCode
		}

		// Issue token
		tokenPayload := token.OTPTokenPayload{
			Email:  code.Email,
			Action: code.Action,
		}
		tokenPayloadJson, err := json.Marshal(tokenPayload)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}
		token, err := s.Signer.Sign(string(tokenPayloadJson), time.Hour)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}

		// Return
		output := i.VerifyCodeOutput{
			Token: token,
		}
		return &output, nil
	}
}
