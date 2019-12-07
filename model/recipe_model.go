package model

// model that represents real world values

import (
	"encoding/json"
	"go/types"
)

type Recipe struct {
	Id          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Directions  json.RawMessage `json:"directions"`
	Ingredients types.Array     `json:"ingredients"`
	PrepTime    int             `json:"prepTime"`
	CookTime    int             `json:"cookTime"`
	Servings    int             `json:"servings"`
	Type        string          `json:"type"`
	Method      string          `json:"method"`
}

type Type struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Method struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
