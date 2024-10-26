package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/iusecases"
)

type IssueCodeReqBody struct {
	Email  string `json:"email"`
	Action string `json:"action"`
}

type IssueCodeResBody struct {
	CodeId string `json:"code_id"`
}

func IssueCode(u iusecases.IssueCode) echo.HandlerFunc {
	return func(c echo.Context) error {
		var json IssueCodeReqBody
		if err := c.Bind(&json); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.IssueCodeInput{
			Email:  json.Email,
			Action: json.Action,
		}
		err := input.Validate()
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		output, err := u(input)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}

		resBody := IssueCodeResBody{
			CodeId: output.CodeId,
		}
		return c.JSON(http.StatusOK, resBody)
	}
}
