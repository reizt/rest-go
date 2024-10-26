package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/iusecases"
)

type UpdateUserReqBody struct {
	Token string `json:"token"`
	Data  struct {
		Name string `json:"name"`
	} `json:"data"`
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
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		if _, err := u(input); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
