package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/iusecases"
)

type CreateUserReqBody struct {
	Token    string `json:"token"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CreateUser(u iusecases.CreateUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var json CreateUserReqBody
		if err := c.Bind(&json); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.CreateUserInput{
			OTPToken: json.Token,
			Name:     json.Name,
			Password: json.Password,
		}
		err := input.Validate()
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		output, err := u(input)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}

		cookie := http.Cookie{
			Name:     LoginTokenCookieName,
			Value:    output.LoginToken,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		c.SetCookie(&cookie)
		return c.NoContent(http.StatusCreated)
	}
}
