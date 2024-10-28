package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	i "reij.uno/iusecases"
	"reij.uno/mservices"
)

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	sample := struct {
		UserId           string
		UserEmail        string
		UserName         string
		UserPasswordHash string
		LoginToken       string
	}{
		UserId:           "xxx",
		UserEmail:        "john@example.com",
		UserName:         "John",
		UserPasswordHash: "_PasswordHash_",
		LoginToken:       "_LoginToken_",
	}
	input := i.UpdateUserInput{
		LoginToken: sample.LoginToken,
		Data: i.UpdateUserInputData{
			Name: "Jane",
		},
	}
	user := idatabase.User{
		Id:           sample.UserId,
		Email:        sample.UserEmail,
		Name:         sample.UserName,
		PasswordHash: sample.UserPasswordHash,
	}

	auth := func(tk string, ctx context.Context) (*idatabase.User, error) {
		assert.Equal(t, tk, input.LoginToken)
		return &user, nil
	}

	type testCase struct {
		Skip                bool
		Name                string
		sDatabaseUserUpdate func(data idatabase.User, ctx context.Context) error
		expectedErr         error
	}

	sDatabaseUserUpdateOK := func(data idatabase.User, ctx context.Context) error {
		assert.Equal(t, data.Id, user.Id)
		assert.Equal(t, data.Email, user.Email)
		assert.Equal(t, data.Name, input.Data.Name)
		assert.Equal(t, data.PasswordHash, user.PasswordHash)
		return nil
	}
	sDatabaseUserUpdateNG := func(data idatabase.User, ctx context.Context) error {
		return errors.New("failed to update user")
	}

	testCases := []testCase{
		{
			Name:                "👎 update failed",
			sDatabaseUserUpdate: sDatabaseUserUpdateNG,
			expectedErr:         i.ErrUnexpected,
		},
		{
			Name:                "👍 works",
			sDatabaseUserUpdate: sDatabaseUserUpdateOK,
			expectedErr:         nil,
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
						Update_: tc.sDatabaseUserUpdate,
					},
				},
			}
			u := updateUser(&s, auth)

			// Act
			ctx := context.Background()
			_, err := u(input, ctx)

			// Assert
			if tc.expectedErr != nil {
				assert.Error(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
