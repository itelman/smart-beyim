package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

func NewDB(storagePath string) (*Sqlite, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Ensure the database connection is closed if the subsequent initialization fails
	defer func() {
		if err != nil {
			db.Close()
		}
	}()

	tableCreationQueries := []string{
		`CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
            created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS test_type (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL
        );`,
		`CREATE TABLE IF NOT EXISTS tests (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
            test_type_id INTEGER,
            score FLOAT NOT NULL,
            feedback TEXT NOT NULL,
			time_spent TIMESTAMP,
			time_passed TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id),
            FOREIGN KEY (test_type_id) REFERENCES test_type(id)
        );`,
		`CREATE TABLE IF NOT EXISTS ielts (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
			time_passed TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			reading_id INTEGER,
			listening_id INTEGER,
			writing_id INTEGER,
			speaking_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (reading_id) REFERENCES tests(id),
			FOREIGN KEY (listening_id) REFERENCES tests(id),
			FOREIGN KEY (writing_id) REFERENCES tests(id),
			FOREIGN KEY (speaking_id) REFERENCES tests(id)
        );`,
		`CREATE TABLE IF NOT EXISTS chat_requests (
			id INTEGER PRIMARY KEY,
			user_id INTEGER UNIQUE,
			request_data TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)			
		);`,
		`CREATE TABLE IF NOT EXISTS teachers (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS teacher_user (
			teacher_id INTEGER,
			user_id INTEGER,
			PRIMARY KEY (teacher_id, user_id),
			FOREIGN KEY (teacher_id) REFERENCES teachers(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);`,
	}

	for _, query := range tableCreationQueries {
		_, err = db.Exec(query)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	// Uncomment and modify the default test types insertion as needed
	// Example of inserting default test types using transaction

	// tx, err := db.Begin()
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: transaction start error: %w", op, err)
	// }
	// defer tx.Rollback() // Safe rollback in case of error

	// stmt, err := tx.Prepare("INSERT INTO test_type (name) VALUES (?)")
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: prepare insert error: %w", op, err)
	// }
	// defer stmt.Close()

	// defaultTests := []string{"Reading", "Listening", "Writing", "Speaking"}
	// for _, category := range defaultTests {
	// 	_, err = stmt.Exec(category)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("%s: insert default test type error: %w", op, err)
	// 	}
	// }

	// if err = tx.Commit(); err != nil {
	// 	return nil, fmt.Errorf("%s: transaction commit error: %w", op, err)
	// }

	return &Sqlite{db: db}, nil
}
