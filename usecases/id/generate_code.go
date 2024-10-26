package id

import (
	"crypto/rand"
	"os"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const length = 6

func GenerateCode() (string, error) {
	fixedValue := os.Getenv("TEST_GENERATE_CODE_FIXED_VALUE")
	if fixedValue != "" {
		return fixedValue, nil
	}
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b), nil
}
