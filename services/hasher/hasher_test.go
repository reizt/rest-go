package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	s := New()

	password := "password"
	hash1, err := s.Hash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash1)
	hash2, err := s.Hash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash2)
	assert.NotEqual(t, hash1, hash2) // To be sure that hash is different
}

func TestValidate(t *testing.T) {
	s := New()

	password := "password"
	hash, err := s.Hash(password)
	assert.NoError(t, err)
	err = s.Validate(password, hash)
	assert.NoError(t, err)
}
