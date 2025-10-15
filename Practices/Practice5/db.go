package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", "user=postgres password=KaToN2006 dbname=users_db sslmode=disable host=localhost port=5432")
	if err != nil {
		log.Fatalln("DB open error:", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalln("DB connection failed:", err)
	}

	return db
}
