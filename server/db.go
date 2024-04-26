package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func StartDB() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))
	log.Println("Connecting Database...")

	if err != nil {
		log.Println("Error! Unable to create Database", err)

		return nil, err
	}

	const create string = `
  CREATE TABLE IF NOT EXISTS activities (
  id INTEGER NOT NULL PRIMARY KEY,
  time DATETIME NOT NU	LL,
  description TEXT
  );`

	if _, err := db.Exec(create); err != nil {
		log.Println("Error! Unable to run script: ", err)

		return nil, err
	}

	return db, nil
}
