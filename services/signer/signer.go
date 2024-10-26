package signer

import (
	"time"

	"github.com/reizt/rest-go/iservices/isigner"
)

func New() isigner.Service {
	return Service{}
}

type Service struct{}

func (s Service) Sign(json string, expiresIn time.Duration) (string, error) {
	return "", nil
}

func (s Service) Verify(token string) (string, error) {
	return "", nil
}
