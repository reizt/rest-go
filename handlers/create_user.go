package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/iusecases"
)

type CreateUserReqBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CreateUser(u iusecases.CreateUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var json CreateUserReqBody
		if err := c.Bind(&json); err != nil {
			fmt.Println("json parse error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		otpToken, err := c.Cookie(OTPTokenCookieName)
		if err != nil {
			fmt.Println("cookie error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.CreateUserInput{
			OTPToken: otpToken.Value,
			Name:     json.Name,
			Password: json.Password,
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
