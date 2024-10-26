package signer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	privateKey := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgUfx+5tNZ0dTOaFMX
FTrXoXwukwY2/jiu7V8c1z2YakGhRANCAARB9ynRk3QhL3KQAwxusvBI14+XXmwE
gprEy+PrKc/trp4ig5olG407OdLvaonbXLGjFHsSvPYO7M6HGgoqCFn6
-----END PRIVATE KEY-----`
	publicKey := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEQfcp0ZN0IS9ykAMMbrLwSNePl15s
BIKaxMvj6ynP7a6eIoOaJRuNOznS72qJ21yxoxR7Erz2DuzOhxoKKghZ+g==
-----END PUBLIC KEY-----`
	s, err := New(privateKey, publicKey)
	assert.NoError(t, err)

	// Sign
	json := `{"name":"John"}`
	expiresIn := time.Hour
	token, err := s.Sign(json, expiresIn)
	println(token)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify
	payload, err := s.Verify(token)
	assert.NoError(t, err)
	assert.Equal(t, json, payload)
}
