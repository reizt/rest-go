package usecases

import (
	"reij.uno/iservices"
	"reij.uno/iusecases"
)

func New(s *iservices.All) *iusecases.All {
	auth := createAuthenticator(s)
	return &iusecases.All{
		SayHello:       sayHello(s),
		IssueCode:      issueCode(s),
		VerifyCode:     verifyCode(s),
		CreateUser:     createUser(s),
		GetUser:        getUser(s, auth),
		UpdateUser:     updateUser(s, auth),
		UpdatePassword: updatePassword(s),
	}
}
