package handlers

import (
	"fmt"
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
			fmt.Println("json parse error:", err)
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		input := iusecases.IssueCodeInput{
			Email:  json.Email,
			Action: json.Action,
		}
		if err := input.Validate(); err != nil {
			fmt.Println("input validation error:", err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		output, err := u(input, c.Request().Context())
		if err != nil {
			fmt.Println("usecase error:", err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		resBody := IssueCodeResBody{
			CodeId: output.CodeId,
		}
		return c.JSON(http.StatusOK, resBody)
	}
}
