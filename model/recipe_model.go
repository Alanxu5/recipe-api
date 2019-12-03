package model

// model that represents real world values

import (
	"encoding/json"
)

type Recipe struct {
	Id          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Directions  json.RawMessage `json:"directions"`
	PrepTime    int             `json:"prepTime"`
	CookTime    int             `json:"cookTime"`
	Feeds       int             `json:"feeds"`
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
