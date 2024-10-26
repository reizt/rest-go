package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/entities"
	"github.com/reizt/rest-go/iusecases"
)

type GetUserReqBody struct {
	Token string `json:"token"`
}

type GetUserResBody struct {
	User entities.User `json:"user"`
}

func GetUser(u iusecases.GetUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var json GetUserReqBody
		if err := c.Bind(&json); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.GetUserInput{
			LoginToken: json.Token,
		}
		err := input.Validate()
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		output, err := u(input, c.Request().Context())
		if err != nil {
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
