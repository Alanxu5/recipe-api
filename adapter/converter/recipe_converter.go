package converter

import (
	"errors"
	gateway "recipe-api/entity"
	"recipe-api/model"
)

func ConvertRecipe(recipe *gateway.Recipe, ingredients []gateway.Ingredient, equipment []gateway.Equipment) (model.Recipe, error) {
	ing, err := ConvertIngredients(ingredients)
	if err != nil {
		return model.Recipe{}, err
	}

	equip, err := ConvertEquipment(equipment)

	rec := model.Recipe{
		Id:          recipe.Id,
		Name:        recipe.Name,
		Description: recipe.Description,
		Equipment:   equip,
		Directions:  recipe.Directions,
		Ingredients: ing,
		PrepTime:    recipe.PrepTime,
		CookTime:    recipe.CookTime,
		Servings:    recipe.Servings,
		Type:        recipe.Type,
		Method:      recipe.Method,
	}

	return rec, nil
}

func ConvertIngredients(ingredients []gateway.Ingredient) ([]model.Ingredient, error) {
	var recipeIngredients []model.Ingredient
	for _, ing := range ingredients {
		recipeIng := model.Ingredient{
			Id:          ing.Id,
			Amount:      ing.Amount,
			Ingredient:  ing.Ingredient,
			Preparation: ing.Preparation,
			Unit:        ing.Unit,
		}
		recipeIngredients = append(recipeIngredients, recipeIng)
	}

	return recipeIngredients, nil
}

func ConvertEquipment(equipment []gateway.Equipment) ([]model.Equipment, error) {
	var recipeEquipment []model.Equipment
	for _, ing := range equipment {
		recipeEquip := model.Equipment{
			Id:          ing.Id,
			Description: ing.Description,
			Equipment:   ing.Equipment,
		}
		recipeEquipment = append(recipeEquipment, recipeEquip)
	}

	return recipeEquipment, nil
}

func ConvertTypes(types []gateway.Type) ([]model.Type, error) {
	var recipeTypes []model.Type
	for _, t := range types {
		recipeType := model.Type{
			Id:   t.Id,
			Name: t.Name,
		}
		recipeTypes = append(recipeTypes, recipeType)
	}

	return recipeTypes, nil
}

func ConvertMethods(methods []gateway.Method) ([]model.Method, error) {
	var recipeMethods []model.Method
	for _, m := range methods {
		recipeMethod := model.Method{
			Id:   m.Id,
			Name: m.Name,
		}
		recipeMethods = append(recipeMethods, recipeMethod)
	}

	return recipeMethods, nil
}

func ConvertRecipeToEntity(recipe model.Recipe) gateway.Recipe {
	rec := gateway.Recipe{
		Id:          recipe.Id,
		Name:        recipe.Name,
		PrepTime:    recipe.PrepTime,
		CookTime:    recipe.CookTime,
		Servings:    recipe.Servings,
		Method:      recipe.Method,
		Type:        recipe.Type,
		Description: recipe.Description,
		Directions:  recipe.Directions,
	}

	return rec
}

func ConvertEquipToEntity(equipment []model.Equipment) []gateway.Equipment {
	var equipList []gateway.Equipment
	for _, e := range equipment {
		equip := gateway.Equipment{
			Id:          e.Id,
			Description: e.Description,
			Equipment:   e.Equipment,
		}
		equipList = append(equipList, equip)
	}

	return equipList
}

func ConvertIngredientsToEntity(ingredients []model.Ingredient) []gateway.Ingredient {
	var ingredientList []gateway.Ingredient
	for _, i := range ingredients {
		ingredient := gateway.Ingredient{
			Id:          i.Id,
			Ingredient:  i.Ingredient,
			RecipeId:    0,
			Unit:        i.Unit,
			Amount:      i.Amount,
			Preparation: i.Preparation,
		}
		ingredientList = append(ingredientList, ingredient)
	}

	return ingredientList
}

func ConvertTypeStringToId(recType string, types []gateway.Type) (int, error) {
	for t := range types {
		if types[t].Name == recType {
			return types[t].Id, nil
		}
	}

	return 0, errors.New("recipe type could not be converted to id")
}

func ConvertMethodStringToId(recMethod string, methods []gateway.Method) (int, error) {
	for m := range methods {
		if methods[m].Name == recMethod {
			return methods[m].Id, nil
		}
	}

	return 0, errors.New("recipe method could not be converted to id")
}
