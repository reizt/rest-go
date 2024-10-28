package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reij.uno/iservices/idatabase"
)

var (
	sampleUser = idatabase.User{
		Id:           "xxx",
		Email:        "john@example.com",
		Name:         "John",
		PasswordHash: "foo",
	}
)

func TestUserGetById(t *testing.T) {
	cleanup()
	defer cleanup()
	// Arrange
	repo := UserRepo{client}
	ctx := context.Background()
	repo.Create(sampleUser, ctx)

	// Act
	actual, err := repo.GetById(sampleUser.Id, ctx)

	// Assert
	require.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.True(t, reflect.DeepEqual(sampleUser, *actual))
}

func TestUserGetByEmail(t *testing.T) {
	cleanup()
	defer cleanup()
	// Arrange
	repo := UserRepo{client}
	ctx := context.Background()
	repo.Create(sampleUser, ctx)

	// Act
	actual, err := repo.GetByEmail(sampleUser.Email, ctx)

	// Assert
	require.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.True(t, reflect.DeepEqual(sampleUser, *actual))
}

func TestUserCreate(t *testing.T) {
	cleanup()
	defer cleanup()
	// Arrange
	repo := UserRepo{client}
	ctx := context.Background()

	// Act
	err := repo.Create(sampleUser, ctx)

	// Assert
	require.NoError(t, err)
}

func TestUserUpdate(t *testing.T) {
	cleanup()
	defer cleanup()
	// Arrange
	repo := UserRepo{client}
	ctx := context.Background()
	repo.Create(sampleUser, ctx)

	// Act
	newUser := idatabase.User{
		Id:           sampleUser.Id,
		Email:        "jane@example.com",
		Name:         "Jane",
		PasswordHash: "bar",
	}
	err := repo.Update(newUser, ctx)

	// Assert
	require.NoError(t, err)
	found, err := repo.GetById(sampleUser.Id, ctx)
	require.NoError(t, err)
	assert.True(t, reflect.DeepEqual(newUser, *found))
}

func TestUserDelete(t *testing.T) {
	cleanup()
	defer cleanup()
	// Arrange
	repo := UserRepo{client}
	ctx := context.Background()
	repo.Create(sampleUser, ctx)

	// Act
	err := repo.Delete(sampleUser.Id, ctx)

	// Assert
	require.NoError(t, err)
	_, err = repo.GetById(sampleUser.Id, ctx)
	assert.Error(t, err)
}
