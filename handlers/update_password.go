package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"reij.uno/iusecases"
)

type UpdatePasswordReqBody struct {
	Data UpdatePasswordReqBodyData `json:"data"`
}

type UpdatePasswordReqBodyData struct {
	Password string `json:"password"`
}

func UpdatePassword(u iusecases.UpdatePassword) echo.HandlerFunc {
	return func(c echo.Context) error {
		otpToken, err := c.Cookie(OTPTokenCookieName)
		if err != nil {
			fmt.Println("cookie error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		var reqBody UpdatePasswordReqBody
		if err := c.Bind(&reqBody); err != nil {
			fmt.Println("json parse error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.UpdatePasswordInput{
			OTPToken: otpToken.Value,
			Password: reqBody.Data.Password,
		}
		if err := input.Validate(); err != nil {
			fmt.Println("input validation error:", err)
			switch err {
			case iusecases.ErrInvalidToken:
				return c.String(http.StatusUnauthorized, err.Error())
			default:
				return c.String(http.StatusBadRequest, err.Error())
			}
		}

		if _, err := u(input, c.Request().Context()); err != nil {
			fmt.Println("usecase error:", err)
			switch err {
			case iusecases.ErrInvalidToken:
				return c.String(http.StatusUnauthorized, err.Error())
			default:
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}

		return c.NoContent(http.StatusNoContent)
	}
}
