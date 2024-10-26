package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/iservices/idatabase"
	"github.com/reizt/rest-go/iservices/imailer"
	i "github.com/reizt/rest-go/iusecases"
	"github.com/reizt/rest-go/mservices"
	"github.com/reizt/rest-go/usecases/token"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Skip                bool
		Name                string
		sSignerSign         func(json string, expiresIn time.Duration) (string, error)
		sSignerVerify       func(token string) (string, error)
		sHasherHash         func(value string) (string, error)
		sMailerCode         func(input imailer.CodeInput) error
		sDatabaseUserCreate func(data idatabase.User, ctx context.Context) error
		expectedErr         error
		assertOutput        func(t *testing.T, output *i.CreateUserOutput)
	}

	sample := struct {
		Email        string
		PasswordHash string
		LoginToken   string
	}{
		Email:        "john@example.com",
		PasswordHash: "_PasswordHash_",
		LoginToken:   "_LoginToken_",
	}
	input := i.CreateUserInput{
		OTPToken: "_OTPToken_",
		Name:     "_Name_",
		Password: "_Password_",
	}

	sDatabaseUserCreateOK := func(data idatabase.User, ctx context.Context) error {
		assert.NotEmpty(t, data.Id)
		assert.Equal(t, data.Email, sample.Email)
		assert.Equal(t, data.Name, input.Name)
		assert.Equal(t, data.PasswordHash, sample.PasswordHash)
		return nil
	}
	sDatabaseUserCreateNG := func(data idatabase.User, ctx context.Context) error {
		return errors.New("failed to create user")
	}
	sSignerSignOK := func(payloadJson string, expiresIn time.Duration) (string, error) {
		var payload token.OTPTokenPayload
		err := json.Unmarshal([]byte(payloadJson), &payload)
		assert.NoError(t, err)
		assert.Equal(t, expiresIn, LoginTokenExpiration)
		return sample.LoginToken, nil
	}
	sSignerSignNG := func(json string, expiresIn time.Duration) (string, error) {
		return "", errors.New("sign failed")
	}
	sSignerVerifyOK := func(tk string) (string, error) {
		assert.Equal(t, tk, input.OTPToken)
		payload := token.OTPTokenPayload{Email: sample.Email}
		payloadJson, _ := json.Marshal(payload)
		return string(payloadJson), nil
	}
	sSignerVerifyNG := func(token string) (string, error) {
		assert.Equal(t, token, input.OTPToken)
		return "", errors.New("invalid token")
	}
	sHasherHashOK := func(value string) (string, error) {
		assert.Equal(t, value, input.Password)
		return sample.PasswordHash, nil
	}
	sHasherHashNG := func(value string) (string, error) {
		return "", errors.New("hash failed")
	}

	testCases := []TestCase{
		{
			Name:          "👎 invalid token",
			sSignerVerify: sSignerVerifyNG,
			expectedErr:   i.ErrInvalidToken,
		},
		{
			Name:          "👎 hash failed",
			sSignerVerify: sSignerVerifyOK,
			sHasherHash:   sHasherHashNG,
			expectedErr:   i.ErrUnexpected,
		},
		{
			Name:                "👎 create user failed",
			sSignerVerify:       sSignerVerifyOK,
			sHasherHash:         sHasherHashOK,
			sDatabaseUserCreate: sDatabaseUserCreateNG,
			expectedErr:         i.ErrUnexpected,
		},
		{
			Name:                "👎 sign failed",
			sSignerSign:         sSignerSignNG,
			sSignerVerify:       sSignerVerifyOK,
			sHasherHash:         sHasherHashOK,
			sDatabaseUserCreate: sDatabaseUserCreateNG,
			expectedErr:         i.ErrUnexpected,
		},
		{
			Name:                "👍 works",
			sSignerSign:         sSignerSignOK,
			sSignerVerify:       sSignerVerifyOK,
			sHasherHash:         sHasherHashOK,
			sDatabaseUserCreate: sDatabaseUserCreateOK,
			expectedErr:         nil,
			assertOutput: func(t *testing.T, output *i.CreateUserOutput) {
				assert.Equal(t, sample.LoginToken, output.LoginToken)
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
						Create_: tc.sDatabaseUserCreate,
					},
					Code: mservices.CodeRepo{},
				},
				Hasher: mservices.Hasher{
					Hash_: tc.sHasherHash,
				},
				Mailer: mservices.Mailer{
					Code_: tc.sMailerCode,
				},
				Signer: mservices.Signer{
					Sign_:   tc.sSignerSign,
					Verify_: tc.sSignerVerify,
				},
			}
			u := CreateUser(&s)

			// Act
			ctx := context.Background()
			output, err := u(input, ctx)

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
