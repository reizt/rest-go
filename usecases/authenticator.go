package usecases

import (
	"context"
	"encoding/json"
	"fmt"

	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	i "reij.uno/iusecases"
	"reij.uno/usecases/token"
)

type authenticator func(tk string, ctx context.Context) (*idatabase.User, error)

func createAuthenticator(s *iservices.All) authenticator {
	return func(tk string, ctx context.Context) (*idatabase.User, error) {
		// Verify token
		payload, err := s.Signer.Verify(tk)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrInvalidToken
		}
		var loginTokenPayload token.LoginTokenPayload
		if err := json.Unmarshal([]byte(payload), &loginTokenPayload); err != nil {
			fmt.Println(err)
			return nil, i.ErrInvalidToken
		}

		// Get user
		user, err := s.Database.User.GetById(loginTokenPayload.UserId, ctx)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrInvalidToken // to avoid leaking information
		}

		return user, nil
	}
}
