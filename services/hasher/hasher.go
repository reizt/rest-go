package hasher

import (
	"fmt"

	"github.com/reizt/rest-go/iservices/ihasher"
	"golang.org/x/crypto/bcrypt"
)

type service struct{}

func New() ihasher.Service {
	return &service{}
}

// Hash は文字列をbcryptでハッシュ化します
func (s *service) Hash(value string) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value cannot be empty")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Validate はハッシュ値と平文を比較検証します
func (s *service) Validate(value, hash string) error {
	if value == "" || hash == "" {
		return fmt.Errorf("value and hash cannot be empty")
	}

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
}
