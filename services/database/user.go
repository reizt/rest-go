package database

import (
	"github.com/reizt/rest-go/iservices/idatabase"
)

type UserRepo struct {
}

func (r UserRepo) GetById(id string) (idatabase.User, error) {
	user := idatabase.User{
		Id:           "xxx",
		Name:         "John",
		Email:        "john@example.com",
		PasswordHash: "fdsa",
	}
	return user, nil
}

func (r UserRepo) GetByEmail(email string) (idatabase.User, error) {
	user := idatabase.User{
		Id:           "xxx",
		Name:         "John",
		Email:        "john@example.com",
		PasswordHash: "fdsa",
	}
	return user, nil
}

func (r UserRepo) Create(data idatabase.User) error {
	return nil
}

func (r UserRepo) Update(id string, data idatabase.UserUpdate) error {
	return nil
}

func (r UserRepo) Delete(id string) error {
	return nil
}
