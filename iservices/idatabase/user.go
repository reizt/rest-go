package idatabase

import (
	"context"

	"reij.uno/entities"
)

type User struct {
	Id           string
	Name         string
	Email        string
	PasswordHash string
}

func (u User) ToEntity() entities.User {
	return entities.User{
		Id:    u.Id,
		Email: u.Email,
		Name:  u.Name,
	}
}

type UserUpdate struct {
	Name         string
	Email        string
	PasswordHash string
}

type UserRepo interface {
	GetById(id string, ctx context.Context) (*User, error)
	GetByEmail(email string, ctx context.Context) (*User, error)
	Create(data User, ctx context.Context) error
	Update(data User, ctx context.Context) error
	Delete(id string, ctx context.Context) error
}
