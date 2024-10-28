package usecases

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"reij.uno/entities"
	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	"reij.uno/iservices/imailer"
	i "reij.uno/iusecases"
	"reij.uno/mservices"
)

func TestIssueCode(t *testing.T) {
	t.Parallel()

	sample := struct {
		UserId           string
		UserEmail        string
		UserName         string
		UserPasswordHash string
		CodeId           string
		CodeValue        string
		CodeValueHash    string
		LoginToken       string
	}{
		UserId:           "xxx",
		UserEmail:        "john@example.com",
		UserName:         "John",
		UserPasswordHash: "_PasswordHash_",
		CodeId:           "_CodeId_",
		CodeValue:        "123456",
		CodeValueHash:    "_CodeValueHash_",
		LoginToken:       "_LoginToken_",
	}
	input := i.IssueCodeInput{
		Email:  sample.UserEmail,
		Action: entities.CodeActionCreateUser,
	}
	user := idatabase.User{
		Id:           sample.UserId,
		Email:        sample.UserEmail,
		Name:         sample.UserName,
		PasswordHash: sample.UserPasswordHash,
	}
	code := idatabase.Code{
		Id:        sample.CodeId,
		Email:     sample.UserEmail,
		Action:    entities.CodeActionCreateUser,
		ValueHash: sample.CodeValueHash,
		ExpiresAt: int64(time.Now().Add(time.Hour).Unix()),
		CreatedAt: int64(time.Now().Unix()),
	}

	os.Setenv("TEST_GENERATE_ID_FIXED_VALUE", sample.CodeId)
	os.Setenv("TEST_GENERATE_CODE_FIXED_VALUE", sample.CodeValue)

	type testCase struct {
		Skip                    bool
		Name                    string
		sDatabaseUserGetByEmail func(id string, ctx context.Context) (*idatabase.User, error)
		sHasherHash             func(value string) (string, error)
		sDatabaseCodeCreate     func(data idatabase.Code, ctx context.Context) error
		sMailerCode             func(input imailer.CodeInput) error
		expectedErr             error
		assertOutput            func(t *testing.T, output *i.IssueCodeOutput)
	}

	sDatabaseUserGetByEmailOK := func(email string, ctx context.Context) (*idatabase.User, error) {
		assert.Equal(t, email, sample.UserEmail)
		return &user, nil
	}
	sDatabaseUserGetByEmailNG := func(email string, ctx context.Context) (*idatabase.User, error) {
		assert.Equal(t, email, sample.UserEmail)
		return nil, errors.New("not found")
	}
	sHasherHashOK := func(value string) (string, error) {
		assert.Equal(t, value, sample.CodeValue)
		return sample.CodeValueHash, nil
	}
	sHasherHashNG := func(value string) (string, error) {
		assert.Equal(t, value, sample.CodeValue)
		return "", errors.New("hash failed")
	}
	sDatabaseCodeCreateOK := func(data idatabase.Code, ctx context.Context) error {
		assert.Equal(t, data.Id, code.Id)
		assert.Equal(t, data.Email, code.Email)
		assert.Equal(t, data.Action, code.Action)
		assert.Equal(t, data.ValueHash, code.ValueHash)
		assert.Greater(t, data.ExpiresAt, int64(0))
		assert.Greater(t, data.CreatedAt, int64(0))
		return nil
	}
	sDatabaseCodeCreateNG := func(data idatabase.Code, ctx context.Context) error {
		return errors.New("failed to create code")
	}
	sMailerCodeOK := func(input imailer.CodeInput) error {
		assert.Equal(t, input.To, sample.UserEmail)
		assert.Equal(t, input.Code, sample.CodeValue)
		return nil
	}
	sMailerCodeNG := func(input imailer.CodeInput) error {
		return errors.New("failed to send email")
	}

	testCases := []testCase{
		{
			Name:                    "👎 user already exists",
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailOK,
			expectedErr:             i.ErrUserAlreadyExists,
		},
		{
			Name:                    "👎 hash failed",
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailNG,
			sHasherHash:             sHasherHashNG,
			expectedErr:             i.ErrUnexpected,
		},
		{
			Name:                    "👎 create code failed",
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailNG,
			sHasherHash:             sHasherHashOK,
			sDatabaseCodeCreate:     sDatabaseCodeCreateNG,
			expectedErr:             i.ErrUnexpected,
		},
		{
			Name:                    "👎 send email failed",
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailNG,
			sHasherHash:             sHasherHashOK,
			sDatabaseCodeCreate:     sDatabaseCodeCreateOK,
			sMailerCode:             sMailerCodeNG,
			expectedErr:             i.ErrUnexpected,
		},
		{
			Name:                    "👍 works",
			sDatabaseUserGetByEmail: sDatabaseUserGetByEmailNG,
			sHasherHash:             sHasherHashOK,
			sDatabaseCodeCreate:     sDatabaseCodeCreateOK,
			sMailerCode:             sMailerCodeOK,
			expectedErr:             nil,
			assertOutput: func(t *testing.T, output *i.IssueCodeOutput) {
				assert.Equal(t, code.Id, output.CodeId)
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
					},
					Code: mservices.CodeRepo{
						Create_: tc.sDatabaseCodeCreate,
					},
				},
				Hasher: mservices.Hasher{
					Hash_: tc.sHasherHash,
				},
				Mailer: mservices.Mailer{
					Code_: tc.sMailerCode,
				},
			}
			u := IssueCode(&s)

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
