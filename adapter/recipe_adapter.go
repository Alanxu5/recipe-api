package adapter

import (
	"github.com/labstack/echo"
	"recipe-api/gateway"
	"recipe-api/model"
)

type RecipeAdapter struct {
	recipeDBGateway gateway.RecipeDBGatewayInterface
	context         echo.Context
}

func NewRecipeAdapter(recipeDBGateway gateway.RecipeDBGatewayInterface, c echo.Context) RecipeAdapter {
	return RecipeAdapter{
		recipeDBGateway: recipeDBGateway,
		context:         c,
	}
}

func (ra RecipeAdapter) GetAllRecipes() ([]*model.Recipe, error) {
	recipes, err := ra.recipeDBGateway.GetAllRecipes()
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (ra RecipeAdapter) GetRecipe(id int) (*model.Recipe, error) {
	recipe, err := ra.recipeDBGateway.GetRecipe(id)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (ra RecipeAdapter) CreateRecipe(recipe model.Recipe) (int64, error) {
	recipeId, err := ra.recipeDBGateway.CreateRecipe(recipe)
	if err != nil {
		return 0, err
	}
	return recipeId, nil
}

func (ra RecipeAdapter) GetTypes() ([]*model.Type, error) {
	types, err := ra.recipeDBGateway.GetTypes()
	if err != nil {
		return nil, err
	}
	return types, nil
}

func (ra RecipeAdapter) GetMethods() ([]*model.Method, error) {
	methods, err := ra.recipeDBGateway.GetMethods()
	if err != nil {
		return nil, err
	}
	return methods, nil
}
