package router

import (
	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/handlers"
	"github.com/reizt/rest-go/iusecases"
)

func New(u *iusecases.All) *echo.Echo {
	e := echo.New()

	e.GET("/hello", handlers.SayHello(u.SayHello))
	e.POST("/auth/code/issue", handlers.IssueCode(u.IssueCode))
	e.POST("/auth/code/verify", handlers.VerifyCode(u.VerifyCode))
	e.POST("/auth/create-user", handlers.CreateUser(u.CreateUser))
	e.GET("/user", handlers.GetUser(u.GetUser))

	return e
}
