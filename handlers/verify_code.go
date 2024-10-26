package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/iusecases"
)

type VerifyCodeReqBody struct {
	CodeId string `json:"codeId"`
}

func VerifyCode(u iusecases.VerifyCode) echo.HandlerFunc {
	return func(c echo.Context) error {
		var json VerifyCodeReqBody
		if err := c.Bind(&json); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.VerifyCodeInput{
			CodeId: json.CodeId,
		}
		err := input.Validate()
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		output, err := u(input)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
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
