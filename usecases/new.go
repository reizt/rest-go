package usecases

import (
	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/iusecases"
)

func New(s *iservices.All) *iusecases.All {
	return &iusecases.All{
		SayHello:   SayHello(s),
		IssueCode:  IssueCode(s),
		VerifyCode: VerifyCode(s),
		CreateUser: CreateUser(s),
		GetUser:    GetUser(s),
		UpdateUser: UpdateUser(s),
	}
}
