package model

import (
	"encoding/json"
)

type Recipe struct {
	Id          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Equipment   []Equip         `json:"equipment"`
	Directions  json.RawMessage `json:"directions"`
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

type Equip struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Item        string `json:"item"`
}

type Ingredient struct {
	Id          int     `json:"id"`
	Amount      float32 `json:"amount"`
	Ingredient  string  `json:"ingredient"`
	Preparation string  `json:"preparation,omitempty"`
	Unit        string  `json:"unit,omitempty"`
}
