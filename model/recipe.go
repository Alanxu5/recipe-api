package model

import (
	"encoding/json"
	gateway "recipe-api/gateway/entities"
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
			&recipe.Servings, &recipe.Method, &recipe.Type, &recipe.Description, &recipe.Directions)

		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (db *DB) GetRecipe(id int) (*Recipe, error) {
	recSql := `SELECT r.id, r.name, r.prep_time, r.cook_time, r.servings, m.name AS method, rt.name AS type, r.description, r.directions
					FROM recipe AS r
					JOIN method AS m 
					ON r.method = m.id
					JOIN type AS rt
					ON r.type = rt.id
					WHERE r.id = ?`
	row := db.QueryRow(recSql, id)

	ingSql := `SELECT * FROM ingredient WHERE recipe_id = ?`
	rows, queryErr := db.Query(ingSql, id)
	if queryErr != nil {
		return nil, queryErr
	}

	var ingredients []Ingredient
	for rows.Next() {
		var ingredientSQL gateway.IngredientSQL
		errScan := rows.Scan(&ingredientSQL.ID, &ingredientSQL.Ingredient, &ingredientSQL.RecipeID, &ingredientSQL.Unit, &ingredientSQL.Amount, &ingredientSQL.Preparation)
		if errScan != nil {
			return nil, errScan
		}
		ingredients = append(ingredients, Ingredient{Ingredient: ingredientSQL.Ingredient, Unit: ingredientSQL.Unit, Amount: ingredientSQL.Amount, Preparation: ingredientSQL.Preparation})
	}

	recipe := new(Recipe)
	recipe.Ingredients = ingredients
	// has to be in the same order as DB columns
	err := row.Scan(&recipe.Id, &recipe.Name, &recipe.PrepTime, &recipe.CookTime,
		&recipe.Servings, &recipe.Method, &recipe.Type, &recipe.Description, &recipe.Directions)

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
	tx, txErr := db.Begin()
	if txErr != nil {
		return 0, txErr
	}
	defer tx.Rollback()

	res, err := db.Exec("INSERT INTO recipe (name, description, prep_time, cook_time, servings, method, type, directions) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", recipe.Name, recipe.Description, recipe.PrepTime, recipe.CookTime, recipe.Servings, recipeMethod.Id, recipeType.Id, jsonString)
	if err != nil {
		return 0, err
	}

	lastInsertId, insertError := res.LastInsertId()

	if insertError != nil {
		return 0, err
	}

	for _, ingredient := range recipe.Ingredients {
		if _, ingErr := db.Exec("INSERT INTO ingredient (food, recipe_id, unit, amount, preparation) VALUES (?, ?, ?, ?, ?)", ingredient.Ingredient, lastInsertId, ingredient.Unit, ingredient.Amount, ingredient.Preparation); ingErr != nil {
			return 0, ingErr
		}
	}

	if commErr := tx.Commit(); commErr != nil {
		return 0, commErr
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
