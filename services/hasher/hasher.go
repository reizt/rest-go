package hasher

import (
	"crypto/rand"

	"github.com/reizt/rest-go/iservices/ihasher"
	"golang.org/x/crypto/bcrypt"
)

type service struct{}

func New() ihasher.Service {
	return &service{}
}

func (s *service) Hash(password string) (string, error) {
	salt := make([]byte, bcrypt.MinCost)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword(append([]byte(password), salt...), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(append(hash, salt...)), nil
}

func (s *service) Validate(password, hash string) error {
	hashBytes := []byte(hash)
	salt := hashBytes[bcrypt.MinCost:]
	hashBytes = hashBytes[:bcrypt.MinCost]

	return bcrypt.CompareHashAndPassword(hashBytes, append([]byte(password), salt...))
}
