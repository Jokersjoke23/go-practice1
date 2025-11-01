package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=KaToN2006 dbname=movies_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL.")
	return db, nil
}
