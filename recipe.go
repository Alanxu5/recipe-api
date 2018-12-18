package main

import (
	"database/sql"
	"fmt"
	"os"
	"recipe/api/handlers"

	"github.com/labstack/echo"
	_ "github.com/lib/pq" // loading the driver anonymously, using _ so none of its exported names are visible
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
	e.GET("/recipes", handlers.CreateRecipe(db))
	e.Logger.Fatal(e.Start(":8000"))
}

func initDb() {
	config := dbConfig()
	var err error

	// all the information needed to connect to DB
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	// sql.Open() does not establish any connection to the DB
	db, err = sql.Open("postgres", psqlInfo)
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
