package hasher

import "github.com/reizt/rest-go/iservices/ihasher"

func New() ihasher.Service {
	return Service{}
}

type Service struct{}

func (h Service) Hash(password string) (string, error) {
	return "", nil
}

func (h Service) Validate(password, hash string) error {
	return nil
}
