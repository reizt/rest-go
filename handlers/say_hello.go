package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/iusecases"
)

func SayHello(u iusecases.SayHello) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")

		input := iusecases.SayHelloInput{
			Name: name,
		}
		err := input.Validate()
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		output, err := u(input, c.Request().Context())
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, output.Message)
	}
}
