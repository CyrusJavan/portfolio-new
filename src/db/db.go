package db

import (
	"os"

	"github.com/apex/log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var conn *sqlx.DB

// GetInstance returns a connection to the database.
func GetInstance(c *gin.Context) *sqlx.DB {
	// Use an existing connection if you can
	if conn != nil {
		return conn
	}

	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		l := c.Keys["logEntry"].(*log.Entry)
		l.WithError(err).WithFields(log.Fields{
			"func": "GetInstance",
		}).Error("could not open db connection")
	}

	conn = db
	return db
}

// GetInstanceNoContext returns a connection to the database without using a
// gin.Context for logging.
func GetInstanceNoContext() *sqlx.DB {
	if conn != nil {
		return conn
	}

	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.WithError(err).Error("could not open DB connection")
	}

	conn = db
	return db
}
