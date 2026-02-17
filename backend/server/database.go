package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDb() error {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?connect_timeout=10&sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Encountered an error while trying to connect to database %s\n", os.Getenv("DB_NAME"))
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Encountered an error with connection to dabase %s\n", os.Getenv("DB_NAME"))
		log.Fatal(err)
	}

	log.Printf("Successfully connected to database %s\n", os.Getenv("DB_NAME"))
	return nil
}
