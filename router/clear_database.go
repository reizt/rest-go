package router

import (
	"github.com/labstack/echo/v4"
	"reij.uno/services/database"
)

func clearDatabase(c echo.Context) error {
	database.Clean()
	return c.NoContent(204)
}
