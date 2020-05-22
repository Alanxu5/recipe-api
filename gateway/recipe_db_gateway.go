package gateway

import (
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo"
	"recipe-api/entity"
	"recipe-api/model"
)

type RecipeDbGatewayInterface interface {
	GetAllRecipes() ([]*gateway.Recipe, error)
	GetRecipe(id int) (*gateway.Recipe, error)
	CreateRecipe(recipe model.Recipe) (int64, error)
	GetTypes() ([]*gateway.Type, error)
	GetMethods() ([]*gateway.Method, error)
}

type RecipeDbGateway struct {
	Context echo.Context
	Db      *sql.DB
}

func NewRecipeDbGateway(context echo.Context, db *sql.DB) RecipeDbGateway {
	return RecipeDbGateway{
		Context: context,
		Db:      db,
	}
}

func (rg RecipeDbGateway) GetAllRecipes() ([]*gateway.Recipe, error) {
	sql := `SELECT r.Id, r.name, r.prep_time, r.cook_time, r.servings, m.name AS method, rt.name AS type, r.description, r.directions
					FROM recipe AS r
					JOIN method AS m
					ON r.method = m.Id
					JOIN type AS rt
					ON r.type = rt.Id`

	rows, err := rg.Db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recipes := make([]*gateway.Recipe, 0)
	for rows.Next() {
		recipe := new(gateway.Recipe)

		// has to be in the same order as Db columns
		err := rows.Scan(&recipe.Id, &recipe.Name, &recipe.PrepTime, &recipe.CookTime,
			&recipe.Servings, &recipe.Method, &recipe.Type, &recipe.Description, &recipe.Directions)

		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (rg RecipeDbGateway) GetRecipe(id int) (*gateway.Recipe, error) {
	recSql := `SELECT r.id, r.name, r.prep_time, r.cook_time, r.servings, m.name AS method, rt.name AS type, r.description, r.directions
					FROM recipe AS r
					JOIN method AS m 
					ON r.method = m.id
					JOIN type AS rt
					ON r.type = rt.id
					WHERE r.id = ?`

	row := rg.Db.QueryRow(recSql, id)

	recipe := new(gateway.Recipe)
	err := row.Scan(&recipe.Id, &recipe.Name, &recipe.PrepTime, &recipe.CookTime,
		&recipe.Servings, &recipe.Method, &recipe.Type, &recipe.Description, &recipe.Directions)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (rg RecipeDbGateway) GetIngredients(id int) (*[]gateway.Ingredient, error) {
	var ingredients []gateway.Ingredient
	ingSql := `SELECT * FROM ingredient WHERE recipe_id = ?`
	rows, err := rg.Db.Query(ingSql, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ingredient gateway.Ingredient
		err = rows.Scan(&ingredient.Id, &ingredient.Ingredient, &ingredient.RecipeId, &ingredient.Unit, &ingredient.Amount, &ingredient.Preparation)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}

	return &ingredients, nil
}

func (rg RecipeDbGateway) CreateRecipe(recipe model.Recipe) (int64, error) {
	jsonString, err := json.Marshal(recipe.Directions)
	if err != nil {
		return 0, err
	}

	methodSql := "SELECT * FROM method WHERE name = ?"
	row := rg.Db.QueryRow(methodSql, recipe.Method)
	recipeMethod := new(model.Method)
	errMethod := row.Scan(&recipeMethod.Id, &recipeMethod.Name)

	if errMethod != nil {
		return 0, errMethod
	}

	typeSql := "SELECT * FROM type WHERE name = ?"
	typeRow := rg.Db.QueryRow(typeSql, recipe.Type)
	recipeType := new(model.Type)
	errType := typeRow.Scan(&recipeType.Id, &recipeType.Name)

	if errType != nil {
		return 0, errType
	}

	var lastInsertId int64
	tx, txErr := rg.Db.Begin()
	if txErr != nil {
		return 0, txErr
	}
	defer tx.Rollback()

	res, err := rg.Db.Exec("INSERT INTO recipe (name, description, prep_time, cook_time, servings, method, type, directions) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", recipe.Name, recipe.Description, recipe.PrepTime, recipe.CookTime, recipe.Servings, recipeMethod.Id, recipeType.Id, jsonString)
	if err != nil {
		return 0, err
	}

	lastInsertId, insertError := res.LastInsertId()

	if insertError != nil {
		return 0, err
	}

	// check if equipment already exits in equipment table
	equipSql := "SELECT id, description, equipment FROM equipment WHERE description = ? and equipment = ?"
	insertRecipeEquipSql := "INSERT INTO recipe_equipment (recipe_id, equipment_id) VALUES (?, ?)"
	for _, equip := range recipe.Equipment {
		equipRow := rg.Db.QueryRow(equipSql, equip.Description, equip.Item)
		equipModel := new(model.Equip)
		equipErr := equipRow.Scan(&equipModel.Id, &equipModel.Description, &equipModel.Item)
		if equipErr == sql.ErrNoRows {
			insertEquipSql := "INSERT INTO equipment (description, equipment) VALUES (?, ?)"
			res, err = rg.Db.Exec(insertEquipSql, equip.Description, equip.Item)
			if err != nil {
				return 0, err
			}
			lastEquipInsertID, insertEquipErr := res.LastInsertId()
			if insertEquipErr != nil {
				return 0, err
			}
			res, err = rg.Db.Exec(insertRecipeEquipSql, lastInsertId, lastEquipInsertID)
			if err != nil {
				return 0, err
			}
		}
		res, err = rg.Db.Exec(insertRecipeEquipSql, lastInsertId, equipModel.Id)
		if err != nil {
			return 0, err
		}
	}

	for _, ingredient := range recipe.Ingredients {
		if _, ingErr := rg.Db.Exec("INSERT INTO ingredient (food, recipe_id, unit, amount, preparation) VALUES (?, ?, ?, ?, ?)", ingredient.Ingredient, lastInsertId, ingredient.Unit, ingredient.Amount, ingredient.Preparation); ingErr != nil {
			return 0, ingErr
		}
	}

	if commErr := tx.Commit(); commErr != nil {
		return 0, commErr
	}
	return lastInsertId, nil
}

func (rg RecipeDbGateway) GetTypes() ([]*gateway.Type, error) {
	sql := "SELECT * FROM type"

	rows, err := rg.Db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	types := make([]*gateway.Type, 0)

	for rows.Next() {
		recipeType := new(gateway.Type)

		err := rows.Scan(&recipeType.Id, &recipeType.Name)
		if err != nil {
			return nil, err
		}

		types = append(types, recipeType)
	}

	return types, nil
}

func (rg RecipeDbGateway) GetMethods() ([]*gateway.Method, error) {
	sql := "SELECT * FROM method"

	rows, err := rg.Db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	methods := make([]*gateway.Method, 0)

	for rows.Next() {
		recipeMethod := new(gateway.Method)

		err := rows.Scan(&recipeMethod.Id, &recipeMethod.Name)
		if err != nil {
			return nil, err
		}

		methods = append(methods, recipeMethod)
	}

	return methods, nil
}
