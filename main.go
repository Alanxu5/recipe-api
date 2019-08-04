package main

import (
	"recipe-api/handlers"
	"recipe-api/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db := db.InitDb()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO: need to restrict
	e.Use(middleware.CORS())

	e.GET("/recipes", handlers.GetAllRecipes(db))
	e.POST("/recipes", handlers.CreateRecipe(db))
	e.DELETE("/recipes/:id", handlers.DeleteRecipe(db))

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
