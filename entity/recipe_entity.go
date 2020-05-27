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
	Ingredient  string
	RecipeId    int
	Unit        string
	Amount      float32
	Preparation string
}

type Equipment struct {
	Id          int
	Description string
	Equipment   string
}

type Type struct {
	Id   int
	Name string
}

type Method struct {
	Id   int
	Name string
}
