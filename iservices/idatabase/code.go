package idatabase

import "context"

type Code struct {
	Id        string
	Email     string
	Action    string
	ValueHash string
	ExpiresAt int64
	CreatedAt int64
}

type CodeRepo interface {
	GetById(id string, ctx context.Context) (*Code, error)
	GetByEmail(email string, ctx context.Context) (*Code, error)
	Create(data Code, ctx context.Context) error
	Delete(id string, ctx context.Context) error
}
