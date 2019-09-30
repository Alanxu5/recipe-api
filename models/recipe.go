package models

import (
	"encoding/json"
)

type Recipe struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Directions  json.RawMessage `json:"directions"`
	PrepTime    int             `json:"prepTime"`
	CookTime    int             `json:"cookTime"`
	Feeds       int             `json:"feeds"`
	Type        int             `json:"type"`
	Method      int             `json:"method"`
}

func (db *DB) GetAllRecipes() ([]*Recipe, error) {
	sql := "SELECT * FROM recipe"

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recipes := make([]*Recipe, 0)
	for rows.Next() {
		recipe := new(Recipe)

		// has to be in the same order as DB columns
		err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.PrepTime, &recipe.CookTime,
			&recipe.Feeds, &recipe.Method, &recipe.Type, &recipe.Description, &recipe.Directions)

		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (db *DB) GetRecipe(id int) (*Recipe, error) {
	sql := "SELECT * FROM recipe WHERE id = ?"

	row := db.QueryRow(sql, id)

	recipe := new(Recipe)

	// has to be in the same order as DB columns
	err := row.Scan(&recipe.ID, &recipe.Name, &recipe.PrepTime, &recipe.CookTime,
		&recipe.Feeds, &recipe.Method, &recipe.Type, &recipe.Description, &recipe.Directions)

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (db *DB) CreateRecipe(recipe Recipe) (int64, error) {
	jsonString, err := json.Marshal(recipe.Directions)
	if err != nil {
		return 0, err
	}

	var lastInsertId int64
	res, err := db.Exec("INSERT INTO recipe (name, description, prep_time, cook_time, servings, method, type, directions) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", recipe.Name, recipe.Description, recipe.PrepTime, recipe.CookTime, recipe.Feeds, recipe.Method, recipe.Type, jsonString)
	if err != nil {
		return 0, err
	}

	lastInsertId, error := res.LastInsertId()

	if error != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (db *DB) DeleteRecipe(id int) (int64, error) {
	var deletedId int64
	err := db.QueryRow("DELETE FROM recipes WHERE id = $1 RETURNING id", id).Scan(&deletedId)

	if err != nil {
		return 0, err
	}

	return deletedId, nil
}
