package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

func CreateRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "recipes")
	}
}
