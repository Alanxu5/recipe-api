package adapter

import (
	"github.com/labstack/echo/v4"
	"recipe-api/adapter/converter"
	"recipe-api/gateway"
	"recipe-api/model"
)

type RecipeAdapterInterface interface {
	GetAllRecipes() ([]model.Recipe, error)
	GetRecipe(id int) (*model.Recipe, error)
	CreateRecipe(recipe model.Recipe) (int64, error)
	GetTypes() ([]model.Type, error)
	GetMethods() ([]model.Method, error)
}

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

func (ra RecipeAdapter) GetAllRecipes() ([]model.Recipe, error) {
	recipes, err := ra.recipeDbGateway.GetAllRecipes()
	if err != nil {
		return nil, err
	}

	var recs []model.Recipe
	for _, rec := range recipes {
		ingredients, ingErr := ra.recipeDbGateway.GetIngredients(rec.Id)
		if ingErr != nil {
			return nil, ingErr
		}

		convertedRec, convertErr := converter.ConvertRecipe(&rec, ingredients, nil)
		if convertErr != nil {
			return nil, convertErr
		}

		recs = append(recs, convertedRec)
	}

	return recs, nil
}

func (ra RecipeAdapter) GetRecipe(id int) (*model.Recipe, error) {
	recipe, err := ra.recipeDbGateway.GetRecipe(id)
	if err != nil {
		return nil, err
	}

	ingredients, ingErr := ra.recipeDbGateway.GetIngredients(id)
	if ingErr != nil {
		return nil, err
	}

	equipment, equipErr := ra.recipeDbGateway.GetEquipment(id)
	if equipErr != nil {
		return nil, equipErr
	}

	rec, convertErr := converter.ConvertRecipe(recipe, ingredients, equipment)
	if convertErr != nil {
		return nil, err
	}

	return &rec, nil
}

func (ra RecipeAdapter) CreateRecipe(recipe model.Recipe) (int64, error) {
	recipeEntity := converter.ConvertRecipeToEntity(recipe)
	equipEntity := converter.ConvertEquipToEntity(recipe.Equipment)
	ingredientEntity := converter.ConvertIngredientsToEntity(recipe.Ingredients)

	typeEntities, typesErr := ra.recipeDbGateway.GetTypes()
	if typesErr != nil {
		return 0, typesErr
	}

	methodEntities, methodsErr := ra.recipeDbGateway.GetMethods()
	if methodsErr != nil {
		return 0, methodsErr
	}

	typeId, typeErr := converter.ConvertTypeStringToId(recipe.Type, typeEntities)
	if typeErr != nil {
		return 0, typeErr
	}

	methodId, methodErr := converter.ConvertMethodStringToId(recipe.Method, methodEntities)
	if methodErr != nil {
		return 0, methodErr
	}

	recipeId, createErr := ra.recipeDbGateway.CreateRecipe(recipeEntity, equipEntity, ingredientEntity, typeId, methodId)
	if createErr != nil {
		return 0, createErr
	}

	return recipeId, nil
}

func (ra RecipeAdapter) GetTypes() ([]model.Type, error) {
	types, err := ra.recipeDbGateway.GetTypes()
	if err != nil {
		return nil, err
	}

	recipeTypes, convertErr := converter.ConvertTypes(types)
	if convertErr != nil {
		return nil, convertErr
	}

	return recipeTypes, nil
}

func (ra RecipeAdapter) GetMethods() ([]model.Method, error) {
	methods, err := ra.recipeDbGateway.GetMethods()
	if err != nil {
		return nil, err
	}

	recipeMethods, convertErr := converter.ConvertMethods(methods)
	if convertErr != nil {
		return nil, convertErr
	}

	return recipeMethods, nil
}
