package database

import (
	"database/sql"
	"log"
	"smart-beyim/models"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() *sql.DB {
	db, err := sql.Open("sqlite3", "database/storage.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InsertUser(user models.User) {
	db := OpenDB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", user.ID, user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
	}
}