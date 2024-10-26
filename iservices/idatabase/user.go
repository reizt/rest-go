package idatabase

import "github.com/reizt/rest-go/entities"

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
	GetById(id string) (User, error)
	GetByEmail(email string) (User, error)
	Create(data User) error
	Update(id string, data UserUpdate) error
	Delete(id string) error
}
