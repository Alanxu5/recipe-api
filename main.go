package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"recipe-api/gateway"
	"recipe-api/handler"
)

func main() {
	db, err := gateway.InitDB()

	if err != nil {
		log.Panic(err)
	}

	env := &handler.Env{
		DB: db,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO: need to restrict
	e.Use(middleware.CORS())

	e.GET("/recipes", env.GetAllRecipes)
	e.GET("/recipes/:id", env.GetRecipe)
	e.POST("/recipes", env.CreateRecipe)
	e.GET("/recipes/types", env.GetTypes)
	e.GET("recipes/methods", env.GetMethods)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
