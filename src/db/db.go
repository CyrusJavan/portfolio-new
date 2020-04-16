package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var conn *sqlx.DB

// GetInstance returns a connection to the database.
func GetInstance() *sqlx.DB {
	// Use an existing connection if you can
	if conn != nil {
		return conn
	}

	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	conn = db
	return db
}
