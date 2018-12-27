package models

import (
	"database/sql"
)

type Recipe struct {
	ID   int    `json:id`
	Name string `json:name`
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

func CreateRecipe(db *sql.DB, name string) (int64, error) {
	var lastInsertId int64
	err := db.QueryRow("INSERT INTO recipes (name) VALUES ($1) RETURNING id", name).Scan(&lastInsertId)

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
