package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Datastore interface {
	GetAllRecipes() ([]*Recipe, error)
	CreateRecipe(recipe Recipe) (int64, error)
	DeleteRecipe(id int) (int64, error)
}

type DB struct {
	*sql.DB
}

func InitDb() (*DB, error) {
	config := dbConfig()
	var err error

	// all the information needed to connect to DB
	mysqlInfo := fmt.Sprintf("%s:%s@/%s",
		config["DBUSER"], config["DBPASS"], config["DBNAME"])

	println(mysqlInfo)

	// sql.Open() does not establish any connection to the DB
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		return nil, err
	}

	// db.Ping() checks if the DB is available and accessible
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")

	return &DB{db}, nil
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	const (
		dbhost = "DBHOST"
		dbport = "DBPORT"
		dbuser = "DBUSER"
		dbpass = "DBPASS"
		dbname = "DBNAME"
	)
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
