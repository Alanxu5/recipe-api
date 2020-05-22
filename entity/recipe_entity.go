package gateway

import (
	"encoding/json"
)

type Recipe struct {
	Id          int
	Name        string
	PrepTime    int
	CookTime    int
	Servings    int
	Method      string
	Type        string
	Description string
	Directions  json.RawMessage
}

type Ingredient struct {
	Id          int
	Amount      float32
	RecipeId    int
	Ingredient  string
	Preparation string
	Unit        string
}

type Type struct {
	Id   int
	Name string
}

type Method struct {
	Id   int
	Name string
}
