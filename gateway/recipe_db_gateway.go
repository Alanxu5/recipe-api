package gateway

import (
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"recipe-api/entity"
)

type RecipeDbGatewayInterface interface {
	GetAllRecipes() ([]gateway.Recipe, error)
	GetRecipe(id int) (*gateway.Recipe, error)
	GetIngredients(id int) ([]gateway.Ingredient, error)
	GetEquipment(id int) ([]gateway.Equipment, error)
	CreateRecipe(recipe gateway.Recipe, equipment []gateway.Equipment, ingredients []gateway.Ingredient, typeId int, methodId int) (int64, error)
	GetTypes() ([]gateway.Type, error)
	GetMethods() ([]gateway.Method, error)
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

func (rg RecipeDbGateway) GetAllRecipes() ([]gateway.Recipe, error) {
	query := `SELECT r.Id, r.name, r.prep_time, r.cook_time, r.servings, m.name AS method, t.name AS type, r.description, r.directions
					FROM recipe AS r
					JOIN method AS m
					ON r.method = m.Id
					JOIN type AS t
					ON r.type = t.Id`

	rows, err := rg.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recipes := make([]gateway.Recipe, 0)
	for rows.Next() {
		recipe := gateway.Recipe{}

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
	query := `SELECT r.id, r.name, r.prep_time, r.cook_time, r.servings, m.name AS method, t.name AS type, r.description, r.directions
					FROM recipe AS r
					JOIN method AS m 
					ON r.method = m.id
					JOIN type AS t
					ON r.type = t.id
					WHERE r.id = ?`

	row := rg.Db.QueryRow(query, id)

	recipe := new(gateway.Recipe)
	err := row.Scan(&recipe.Id, &recipe.Name, &recipe.PrepTime, &recipe.CookTime,
		&recipe.Servings, &recipe.Method, &recipe.Type, &recipe.Description, &recipe.Directions)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (rg RecipeDbGateway) GetIngredients(id int) ([]gateway.Ingredient, error) {
	var ingredients []gateway.Ingredient
	query := `SELECT * FROM ingredient WHERE recipe_id = ?`
	rows, err := rg.Db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		ingredient := gateway.Ingredient{}
		err = rows.Scan(&ingredient.Id, &ingredient.Ingredient, &ingredient.RecipeId, &ingredient.Unit, &ingredient.Amount, &ingredient.Preparation)
		if err != nil {
			return nil, err
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (rg RecipeDbGateway) GetEquipment(id int) ([]gateway.Equipment, error) {
	var equipment []gateway.Equipment
	query := `SELECT e.id, e.description, e.equipment
			  FROM equipment as e 
			  INNER JOIN recipe_equipment as re 
			  ON e.id = re.equipment_id
			  INNER JOIN recipe as r
			  ON r.id = re.recipe_id
			  WHERE recipe_id = ?`
	rows, err := rg.Db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		equip := gateway.Equipment{}
		err := rows.Scan(&equip.Id, &equip.Description, &equip.Equipment)
		if err != nil {
			return nil, err
		}

		equipment = append(equipment, equip)
	}

	return equipment, nil
}

func (rg RecipeDbGateway) CreateRecipe(recipe gateway.Recipe, equipment []gateway.Equipment, ingredients []gateway.Ingredient, typeId int, methodId int) (int64, error) {
	var lastInsertId int64
	tx, txErr := rg.Db.Begin()
	if txErr != nil {
		return 0, txErr
	}
	defer tx.Rollback()

	directionsJson, err := json.Marshal(recipe.Directions)
	if err != nil {
		return 0, err
	}

	res, err := rg.Db.Exec("INSERT INTO recipe (name, description, prep_time, cook_time, servings, method, type, directions) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", recipe.Name, recipe.Description, recipe.PrepTime, recipe.CookTime, recipe.Servings, methodId, typeId, directionsJson)
	if err != nil {
		return 0, err
	}

	lastInsertId, insertError := res.LastInsertId()
	if insertError != nil {
		return 0, err
	}

	// check if equipment already exits in equipment table
	equipQuery := "SELECT id, description, equipment FROM equipment WHERE description = ? and equipment = ?"
	insertRecipeEquipQuery := "INSERT INTO recipe_equipment (recipe_id, equipment_id) VALUES (?, ?)"
	for _, equip := range equipment {
		equipRow := rg.Db.QueryRow(equipQuery, equip.Description, equip.Equipment)
		equipEntity := new(gateway.Equipment)
		equipErr := equipRow.Scan(&equipEntity.Id, &equipEntity.Description, &equipEntity.Equipment)
		// no row matches
		if equipErr == sql.ErrNoRows {
			insertEquipQuery := "INSERT INTO equipment (description, equipment) VALUES (?, ?)"
			res, err = rg.Db.Exec(insertEquipQuery, equip.Description, equip.Equipment)
			if err != nil {
				return 0, err
			}

			lastEquipInsertID, insertEquipErr := res.LastInsertId()
			if insertEquipErr != nil {
				return 0, err
			}

			res, err = rg.Db.Exec(insertRecipeEquipQuery, lastInsertId, lastEquipInsertID)
			if err != nil {
				return 0, err
			}
		} else {
			res, err = rg.Db.Exec(insertRecipeEquipQuery, lastInsertId, equipEntity.Id)
			if err != nil {
				return 0, err
			}
		}
	}

	for _, ingredient := range ingredients {
		if _, ingErr := rg.Db.Exec("INSERT INTO ingredient (food, recipe_id, unit, amount, preparation) VALUES (?, ?, ?, ?, ?)", ingredient.Ingredient, lastInsertId, ingredient.Unit, ingredient.Amount, ingredient.Preparation); ingErr != nil {
			return 0, ingErr
		}
	}

	if commErr := tx.Commit(); commErr != nil {
		return 0, commErr
	}
	return lastInsertId, nil
}

func (rg RecipeDbGateway) GetTypes() ([]gateway.Type, error) {
	query := "SELECT * FROM type"

	rows, err := rg.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	types := make([]gateway.Type, 0)
	for rows.Next() {
		recipeType := gateway.Type{}
		err := rows.Scan(&recipeType.Id, &recipeType.Name)
		if err != nil {
			return nil, err
		}

		types = append(types, recipeType)
	}

	return types, nil
}

func (rg RecipeDbGateway) GetMethods() ([]gateway.Method, error) {
	query := "SELECT * FROM method"

	rows, err := rg.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	methods := make([]gateway.Method, 0)
	for rows.Next() {
		recipeMethod := gateway.Method{}
		err := rows.Scan(&recipeMethod.Id, &recipeMethod.Name)
		if err != nil {
			return nil, err
		}

		methods = append(methods, recipeMethod)
	}

	return methods, nil
}
