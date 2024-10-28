package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"reij.uno/iusecases"
)

type VerifyCodeReqBody struct {
	CodeId string `json:"codeId"`
	Code   string `json:"code"`
}

func VerifyCode(u iusecases.VerifyCode) echo.HandlerFunc {
	return func(c echo.Context) error {
		var json VerifyCodeReqBody
		if err := c.Bind(&json); err != nil {
			fmt.Println("json parse error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.VerifyCodeInput{
			CodeId: json.CodeId,
			Code:   json.Code,
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
			case iusecases.ErrInvalidCode:
				return c.String(http.StatusUnauthorized, err.Error())
			case iusecases.ErrCodeNotFound:
				return c.String(http.StatusNotFound, err.Error())
			default:
				return c.String(http.StatusInternalServerError, err.Error())
			}
		}

		cookie := http.Cookie{
			Name:     OTPTokenCookieName,
			Value:    output.Token,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		c.SetCookie(&cookie)
		return c.String(http.StatusOK, "OK")
	}
}
