package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reij.uno/entities"
	"reij.uno/iservices/idatabase"
)

var (
	sampleCode = idatabase.Code{
		Id:        "xxx",
		Email:     "john@example.com",
		Action:    entities.CodeActionCreateUser,
		ValueHash: "fdsa",
		ExpiresAt: 0,
		CreatedAt: 0,
	}
)

func TestCodeGetById(t *testing.T) {
	cleanup()
	defer cleanup()
	// Arrange
	repo := CodeRepo{client}
	ctx := context.Background()
	repo.Create(sampleCode, ctx)

	// Act
	actual, err := repo.GetById(sampleCode.Id, ctx)

	// Assert
	require.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.True(t, reflect.DeepEqual(sampleCode, *actual))
}

func TestCodeGetByEmail(t *testing.T) {
	cleanup()
	defer cleanup()
	// Arrange
	repo := CodeRepo{client}
	ctx := context.Background()
	repo.Create(sampleCode, ctx)

	// Act
	actual, err := repo.GetByEmail(sampleCode.Email, ctx)

	// Assert
	require.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.True(t, reflect.DeepEqual(sampleCode, *actual))
}

func TestCodeCreate(t *testing.T) {
	cleanup()
	defer cleanup()
	// Arrange
	repo := CodeRepo{client}
	ctx := context.Background()

	// Act
	err := repo.Create(sampleCode, ctx)

	// Assert
	require.NoError(t, err)
}

func TestCodeDelete(t *testing.T) {
	cleanup()
	defer cleanup()
	// Arrange
	repo := CodeRepo{client}
	ctx := context.Background()
	repo.Create(sampleCode, ctx)

	// Act
	err := repo.Delete(sampleCode.Id, ctx)

	// Assert
	require.NoError(t, err)
	_, err = repo.GetById(sampleCode.Id, ctx)
	assert.Error(t, err)
}
