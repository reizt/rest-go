package database

import "github.com/reizt/rest-go/iservices/idatabase"

func New() idatabase.Service {
	return idatabase.Service{
		User: UserRepo{},
		Code: CodeRepo{},
	}
}
