package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	i "reij.uno/iusecases"
	"reij.uno/mservices"
	"reij.uno/usecases/token"
)

func TestAuthenticator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		Skip                 bool
		Name                 string
		sSignerVerify        func(token string) (string, error)
		sDatabaseUserGetById func(id string, ctx context.Context) (*idatabase.User, error)
		expectedErr          error
		assertOutput         func(t *testing.T, output *idatabase.User)
	}

	sample := struct {
		User       *idatabase.User
		LoginToken string
	}{
		User: &idatabase.User{
			Id:           "_Id_",
			Email:        "john@example.com",
			Name:         "_Name_",
			PasswordHash: "_PasswordHash_",
		},
		LoginToken: "_LoginToken_",
	}

	sSignerVerifyOK := func(tk string) (string, error) {
		tokenPayload := token.LoginTokenPayload{
			UserId: sample.User.Id,
		}
		tokenPayloadJson, _ := json.Marshal(tokenPayload)
		return string(tokenPayloadJson), nil
	}
	sSignerVerifyNG := func(tk string) (string, error) {
		return "", errors.New("invalid token")
	}
	sDatabaseUserGetByIdOK := func(id string, ctx context.Context) (*idatabase.User, error) {
		assert.Equal(t, id, sample.User.Id)
		return sample.User, nil
	}
	sDatabaseUserGetByIdNG := func(id string, ctx context.Context) (*idatabase.User, error) {
		return nil, errors.New("failed to get user")
	}

	testCases := []testCase{
		{
			Name:          "👎 invalid token",
			sSignerVerify: sSignerVerifyNG,
			expectedErr:   i.ErrInvalidToken,
		},
		{
			Name:                 "👎 user not found",
			sSignerVerify:        sSignerVerifyOK,
			sDatabaseUserGetById: sDatabaseUserGetByIdNG,
			expectedErr:          i.ErrInvalidToken,
		},
		{
			Name:                 "👍 works",
			sSignerVerify:        sSignerVerifyOK,
			sDatabaseUserGetById: sDatabaseUserGetByIdOK,
			expectedErr:          nil,
			assertOutput: func(t *testing.T, output *idatabase.User) {
				assert.Equal(t, sample.User.Id, output.Id)
				assert.Equal(t, sample.User.Email, output.Email)
				assert.Equal(t, sample.User.Name, output.Name)
				assert.Equal(t, sample.User.PasswordHash, output.PasswordHash)
			},
		},
	}

	for _, tc := range testCases {
		if tc.Skip {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			// Arrange
			s := iservices.All{
				Database: idatabase.Service{
					User: mservices.UserRepo{
						GetById_: tc.sDatabaseUserGetById,
					},
				},
				Signer: mservices.Signer{
					Verify_: tc.sSignerVerify,
				},
			}
			u := createAuthenticator(&s)

			// Act
			ctx := context.Background()
			output, err := u(sample.LoginToken, ctx)

			// Assert
			if tc.expectedErr != nil {
				assert.Error(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
				tc.assertOutput(t, output)
			}
		})
	}
}
