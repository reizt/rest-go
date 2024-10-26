package router

import (
	"github.com/labstack/echo/v4"
	"github.com/reizt/rest-go/services/database"
)

func ClearDatabase(c echo.Context) error {
	database.Clean()
	return c.NoContent(204)
}
