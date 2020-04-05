package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// Postgres
	_ "github.com/lib/pq"
)

var conn *sqlx.DB

// GetInstance of the db
func GetInstance() *sqlx.DB {
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
