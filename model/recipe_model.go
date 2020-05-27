package model

import (
	"encoding/json"
)

type Recipe struct {
	Id          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Equipment   []Equipment     `json:"equipment,omitempty"`
	Directions  json.RawMessage `json:"directions,omitempty"`
	Ingredients []Ingredient    `json:"ingredients"`
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

type Equipment struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Equipment   string `json:"equipment"`
}

type Ingredient struct {
	Id          int     `json:"id"`
	Amount      float32 `json:"amount"`
	Ingredient  string  `json:"ingredient"`
	Preparation string  `json:"preparation,omitempty"`
	Unit        string  `json:"unit,omitempty"`
}
