package model

import (
	"encoding/json"
)

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
		err := rows.Scan(&recipe.Id, &recipe.Name, &recipe.PrepTime, &recipe.CookTime,
			&recipe.Feeds, &recipe.Method, &recipe.Type, &recipe.Description, &recipe.Directions)

		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (db *DB) GetRecipe(id int) (*Recipe, error) {
	sql := `SELECT r.id, r.name, r.prep_time, r.cook_time, r.servings, m.name AS method, rt.name AS type, r.description, r.directions
					FROM recipe AS r
					JOIN method AS m 
					ON r.method = m.id
					JOIN type AS rt
					ON r.type = rt.id
					WHERE r.id = ?`

	row := db.QueryRow(sql, id)

	recipe := new(Recipe)

	// has to be in the same order as DB columns
	err := row.Scan(&recipe.Id, &recipe.Name, &recipe.PrepTime, &recipe.CookTime,
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

	methodSql := "SELECT * FROM method WHERE name = ?"
	row := db.QueryRow(methodSql, recipe.Method)
	recipeMethod := new(Method)
	errMethod := row.Scan(&recipeMethod.Id, &recipeMethod.Name)

	if errMethod != nil {
		return 0, errMethod
	}

	typeSql := "SELECT * FROM type WHERE name = ?"
	typeRow := db.QueryRow(typeSql, recipe.Type)
	recipeType := new(Type)
	errType := typeRow.Scan(&recipeType.Id, &recipeType.Name)

	if errType != nil {
		return 0, errType
	}

	var lastInsertId int64
	res, err := db.Exec("INSERT INTO recipe (name, description, prep_time, cook_time, servings, method, type, directions) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", recipe.Name, recipe.Description, recipe.PrepTime, recipe.CookTime, recipe.Feeds, recipeMethod.Id, recipeType.Id, jsonString)
	if err != nil {
		return 0, err
	}

	lastInsertId, error := res.LastInsertId()

	if error != nil {
		return 0, err
	}

	return lastInsertId, nil
}

// TODO - implement
func (db *DB) DeleteRecipe(id int) (int64, error) {
	var deletedId int64
	err := db.QueryRow("DELETE FROM recipe WHERE id = $1 RETURNING id", id).Scan(&deletedId)

	if err != nil {
		return 0, err
	}

	return deletedId, nil
}

func (db *DB) GetTypes() ([]*Type, error) {
	sql := "SELECT * FROM type"

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	types := make([]*Type, 0)

	for rows.Next() {
		recipeType := new(Type)

		err := rows.Scan(&recipeType.Id, &recipeType.Name)

		if err != nil {
			return nil, err
		}
		types = append(types, recipeType)
	}

	return types, nil
}

func (db *DB) GetMethods() ([]*Method, error) {
	sql := "SELECT * FROM method"

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	methods := make([]*Method, 0)

	for rows.Next() {
		recipeMethod := new(Method)

		err := rows.Scan(&recipeMethod.Id, &recipeMethod.Name)

		if err != nil {
			return nil, err
		}
		methods = append(methods, recipeMethod)
	}

	return methods, nil
}
