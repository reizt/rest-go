package mock

import (
	"context"

	"github.com/reizt/rest-go/iservices/idatabase"
)

type CodeRepo struct {
	GetById_    func(id string, ctx context.Context) (*idatabase.Code, error)
	GetByEmail_ func(email string, ctx context.Context) (*idatabase.Code, error)
	Create_     func(data idatabase.Code, ctx context.Context) error
	Delete_     func(id string, ctx context.Context) error
}

func (r CodeRepo) GetById(id string, ctx context.Context) (*idatabase.Code, error) {
	return r.GetById_(id, ctx)
}

func (r CodeRepo) GetByEmail(email string, ctx context.Context) (*idatabase.Code, error) {
	return r.GetByEmail_(email, ctx)
}

func (r CodeRepo) Create(data idatabase.Code, ctx context.Context) error {
	return r.Create_(data, ctx)
}

func (r CodeRepo) Delete(id string, ctx context.Context) error {
	return r.Delete_(id, ctx)
}
