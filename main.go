package main

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"recipe-api/database"
	"recipe-api/handler"
)

func main() {
	database.DB = database.InitDB()

	if database.DB == nil {
		log.Panic(errors.New("could not connect to the db"))
		return
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO: need to restrict
	e.Use(middleware.CORS())

	e.GET("/recipes", handler.GetAllRecipes)
	e.GET("/recipes/:id", handler.GetRecipe)
	e.POST("/recipes", handler.CreateRecipe)
	e.GET("/recipes/types", handler.GetTypes)
	e.GET("recipes/methods", handler.GetMethods)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
