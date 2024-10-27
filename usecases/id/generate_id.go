package id

import (
	"os"

	"github.com/google/uuid"
)

func GenerateId() string {
	fixedValue := os.Getenv("TEST_GENERATE_ID_FIXED_VALUE")
	if fixedValue != "" {
		return fixedValue
	}
	return uuid.New().String()
}
