package handler

import (
	"database/sql"
	"net/http"
	"recipe-api/adapter"
	"recipe-api/database"
	"recipe-api/gateway"
	"recipe-api/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func createRecipeAdapter(c echo.Context, db *sql.DB) adapter.RecipeAdapter {
	recipeDbGateway := gateway.NewRecipeDbGateway(c, db)
	return adapter.NewRecipeAdapter(recipeDbGateway, c)
}

func GetAllRecipes(c echo.Context) error {
	recipeAdapter := createRecipeAdapter(c, database.Db)
	recipes, err := recipeAdapter.GetAllRecipes()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, recipes)
}

func GetRecipe(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	recipeAdapter := createRecipeAdapter(c, database.Db)
	recipe, getRecipeErr := recipeAdapter.GetRecipe(id)
	if getRecipeErr != nil {
		return c.JSON(http.StatusInternalServerError, getRecipeErr)
	}
	return c.JSON(http.StatusOK, recipe)
}

func CreateRecipe(c echo.Context) error {
	var recipe model.Recipe

	// map incoming JSON body to the new recipe
	err := c.Bind(&recipe)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	recipeAdapter := createRecipeAdapter(c, database.Db)
	id, createRecipeErr := recipeAdapter.CreateRecipe(recipe)
	if createRecipeErr != nil {
		return c.JSON(http.StatusInternalServerError, createRecipeErr)
	}
	return c.JSON(http.StatusOK, id)
}

func GetTypes(c echo.Context) error {
	recipeAdapter := createRecipeAdapter(c, database.Db)
	types, err := recipeAdapter.GetTypes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, types)
}

func GetMethods(c echo.Context) error {
	recipeAdapter := createRecipeAdapter(c, database.Db)
	methods, err := recipeAdapter.GetMethods()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, methods)
}
