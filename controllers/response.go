package controllers

import (
	"database/sql"

	"github.com/labstack/echo/v5"
)

func handleError(c *echo.Context, err error) error {
	if err == sql.ErrNoRows {
		return c.JSON(404, map[string]string{
			"message": "Record not found",
		})
	}

	return c.JSON(400, map[string]string{
		"message": err.Error(),
	})
}
