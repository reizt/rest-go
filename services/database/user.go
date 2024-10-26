package database

import (
	"context"

	"github.com/reizt/rest-go/ent"
	"github.com/reizt/rest-go/ent/user"
	"github.com/reizt/rest-go/iservices/idatabase"
)

type UserRepo struct {
	client *ent.Client
}

func (UserRepo) toInterface(user *ent.User) *idatabase.User {
	return &idatabase.User{
		Id:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func (r UserRepo) GetById(id string, ctx context.Context) (*idatabase.User, error) {
	user, err := r.client.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return r.toInterface(user), nil
}

func (r UserRepo) GetByEmail(email string, ctx context.Context) (*idatabase.User, error) {
	user, err := r.client.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return r.toInterface(user), nil
}

func (r UserRepo) Create(data idatabase.User, ctx context.Context) error {
	_, err := r.client.User.Create().
		SetID(data.Id).
		SetName(data.Name).
		SetEmail(data.Email).
		SetPasswordHash(data.PasswordHash).
		Save(ctx)
	return err
}

func (r UserRepo) Update(data idatabase.User, ctx context.Context) error {
	_, err := r.client.User.UpdateOneID(data.Id).
		SetName(data.Name).
		SetEmail(data.Email).
		SetPasswordHash(data.PasswordHash).
		Save(ctx)
	return err
}

func (r UserRepo) Delete(id string, ctx context.Context) error {
	err := r.client.User.DeleteOneID(id).Exec(ctx)
	return err
}
