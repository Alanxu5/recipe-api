package handlers

import (
	"database/sql"
	"net/http"
	"recipe/api/models"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetRecipes(db))
	}
}

func CreateRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// init a new recipe
		var recipe models.Recipe

		// map incoming JSON body to the new recipe
		c.Bind(&recipe)

		// add a recipe using our model
		id, err := models.CreateRecipe(db, recipe.Name)

		// if creation is successful return a response
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

func DeleteRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		// delete a recipe using our model
		deletedId, err := models.DeleteRecipe(db, id)

		// if deletion is sucessful return a response
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": deletedId,
			})
		} else {
			return err
		}
	}
}
