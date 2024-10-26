package usecases

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/iservices/idatabase"
	i "github.com/reizt/rest-go/iusecases"
	"github.com/reizt/rest-go/usecases/token"
)

func authenticator(s *iservices.All) func(tk string, ctx context.Context) (*idatabase.User, error) {
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
