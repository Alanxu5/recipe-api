package handler

import (
	"database/sql"
	"net/http"
	"recipe-api/adapter"
	"recipe-api/database"
	"recipe-api/gateway"
	"recipe-api/model"
	"strconv"

	"github.com/labstack/echo"
)

func createRecipeAdapter(c echo.Context, db *sql.DB) adapter.RecipeAdapter {
	recipeDbGateway := gateway.NewRecipeDbGateway(c, db)
	return adapter.NewRecipeAdapter(recipeDbGateway, c)
}

func GetAllRecipes(c echo.Context) error {
	recipeAdapter := createRecipeAdapter(c, database.Db)
	recipes, err := recipeAdapter.GetAllRecipes()

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, recipes)
}

func GetRecipe(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	recipeAdapter := createRecipeAdapter(c, database.Db)
	recipe, err := recipeAdapter.GetRecipe(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, recipe)
}

func CreateRecipe(c echo.Context) error {
	var recipe model.Recipe

	// map incoming JSON body to the new recipe
	c.Bind(&recipe)
	recipeAdapter := createRecipeAdapter(c, database.Db)
	id, err := recipeAdapter.CreateRecipe(recipe)

	// if creation is successful return a response
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, id)
}

func GetTypes(c echo.Context) error {
	recipeAdapter := createRecipeAdapter(c, database.Db)
	types, err := recipeAdapter.GetTypes()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, types)
}

func GetMethods(c echo.Context) error {
	recipeAdapter := createRecipeAdapter(c, database.Db)
	methods, err := recipeAdapter.GetMethods()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, methods)
}
