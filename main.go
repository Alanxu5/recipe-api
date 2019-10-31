package main

import (
	"log"
	"recipe-api/handler"
	"recipe-api/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db, err := models.InitDb()

	if err != nil {
		log.Panic(err)
	}

	env := &handler.Env{db}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO: need to restrict
	e.Use(middleware.CORS())

	e.GET("/recipes", env.GetAllRecipes)
	e.GET("/recipes/:id", env.GetRecipe)
	e.POST("/recipes", env.CreateRecipe)
	// e.DELETE("/recipes/:id", env.DeleteRecipe)
	e.GET("/recipes/types", env.GetTypes)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
