package iusecases

import "errors"

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrUnexpected        = errors.New("unexpected error")
	ErrCodeNotFound      = errors.New("code not found")
	ErrCodeExpired       = errors.New("code expired")
	ErrInvalidCode       = errors.New("invalid code")
	ErrUserAlreadyExists = errors.New("user already exists")
)
