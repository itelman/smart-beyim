package repo

import (
	"database/sql"
	"os"
	"smart-beyim/models"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInsertUser(t *testing.T) {
	// Create a temporary database file for testing
	tempDB := "temp.db"
	defer os.Remove(tempDB)

	// Open the temporary database file
	db, err := sql.Open("sqlite3", tempDB)
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	// Create the users table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name TEXT,
		password TEXT
	);
	`)
	if err != nil {
		t.Fatalf("error creating table: %v", err)
	}

	users := [10]models.User{
		{"alice123", "password123"},
		{"Hd9ZGt7A", "p$VJXg*7O#fR"},
		{"u7FZ6QcK", "!Dbn28sXv1L"},
		{"gP2rDmLk", "@A3aJ*0Kb&xW"},
		{"J3LwCpRb", "%Z1!sPqf5vXo"},
		{"k5bN6Xl3", "cU7&$DjX1lG2"},
		{"E9YfM6Xb", "Q8*uS3#tHwJr"},
		{"b2GfYpSj", "!xR5UkN9@pE7"},
		{"N6dQzGkL", "@gF2dN$1qM5X"},
		{"S4gVmC7A", "J2%yU3!cL8Tb"},
	}

	for _, user := range users {
		// Insert the user into the table
		err = InsertUser(db, user)
		if err != nil {
			t.Fatalf("error inserting user: %v", err)
		}

		// Check if the user exists in the table
		var name, password string
		err = db.QueryRow("SELECT name, password FROM users WHERE name = ?", user.Name).Scan(&name, &password)
		if err != nil {
			t.Fatalf("error querying user: %v", err)
		}

		if name != user.Name {
			t.Errorf("expected user name %s, got %s", user.Name, name)
		}

		if password != user.Password {
			t.Errorf("expected user password %s, got %s", user.Password, password)
		}
	}
}

func InsertUser(db *sql.DB, user models.User) error {
	// Execute the INSERT statement
	_, err := db.Exec("INSERT INTO users (name, password) VALUES (?, ?)",
		user.Name, user.Password)
	if err != nil {
		return err
	}
	return nil
}
