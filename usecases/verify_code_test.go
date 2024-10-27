package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/reizt/rest-go/entities"
	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/iservices/idatabase"
	i "github.com/reizt/rest-go/iusecases"
	"github.com/reizt/rest-go/mservices"
	"github.com/reizt/rest-go/usecases/token"
	"github.com/stretchr/testify/assert"
)

func TestVerifyCode(t *testing.T) {
	t.Parallel()

	sample := struct {
		CodeId           string
		CodeAction       string
		CodeValue        string
		CodeValueHash    string
		UserId           string
		UserEmail        string
		UserName         string
		UserPasswordHash string
		OTPToken         string
	}{
		CodeId:           "_CodeId_",
		CodeAction:       entities.CodeActionCreateUser,
		CodeValue:        "123456",
		CodeValueHash:    "_CodeValueHash_",
		UserId:           "xxx",
		UserEmail:        "john@example.com",
		UserName:         "John",
		UserPasswordHash: "_PasswordHash_",
		OTPToken:         "_LoginToken_",
	}
	input := i.VerifyCodeInput{
		CodeId: sample.CodeId,
		Code:   sample.CodeValue,
	}
	code := idatabase.Code{
		Id:        sample.CodeId,
		Email:     sample.UserEmail,
		Action:    sample.CodeAction,
		ValueHash: sample.CodeValueHash,
		ExpiresAt: int64(time.Now().Add(time.Hour).Unix()),
		CreatedAt: int64(time.Now().Unix()),
	}

	type testCase struct {
		Skip                 bool
		Name                 string
		sDatabaseCodeGetById func(id string, ctx context.Context) (*idatabase.Code, error)
		sHasherValidate      func(value, hash string) error
		sSignerSign          func(json string, expiresIn time.Duration) (string, error)
		expectedErr          error
		assertOutput         func(t *testing.T, output *i.VerifyCodeOutput)
	}

	sDatabaseCodeGetByIdOK := func(id string, ctx context.Context) (*idatabase.Code, error) {
		assert.Equal(t, id, sample.CodeId)
		return &code, nil
	}
	sDatabaseCodeGetByIdNG := func(id string, ctx context.Context) (*idatabase.Code, error) {
		assert.Equal(t, id, sample.CodeId)
		return nil, errors.New("not found")
	}
	sHasherValidateOK := func(value, hash string) error {
		assert.Equal(t, value, sample.CodeValue)
		assert.Equal(t, hash, sample.CodeValueHash)
		return nil
	}
	sHasherValidateNG := func(value, hash string) error {
		assert.Equal(t, value, sample.CodeValue)
		assert.Equal(t, hash, sample.CodeValueHash)
		return errors.New("invalid hash")
	}

	sSignerSignOK := func(payloadJson string, expiresIn time.Duration) (string, error) {
		payload := token.OTPTokenPayload{}
		json.Unmarshal([]byte(payloadJson), &payload)
		assert.Equal(t, payload.Email, sample.UserEmail)
		assert.Equal(t, payload.Action, sample.CodeAction)
		return sample.OTPToken, nil
	}

	sSignerSignNG := func(payloadJson string, expiresIn time.Duration) (string, error) {
		payload := token.OTPTokenPayload{}
		json.Unmarshal([]byte(payloadJson), &payload)
		assert.Equal(t, payload.Email, sample.UserEmail)
		assert.Equal(t, payload.Action, sample.CodeAction)
		return "", errors.New("sign failed")
	}

	testCases := []testCase{
		{
			Name:                 "👎 code not found",
			sDatabaseCodeGetById: sDatabaseCodeGetByIdNG,
			expectedErr:          i.ErrCodeNotFound,
		},
		{
			Name:                 "👎 invalid hash",
			sDatabaseCodeGetById: sDatabaseCodeGetByIdOK,
			sHasherValidate:      sHasherValidateNG,
			expectedErr:          i.ErrInvalidCode,
		},
		{
			Name:                 "👎 sign failed",
			sDatabaseCodeGetById: sDatabaseCodeGetByIdOK,
			sHasherValidate:      sHasherValidateOK,
			sSignerSign:          sSignerSignNG,
			expectedErr:          i.ErrUnexpected,
		},
		{
			Name:                 "👍 works",
			sDatabaseCodeGetById: sDatabaseCodeGetByIdOK,
			sHasherValidate:      sHasherValidateOK,
			sSignerSign:          sSignerSignOK,
			expectedErr:          nil,
			assertOutput: func(t *testing.T, output *i.VerifyCodeOutput) {
				assert.Equal(t, sample.OTPToken, output.Token)
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
					Code: mservices.CodeRepo{
						GetById_: tc.sDatabaseCodeGetById,
					},
				},
				Hasher: mservices.Hasher{
					Validate_: tc.sHasherValidate,
				},
				Signer: mservices.Signer{
					Sign_: tc.sSignerSign,
				},
			}
			u := VerifyCode(&s)

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
