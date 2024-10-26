package mservices

import (
	"context"

	"github.com/reizt/rest-go/iservices/idatabase"
)

type UserRepo struct {
	GetById_    func(id string, ctx context.Context) (*idatabase.User, error)
	GetByEmail_ func(email string, ctx context.Context) (*idatabase.User, error)
	Create_     func(data idatabase.User, ctx context.Context) error
	Update_     func(data idatabase.User, ctx context.Context) error
	Delete_     func(id string, ctx context.Context) error
}

func (r UserRepo) GetById(id string, ctx context.Context) (*idatabase.User, error) {
	return r.GetById_(id, ctx)
}

func (r UserRepo) GetByEmail(email string, ctx context.Context) (*idatabase.User, error) {
	return r.GetByEmail_(email, ctx)
}

func (r UserRepo) Create(data idatabase.User, ctx context.Context) error {
	return r.Create_(data, ctx)
}

func (r UserRepo) Update(data idatabase.User, ctx context.Context) error {
	return r.Update_(data, ctx)
}

func (r UserRepo) Delete(id string, ctx context.Context) error {
	return r.Delete_(id, ctx)
}
