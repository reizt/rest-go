package usecases

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	i "reij.uno/iusecases"
)

func TestGetUser(t *testing.T) {
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
	input := i.GetUserInput{
		LoginToken: sample.LoginToken,
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

	s := iservices.All{}

	u := getUser(&s, auth)
	output, err := u(input, context.Background())
	require.NoError(t, err)
	assert.Equal(t, user.Id, output.User.Id)
	assert.Equal(t, user.Email, output.User.Email)
	assert.Equal(t, user.Name, output.User.Name)
}
