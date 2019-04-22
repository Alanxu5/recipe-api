package models

import (
	"database/sql"
)

type Recipe struct {
	ID          int    `json:id`
	Name        string `json:name`
	Description string `json:description`
	Ingredients string `json:ingredients`
	Directions  string `json:directions`
	PrepTime    int    `json:prep_time`
	CookTime    int    `json:cook_time`
	Feeds       int    `json:feeds`
}

type RecipeCollection struct {
	Recipes []Recipe `json:"items"`
}

func GetRecipes(db *sql.DB) RecipeCollection {
	sql := "SELECT * FROM recipes"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := RecipeCollection{}

	for rows.Next() {
		recipe := Recipe{}
		err := rows.Scan(&recipe.ID, &recipe.Name)

		if err != nil {
			panic(err)
		}
		result.Recipes = append(result.Recipes, recipe)
	}

	return result
}

func CreateRecipe(db *sql.DB, recipe Recipe) (int64, error) {
	var lastInsertId int64
	err := db.QueryRow("INSERT INTO recipes (name, description, prep_time, cook_time, feeds) VALUES ($1, $2, $3, $4, $5) RETURNING id", recipe.Name, recipe.Description, recipe.PrepTime, recipe.CookTime, recipe.Feeds).Scan(&lastInsertId)

	if err != nil {
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
