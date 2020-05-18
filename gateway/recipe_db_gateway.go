package gateway

import (
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo"
	"recipe-api/database"
	"recipe-api/entity"
	. "recipe-api/model"
)

type RecipeDBGatewayInterface interface {
	GetAllRecipes() ([]*Recipe, error)
	GetRecipe(id int) (*Recipe, error)
	CreateRecipe(recipe Recipe) (int64, error)
	GetTypes() ([]*Type, error)
	GetMethods() ([]*Method, error)
}

type RecipeDBGateway struct {
	Context echo.Context
}

func NewRecipeDBGateway(context echo.Context) RecipeDBGateway {
	return RecipeDBGateway{
		Context: context,
	}
}

func (rg RecipeDBGateway) GetAllRecipes() ([]*Recipe, error) {
	sql := `SELECT r.id, r.name, r.prep_time, r.cook_time, r.servings, m.name AS method, rt.name AS type, r.description, r.directions
					FROM recipe AS r
					JOIN method AS m
					ON r.method = m.id
					JOIN type AS rt
					ON r.type = rt.id`

	rows, err := database.DB.Query(sql)

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

func (rg RecipeDBGateway) GetRecipe(id int) (*Recipe, error) {
	recSql := `SELECT r.id, r.name, r.prep_time, r.cook_time, r.servings, m.name AS method, rt.name AS type, r.description, r.directions
					FROM recipe AS r
					JOIN method AS m 
					ON r.method = m.id
					JOIN type AS rt
					ON r.type = rt.id
					WHERE r.id = ?`
	row := database.DB.QueryRow(recSql, id)

	ingSql := `SELECT * FROM ingredient WHERE recipe_id = ?`
	rows, queryErr := database.DB.Query(ingSql, id)
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

func (rg RecipeDBGateway) CreateRecipe(recipe Recipe) (int64, error) {
	jsonString, err := json.Marshal(recipe.Directions)
	if err != nil {
		return 0, err
	}

	methodSql := "SELECT * FROM method WHERE name = ?"
	row := database.DB.QueryRow(methodSql, recipe.Method)
	recipeMethod := new(Method)
	errMethod := row.Scan(&recipeMethod.Id, &recipeMethod.Name)

	if errMethod != nil {
		return 0, errMethod
	}

	typeSql := "SELECT * FROM type WHERE name = ?"
	typeRow := database.DB.QueryRow(typeSql, recipe.Type)
	recipeType := new(Type)
	errType := typeRow.Scan(&recipeType.Id, &recipeType.Name)

	if errType != nil {
		return 0, errType
	}

	var lastInsertId int64
	tx, txErr := database.DB.Begin()
	if txErr != nil {
		return 0, txErr
	}
	defer tx.Rollback()

	res, err := database.DB.Exec("INSERT INTO recipe (name, description, prep_time, cook_time, servings, method, type, directions) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", recipe.Name, recipe.Description, recipe.PrepTime, recipe.CookTime, recipe.Servings, recipeMethod.Id, recipeType.Id, jsonString)
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
		equipRow := database.DB.QueryRow(equipSql, equip.Description, equip.Item)
		equipModel := new(Equip)
		equipErr := equipRow.Scan(&equipModel.ID, &equipModel.Description, &equipModel.Item)
		if equipErr == sql.ErrNoRows {
			insertEquipSql := "INSERT INTO equipment (description, equipment) VALUES (?, ?)"
			res, err = database.DB.Exec(insertEquipSql, equip.Description, equip.Item)
			if err != nil {
				return 0, err
			}
			lastEquipInsertID, insertEquipErr := res.LastInsertId()
			if insertEquipErr != nil {
				return 0, err
			}
			res, err = database.DB.Exec(insertRecipeEquipSql, lastInsertId, lastEquipInsertID)
			if err != nil {
				return 0, err
			}
		}
		res, err = database.DB.Exec(insertRecipeEquipSql, lastInsertId, equipModel.ID)
		if err != nil {
			return 0, err
		}
	}

	for _, ingredient := range recipe.Ingredients {
		if _, ingErr := database.DB.Exec("INSERT INTO ingredient (food, recipe_id, unit, amount, preparation) VALUES (?, ?, ?, ?, ?)", ingredient.Ingredient, lastInsertId, ingredient.Unit, ingredient.Amount, ingredient.Preparation); ingErr != nil {
			return 0, ingErr
		}
	}

	if commErr := tx.Commit(); commErr != nil {
		return 0, commErr
	}
	return lastInsertId, nil
}

func (rg RecipeDBGateway) GetTypes() ([]*Type, error) {
	sql := "SELECT * FROM type"

	rows, err := database.DB.Query(sql)

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

func (rg RecipeDBGateway) GetMethods() ([]*Method, error) {
	sql := "SELECT * FROM method"

	rows, err := database.DB.Query(sql)

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
