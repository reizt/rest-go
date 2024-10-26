package usecases

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/reizt/rest-go/iservices"
	i "github.com/reizt/rest-go/iusecases"
	"github.com/reizt/rest-go/usecases/id"
	"github.com/reizt/rest-go/usecases/token"
)

func VerifyCode(s *iservices.All) i.VerifyCode {
	return func(input i.VerifyCodeInput) (*i.VerifyCodeOutput, error) {
		// Get code from database
		code, err := s.Database.Code.GetById(input.CodeId)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("not found")
		}

		// Check if code is expired
		if code.ExpiresAt < time.Now().Unix() {
			return nil, fmt.Errorf("expired")
		}

		// Check if code is valid
		if code.ValueHash != id.GenerateCode() {
			return nil, fmt.Errorf("invalid code")
		}

		// Issue token
		tokenPayload := token.OTPTokenPayload{
			Email: code.Email,
		}
		tokenPayloadJson, err := json.Marshal(tokenPayload)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to issue token")
		}
		token, err := s.Signer.Sign(string(tokenPayloadJson), time.Hour)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to issue token")
		}

		// Return
		output := i.VerifyCodeOutput{
			Token: token,
		}
		return &output, nil
	}
}
