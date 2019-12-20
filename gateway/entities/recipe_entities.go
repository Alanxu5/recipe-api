package gateway

// database models

import "encoding/json"

type RecipeSQL struct {
	Name        string
	Description string
	Directions  json.RawMessage
	PrepTime    int
	CookTime    int
	Servings    int
	Type        int
	Method      string
}

type IngredientSQL struct {
	ID          int
	Amount      float32
	RecipeID    int
	Ingredient  string
	Preparation string
	Unit        string
}
