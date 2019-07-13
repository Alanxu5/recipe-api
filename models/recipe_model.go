package models

import (
	"database/sql"
	"encoding/json"
)

type Recipe struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Ingredients string          `json:"ingredients"`
	Directions  json.RawMessage `json:"directions"`
	PrepTime    int             `json:"prepTime"`
	CookTime    int             `json:"cookTime"`
	Feeds       int             `json:"feeds"`
	Type        int             `json:"type"`
	Method      int             `json:"method"`
}

type RecipeCollection struct {
	Recipes []Recipe `json:"items"`
}

func GetRecipes(db *sql.DB) RecipeCollection {
	sql := "SELECT * FROM recipe"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := RecipeCollection{}

	for rows.Next() {
		recipe := Recipe{}
		// err := rows.Scan(&recipe.ID, &recipe.Name)

		if err != nil {
			panic(err)
		}
		result.Recipes = append(result.Recipes, recipe)
	}

	return result
}

func CreateRecipe(db *sql.DB, recipe Recipe) (int64, error) {
	jsonString, err := json.Marshal(recipe.Directions)
	if err != nil {
		panic(err)
	}

	var lastInsertId int64
	res, err := db.Exec("INSERT INTO recipe (name, description, prep_time, cook_time, servings, method, type, directions) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", recipe.Name, recipe.Description, recipe.PrepTime, recipe.CookTime, recipe.Feeds, recipe.Method, recipe.Type, jsonString)
	if err != nil {
		panic(err)
	}

	lastInsertId, error := res.LastInsertId()

	if error != nil {
		panic(err)
	}

	return lastInsertId, nil
}

func DeleteRecipe(db *sql.DB, id int) (int64, error) {
	var deletedId int64
	err := db.QueryRow("DELETE FROM recipes WHERE id = $1 RETURNING id", id).Scan(&deletedId)

	if err != nil {
		panic(err)
	}

	return deletedId, nil
}
