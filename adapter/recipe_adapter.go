package adapter

import (
	"github.com/labstack/echo"
	"recipe-api/gateway"
	"recipe-api/model"
)

type RecipeAdapter struct {
	recipeDbGateway gateway.RecipeDbGatewayInterface
	context         echo.Context
}

func NewRecipeAdapter(recipeDbGateway gateway.RecipeDbGatewayInterface, c echo.Context) RecipeAdapter {
	return RecipeAdapter{
		recipeDbGateway: recipeDbGateway,
		context:         c,
	}
}

func (ra RecipeAdapter) GetAllRecipes() ([]*model.Recipe, error) {
	recipesSQL, err := ra.recipeDbGateway.GetAllRecipes()
	if err != nil {
		return nil, err
	}

	recipes := make([]*model.Recipe, 0)
	for _, rec := range recipesSQL {
		recipe := model.Recipe{
			Id:          rec.Id,
			Name:        rec.Name,
			Description: rec.Description,
			Equipment:   nil,
			Directions:  rec.Directions,
			Ingredients: nil,
			PrepTime:    rec.PrepTime,
			CookTime:    rec.CookTime,
			Servings:    rec.Servings,
			Type:        rec.Type,
			Method:      rec.Method,
		}
		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}

func (ra RecipeAdapter) GetRecipe(id int) (*model.Recipe, error) {
	recipe, err := ra.recipeDbGateway.GetRecipe(id)
	if err != nil {
		return nil, err
	}

	// should create a function and grab the undefined values
	rec := model.Recipe{
		Id:          recipe.Id,
		Name:        recipe.Name,
		Description: recipe.Description,
		Equipment:   nil,
		Directions:  recipe.Directions,
		Ingredients: nil,
		PrepTime:    recipe.PrepTime,
		CookTime:    recipe.CookTime,
		Servings:    recipe.Servings,
		Type:        "",
		Method:      "",
	}

	return &rec, nil
}

func (ra RecipeAdapter) CreateRecipe(recipe model.Recipe) (int64, error) {
	recipeId, err := ra.recipeDbGateway.CreateRecipe(recipe)
	if err != nil {
		return 0, err
	}

	return recipeId, nil
}

func (ra RecipeAdapter) GetTypes() ([]*model.Type, error) {
	types, err := ra.recipeDbGateway.GetTypes()
	if err != nil {
		return nil, err
	}

	var recipeTypes []*model.Type
	for _, t := range types {
		recipeType := model.Type{
			Id:   t.Id,
			Name: t.Name,
		}

		recipeTypes = append(recipeTypes, &recipeType)
	}

	return recipeTypes, nil
}

func (ra RecipeAdapter) GetMethods() ([]*model.Method, error) {
	methods, err := ra.recipeDbGateway.GetMethods()
	if err != nil {
		return nil, err
	}

	var recipeMethods []*model.Method
	for _, m := range methods {
		recipeMethod := model.Method{
			Id:   m.Id,
			Name: m.Name,
		}

		recipeMethods = append(recipeMethods, &recipeMethod)
	}

	return recipeMethods, nil
}
