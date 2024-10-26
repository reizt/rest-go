package database

import (
	"context"

	"github.com/reizt/rest-go/iservices/idatabase"
)

type UserRepo struct {
}

func (r UserRepo) GetById(id string, ctx context.Context) (*idatabase.User, error) {
	user := idatabase.User{
		Id:           "xxx",
		Name:         "John",
		Email:        "john@example.com",
		PasswordHash: "fdsa",
	}
	return &user, nil
}

func (r UserRepo) GetByEmail(email string, ctx context.Context) (*idatabase.User, error) {
	user := idatabase.User{
		Id:           "xxx",
		Name:         "John",
		Email:        "john@example.com",
		PasswordHash: "fdsa",
	}
	return &user, nil
}

func (r UserRepo) Create(data idatabase.User, ctx context.Context) error {
	return nil
}

func (r UserRepo) Update(data idatabase.User, ctx context.Context) error {
	return nil
}

func (r UserRepo) Delete(id string, ctx context.Context) error {
	return nil
}
