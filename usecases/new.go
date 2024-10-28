package usecases

import (
	"reij.uno/iservices"
	"reij.uno/iusecases"
)

func New(s *iservices.All) *iusecases.All {
	auth := createAuthenticator(s)
	return &iusecases.All{
		SayHello:   SayHello(s),
		IssueCode:  IssueCode(s),
		VerifyCode: VerifyCode(s),
		CreateUser: CreateUser(s),
		GetUser:    GetUser(s, auth),
		UpdateUser: UpdateUser(s, auth),
	}
}
