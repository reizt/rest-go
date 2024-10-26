package database

import (
	"context"

	"github.com/reizt/rest-go/ent"
	"github.com/reizt/rest-go/ent/code"
	"github.com/reizt/rest-go/iservices/idatabase"
)

type CodeRepo struct {
	client *ent.Client
}

func (CodeRepo) toInterface(code *ent.Code) *idatabase.Code {
	return &idatabase.Code{
		Id:        code.ID,
		Email:     code.Email,
		Action:    code.Action,
		ValueHash: code.ValueHash,
		ExpiresAt: int64(code.ExpiresAt),
		CreatedAt: int64(code.CreatedAt),
	}
}

func (r CodeRepo) GetById(id string, ctx context.Context) (*idatabase.Code, error) {
	code, err := r.client.Code.Query().Where(code.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return r.toInterface(code), nil
}

func (r CodeRepo) GetByEmail(email string, ctx context.Context) (*idatabase.Code, error) {
	code, err := r.client.Code.Query().Where(code.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return r.toInterface(code), nil
}

func (r CodeRepo) Create(data idatabase.Code, ctx context.Context) error {
	_, err := r.client.Code.Create().
		SetID(data.Id).
		SetEmail(data.Email).
		SetAction(data.Action).
		SetValueHash(data.ValueHash).
		SetExpiresAt(int(data.ExpiresAt)).
		SetCreatedAt(int(data.CreatedAt)).
		Save(ctx)
	return err
}

func (r CodeRepo) Delete(id string, ctx context.Context) error {
	err := r.client.Code.DeleteOneID(id).Exec(ctx)
	return err
}

func (r CodeRepo) deleteAll(ctx context.Context) error {
	_, err := r.client.Code.Delete().Exec(ctx)
	return err
}
