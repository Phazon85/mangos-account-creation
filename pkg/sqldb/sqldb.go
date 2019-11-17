package sqldb

import (
	"database/sql"
	"log"

	//used for testing in pg_test.go
	_ "github.com/lib/pq"
)

//New returns a sql DB object
func New(file string) *sql.DB {
	driverName, cfg := newConfig(file)
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
