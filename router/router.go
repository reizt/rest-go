package router

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	h "reij.uno/handlers"
	"reij.uno/iusecases"
)

func New(u *iusecases.All) *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:       true,
		LogURI:          true,
		LogStatus:       true,
		LogResponseSize: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			fmt.Printf("%s %s %d\n", c.Request().Method, c.Request().URL.Path, c.Response().Status)
			return nil
		},
	}))
	e.Use(middleware.Recover())

	e.GET("/hello", h.SayHello(u.SayHello))
	e.POST("/auth/code/issue", h.IssueCode(u.IssueCode))
	e.POST("/auth/code/verify", h.VerifyCode(u.VerifyCode))
	e.POST("/auth/create-user", h.CreateUser(u.CreateUser))
	e.GET("/user", h.GetUser(u.GetUser))
	e.PATCH("/user", h.UpdateUser(u.UpdateUser))

	development := os.Getenv("TEST_CLEAR_DATABASE")
	if development == "on" {
		e.POST("/dev/clear-database", clearDatabase)
	}

	return e
}
