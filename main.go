package main

import (
	"database/sql"
	"fmt"
	"os"
	"recipe-api/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	initDb()

	// idiomatic to use if the db should not have a lifetime beyond the scope of the function.
	defer db.Close()

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

func initDb() {
	config := dbConfig()
	var err error

	// all the information needed to connect to DB
	mysqlInfo := fmt.Sprintf("%s:%s@/%s",
		config[dbuser], config[dbpass], config[dbname])

	// sql.Open() does not establish any connection to the DB
	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}

	// db.Ping() checks if the DB is available and accessible
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}
