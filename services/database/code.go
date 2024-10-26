package database

import (
	"github.com/reizt/rest-go/iservices/idatabase"
)

type CodeRepo struct {
}

func (r CodeRepo) GetById(id string) (idatabase.Code, error) {
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

func (r CodeRepo) GetByEmail(email string) (idatabase.Code, error) {
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

func (r CodeRepo) Create(data idatabase.Code) error {
	return nil
}

func (r CodeRepo) Delete(id string) error {
	return nil
}
