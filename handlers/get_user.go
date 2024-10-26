package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/entities"
	"github.com/reizt/rest-go/iusecases"
)

type GetUserResBody struct {
	User entities.User `json:"user"`
}

func GetUser(u iusecases.GetUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		loginToken, err := c.Cookie(LoginTokenCookieName)
		if err != nil {
			fmt.Println("cookie error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.GetUserInput{
			LoginToken: loginToken.Value,
		}
		if err := input.Validate(); err != nil {
			fmt.Println("input validation error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		output, err := u(input, c.Request().Context())
		if err != nil {
			fmt.Println("usecase error:", err)
			switch err {
			case iusecases.ErrInvalidToken:
				return c.String(http.StatusUnauthorized, err.Error())
			default:
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}

		resBody := GetUserResBody{
			User: output.User,
		}
		return c.JSON(http.StatusOK, resBody)
	}
}
