package repo

import (
	"fmt"
	"math/rand"
	"time"
)

func (s *Sqlite) SeedDatabase() error {
	const op = "storage.sqlite.SeedDatabase"

	users := []struct {
		Name     string
		Password string
	}{
		{"Alice", "password123"},
		{"Bob", "password123"},
		{"Charlie", "password123"},
		{"Diana", "password123"},
	}

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: transaction start error: %w", op, err)
	}
	defer tx.Rollback() // Safe rollback in case of error

	// Insert users
	userStmt, err := tx.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("%s: prepare user insert error: %w", op, err)
	}
	defer userStmt.Close()

	userIDs := make(map[string]int)
	for _, user := range users {
		res, err := userStmt.Exec(user.Name, user.Password)
		if err != nil {
			return fmt.Errorf("%s: insert user error: %w", op, err)
		}
		userId, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("%s: getting last insert ID for user error: %w", op, err)
		}
		userIDs[user.Name] = int(userId)
	}

	// Insert tests and feedback
	testStmt, err := tx.Prepare("INSERT INTO tests (user_id, test_type_id, score, feedback, time_spent, time_passed) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("%s: prepare test insert error: %w", op, err)
	}
	defer testStmt.Close()

	ieltsStmt, err := tx.Prepare("INSERT INTO ielts (user_id, reading_id, listening_id, writing_id, speaking_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("%s: prepare ielts insert error: %w", op, err)
	}
	defer ieltsStmt.Close()

	testTypes := map[string]int{"Reading": 1, "Listening": 2, "Writing": 3, "Speaking": 4}
	feedbackTypes := []string{"Needs improvement in...", "Excellent performance in..."}

	for _, userID := range userIDs {
		// Each user gets three full IELTS records
		for i := 0; i < 3; i++ {
			testIDs := make(map[string]int64)
			for testType, testTypeID := range testTypes {
				score := 5.0 + rand.Float64()*(9.0-5.0)
				feedback := feedbackTypes[rand.Intn(len(feedbackTypes))]
				timeSpent := time.Duration(rand.Intn(30)+60) * time.Minute
				timePassed := time.Now().Add(-time.Duration(rand.Intn(30*24*60)) * time.Minute)
				res, err := testStmt.Exec(userID, testTypeID, score, feedback, timeSpent, timePassed)
				if err != nil {
					return fmt.Errorf("%s: insert test error: %w", op, err)
				}
				testId, err := res.LastInsertId()
				if err != nil {
					return fmt.Errorf("%s: getting last insert ID for test error: %w", op, err)
				}
				testIDs[testType] = testId
			}
			// Insert into ielts table
			_, err = ieltsStmt.Exec(userID, testIDs["Reading"], testIDs["Listening"], testIDs["Writing"], testIDs["Speaking"])
			if err != nil {
				return fmt.Errorf("%s: insert ielts record error: %w", op, err)
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: transaction commit error: %w", op, err)
	}

	return nil
}
