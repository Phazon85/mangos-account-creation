package pg

import (
	"database/sql"
	"log"

	//used for testing in pg_test.go
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

//NewSQLDBObject returns a sql DB object, currently only supports postgres.
// TODO: Look into MySQL and MicrosoftSQL support
func NewSQLDBObject(file string) *sql.DB {
	cfg := newConfig(file)
	db, err := sql.Open(driverName, cfg)
	if err != nil {
		log.Printf("Error making config string or opening SQL DB: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging postgres db connection: %s", err.Error())
	}
	log.Println("Successfully connected.")
	return db
}
