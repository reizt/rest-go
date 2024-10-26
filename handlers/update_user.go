package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/iusecases"
)

type UpdateUserReqBodyData struct {
	Name string `json:"name"`
}

type UpdateUserReqBody struct {
	Token string                `json:"token"`
	Data  UpdateUserReqBodyData `json:"data"`
}

func UpdateUser(u iusecases.UpdateUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var json UpdateUserReqBody
		if err := c.Bind(&json); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.UpdateUserInput{
			LoginToken: json.Token,
			Data: iusecases.UpdateUserInputData{
				Name: json.Data.Name,
			},
		}
		err := input.Validate()
		if err != nil {
			switch err {
			case iusecases.ErrInvalidToken:
				return c.String(http.StatusUnauthorized, err.Error())
			default:
				return c.String(http.StatusBadRequest, err.Error())
			}
		}

		if _, err := u(input, c.Request().Context()); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
