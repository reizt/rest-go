package router

import (
	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/handlers"
	"github.com/reizt/rest-go/iusecases"
)

func New(u *iusecases.All) *echo.Echo {
	e := echo.New()

	e.GET("/hello", handlers.SayHello(u.SayHello))

	return e
}
