package mservices

import "time"

type Signer struct {
	Sign_   func(json string, expiresIn time.Duration) (string, error)
	Verify_ func(token string) (string, error)
}

func (s Signer) Sign(json string, expiresIn time.Duration) (string, error) {
	return s.Sign_(json, expiresIn)
}

func (s Signer) Verify(token string) (string, error) {
	return s.Verify_(token)
}
