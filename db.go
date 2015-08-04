package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Connect() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./pre-sales-huddle.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
