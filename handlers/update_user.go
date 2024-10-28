package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"reij.uno/iusecases"
)

type UpdateUserReqBodyData struct {
	Name string `json:"name"`
}

type UpdateUserReqBody struct {
	Data UpdateUserReqBodyData `json:"data"`
}

func UpdateUser(u iusecases.UpdateUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody UpdateUserReqBody
		if err := c.Bind(&reqBody); err != nil {
			fmt.Println("json parse error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		loginToken, err := c.Cookie(LoginTokenCookieName)
		if err != nil {
			fmt.Println("cookie error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.UpdateUserInput{
			LoginToken: loginToken.Value,
			Data: iusecases.UpdateUserInputData{
				Name: reqBody.Data.Name,
			},
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
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
