package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reij.uno/entities"
	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	i "reij.uno/iusecases"
	"reij.uno/mservices"
	"reij.uno/usecases/token"
)

func TestUpdatePassword(t *testing.T) {
	t.Parallel()

	type testCase struct {
		Skip                    bool
		Name                    string
		sSignerSign             func(json string, expiresIn time.Duration) (string, error)
		sSignerVerify           func(token string) (string, error)
		sHasherHash             func(value string) (string, error)
		sDatabaseUserGetByEmail func(email string, ctx context.Context) (*idatabase.User, error)
		sDatabaseUserUpdate     func(data idatabase.User, ctx context.Context) error
		expectedErr             error
		assertOutput            func(t *testing.T, output *i.UpdatePasswordOutput)
	}

	sample := struct {
		UserId       string
		Name         string
		Email        string
		PasswordHash string
		OTPToken     string
		LoginToken   string
	}{
		UserId:       "_UserId_",
		Name:         "_Name_",
		Email:        "john@example.com",
		PasswordHash: "_PasswordHash_",
		OTPToken:     "_OTPToken_",
		LoginToken:   "_LoginToken_",
	}
	input := i.UpdatePasswordInput{
		OTPToken: "_OTPToken_",
		Password: "_Password_",
	}

	sDatabaseUserGetByEmailOK := func(email string, ctx context.Context) (*idatabase.User, error) {
		assert.Equal(t, email, sample.Email)
		return &idatabase.User{
			Id:           sample.UserId,
			Email:        sample.Email,
			Name:         sample.Name,
			PasswordHash: sample.PasswordHash,
		}, nil
	}
	sDatabaseUserGetByEmailNG := func(email string, ctx context.Context) (*idatabase.User, error) {
		assert.Equal(t, email, sample.Email)
		return nil, errors.New("not found")
	}
	sDatabaseUserUpdateOK := func(data idatabase.User, ctx context.Context) error {
		assert.NotEmpty(t, data.Id)
		assert.Equal(t, data.Email, sample.Email)
		assert.Equal(t, data.Name, sample.Name)
		assert.Equal(t, data.PasswordHash, sample.PasswordHash)
		return nil
	}
	sDatabaseUserUpdateNG := func(data idatabase.User, ctx context.Context) error {
		return errors.New("failed to create user")
	}
	sSignerSignOK := func(payloadJson string, expiresIn time.Duration) (string, error) {
		var payload token.OTPTokenPayload
		err := json.Unmarshal([]byte(payloadJson), &payload)
		require.NoError(t, err)
		assert.Equal(t, expiresIn, LoginTokenExpiration)
		return sample.LoginToken, nil
	}
	sSignerSignNG := func(json string, expiresIn time.Duration) (string, error) {
		return "", errors.New("sign failed")
	}
	sSignerVerifyOK := func(tk string) (string, error) {
		assert.Equal(t, tk, input.OTPToken)
		payload := token.OTPTokenPayload{
			Email:  sample.Email,
			Action: entities.CodeActionResetPassword,
		}
		payloadJson, _ := json.Marshal(payload)
		return string(payloadJson), nil
	}
	sSignerVerifyDifferentAction := func(tk string) (string, error) {
		assert.Equal(t, tk, input.OTPToken)
		payload := token.OTPTokenPayload{
			Email:  sample.Email,
			Action: entities.CodeActionCreateUser,
		}
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

	testCases := []testCase{
		{
			Name:          "👎 invalid token",
			sSignerVerify: sSignerVerifyNG,
			expectedErr:   i.ErrInvalidToken,
		},
		{
			Name:          "👎 invalid action of token",
			sSignerVerify: sSignerVerifyDifferentAction,
			expectedErr:   i.ErrInvalidToken,
		},
		{
			Name:                    "👎 user not found",
			sSignerVerify:           sSignerVerifyOK,
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailNG,
			expectedErr:             i.ErrInvalidToken,
		},
		{
			Name:                    "👎 hash failed",
			sSignerVerify:           sSignerVerifyOK,
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailOK,
			sHasherHash:             sHasherHashNG,
			expectedErr:             i.ErrUnexpected,
		},
		{
			Name:                    "👎 create user failed",
			sSignerVerify:           sSignerVerifyOK,
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailOK,
			sHasherHash:             sHasherHashOK,
			sDatabaseUserUpdate:     sDatabaseUserUpdateNG,
			expectedErr:             i.ErrUnexpected,
		},
		{
			Name:                    "👎 sign failed",
			sSignerSign:             sSignerSignNG,
			sSignerVerify:           sSignerVerifyOK,
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailOK,
			sHasherHash:             sHasherHashOK,
			sDatabaseUserUpdate:     sDatabaseUserUpdateNG,
			expectedErr:             i.ErrUnexpected,
		},
		{
			Name:                    "👍 works",
			sSignerSign:             sSignerSignOK,
			sSignerVerify:           sSignerVerifyOK,
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailOK,
			sHasherHash:             sHasherHashOK,
			sDatabaseUserUpdate:     sDatabaseUserUpdateOK,
			expectedErr:             nil,
			assertOutput: func(t *testing.T, output *i.UpdatePasswordOutput) {
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
						GetByEmail_: tc.sDatabaseUserGetByEmail,
						Update_:     tc.sDatabaseUserUpdate,
					},
					Code: mservices.CodeRepo{},
				},
				Hasher: mservices.Hasher{
					Hash_: tc.sHasherHash,
				},
				Signer: mservices.Signer{
					Sign_:   tc.sSignerSign,
					Verify_: tc.sSignerVerify,
				},
			}
			u := updatePassword(&s)

			// Act
			ctx := context.Background()
			output, err := u(input, ctx)

			// Assert
			if tc.expectedErr != nil {
				assert.Error(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
				tc.assertOutput(t, output)
			}
		})
	}
}
