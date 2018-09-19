package persistence

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"database/sql"
	)

// Connect ...
func Connect() *sql.DB {
	resource, err := sql.Open("sqlite3", "simple-go-rest.db")
	if err != nil {
		log.Fatal(err)
	}
	createSchemaIfNotExists(resource)
	return resource
}

func createSchemaIfNotExists(resource *sql.DB) {
	query := "CREATE TABLE IF NOT EXISTS customers (Id INTEGER PRIMARY KEY, Name TEXT, Email TEXT, Age INTEGER, Gender TEXT)"
	stmt, err := resource.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
}
