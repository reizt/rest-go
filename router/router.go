package router

import (
	"github.com/labstack/echo/v4"
	h "github.com/reizt/rest-go/handlers"
	"github.com/reizt/rest-go/iusecases"
)

func New(u *iusecases.All) *echo.Echo {
	e := echo.New()

	e.GET("/hello", h.SayHello(u.SayHello))
	e.POST("/auth/code/issue", h.IssueCode(u.IssueCode))
	e.POST("/auth/code/verify", h.VerifyCode(u.VerifyCode))
	e.POST("/auth/create-user", h.CreateUser(u.CreateUser))
	e.GET("/user", h.GetUser(u.GetUser))
	e.PATCH("/user", h.UpdateUser(u.UpdateUser))

	return e
}
