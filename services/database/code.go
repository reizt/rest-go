package database

import (
	"context"

	"github.com/reizt/rest-go/iservices/idatabase"
)

type CodeRepo struct {
}

func (r CodeRepo) GetById(id string, ctx context.Context) (idatabase.Code, error) {
	code := idatabase.Code{
		Id:        "xxx",
		Email:     "john@example.com",
		Action:    "create-user",
		ValueHash: "fdsa",
		ExpiresAt: 0,
		CreatedAt: 0,
	}
	return code, nil
}

func (r CodeRepo) GetByEmail(email string, ctx context.Context) (idatabase.Code, error) {
	code := idatabase.Code{
		Id:        "xxx",
		Email:     "john@example.com",
		Action:    "create-user",
		ValueHash: "fdsa",
		ExpiresAt: 0,
		CreatedAt: 0,
	}
	return code, nil
}

func (r CodeRepo) Create(data idatabase.Code, ctx context.Context) error {
	return nil
}

func (r CodeRepo) Delete(id string, ctx context.Context) error {
	return nil
}
