package repo

import (
	"fmt"
	"math/rand"
	"time"
)

// func (s *Sqlite) SeedDatabase() error {
// 	const op = "storage.sqlite.SeedDatabase"

// 	users := []struct {
// 		Name     string
// 		Password string
// 	}{
// 		{"Alice", "password123"},
// 		{"Bob", "password123"},
// 		{"Charlie", "password123"},
// 		{"Diana", "password123"},
// 	}

// 	tx, err := s.db.Begin()
// 	if err != nil {
// 		return fmt.Errorf("%s: transaction start error: %w", op, err)
// 	}
// 	defer tx.Rollback() // Safe rollback in case of error

// 	// Insert users
// 	userStmt, err := tx.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
// 	if err != nil {
// 		return fmt.Errorf("%s: prepare user insert error: %w", op, err)
// 	}
// 	defer userStmt.Close()

// 	for _, user := range users {
// 		_, err = userStmt.Exec(user.Name, user.Password)
// 		if err != nil {
// 			return fmt.Errorf("%s: insert user error: %w", op, err)
// 		}
// 	}

// 	// Prepare feedback mapping for each test type
// 	feedbackMapping := map[int][]string{
// 		1: []string{
// 			"Outstanding ability to understand detailed reasoning and complex language.",
// 			"Excellent retention of information and attention to detail even in challenging listening scenarios.",
// 			"Frequent confusion with numbers and proper nouns, which impacts the overall score.",
// 			"Struggles with multi-speaker dialogues and varied accents, needs more focused practice.",
// 		},
// 		2: []string{
// 			"Impressive speed and accuracy in identifying key details and inferring meanings.",
// 			"Demonstrates a strong ability to synthesize information from complex texts.",
// 			"Reading pace needs improvement; misses important details which affects answering precision.",
// 			"Difficulty in handling higher-level vocabulary and implicit questions.",
// 		},
// 		3: []string{
// 			"Articulates complex ideas clearly and effectively, with excellent grammatical control.",
// 			"Writes cohesively, with a logical flow and varied sentence structures.",
// 			"Writing lacks clarity due to poor sentence structure and frequent grammatical errors.",
// 			"Needs to develop ideas more fully, tends to repeat the same points without elaboration.",
// 		},
// 		4: []string{
// 			"Excellent expressive range, using a wide vocabulary and fluent speech without hesitation.",
// 			"Engages well with the examiner with a natural flow and good pronunciation.",
// 			"Shows potential but often pauses and uses fillers, affecting the fluency of speech.",
// 			"Limited vocabulary makes it difficult to fully express complex ideas.",
// 		},
// 	}

// 	// Insert tests and feedback
// 	testStmt, err := tx.Prepare("INSERT INTO tests (user_id, test_type_id, score, feedback, time_spent, time_passed) VALUES (?, ?, ?, ?, ?, ?)")
// 	if err != nil {
// 		return fmt.Errorf("%s: prepare test insert error: %w", op, err)
// 	}
// 	defer testStmt.Close()

// 	for _, user := range users {
// 		for i := 0; i < 3; i++ {
// 			for testType, feedbacks := range feedbackMapping {
// 				score := 5.0 + rand.Float64()*(9.0-5.0)
// 				feedback := feedbacks[rand.Intn(len(feedbacks))]
// 				timeSpent := time.Duration(rand.Intn(30)+60) * time.Minute
// 				timePassed := time.Now().Add(-time.Duration(rand.Intn(30*24*60)) * time.Minute)
// 				_, err = testStmt.Exec(user.Name, testType, score, feedback, timeSpent, timePassed)
// 				if err != nil {
// 					return fmt.Errorf("%s: insert test error: %w", op, err)
// 				}
// 			}
// 		}
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("%s: transaction commit error: %w", op, err)
// 	}

// 	return nil
// }

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

	// Prepare feedback mapping and test statements
	feedbackMapping := map[int][]string{
		1: {"Outstanding ability to understand detailed reasoning.", "Excellent retention of information.", "Frequent confusion with numbers.", "Struggles with varied accents."},
		2: {"Impressive speed in identifying key details.", "Strong ability to synthesize complex texts.", "Reading pace needs improvement.", "Difficulty handling higher-level vocabulary."},
		3: {"Articulates complex ideas clearly.", "Writes cohesively with varied sentence structures.", "Poor sentence structure and errors.", "Needs to develop ideas more fully."},
		4: {"Excellent expressive range and fluent speech.", "Engages well with natural flow.", "Often pauses and uses fillers.", "Limited vocabulary for complex ideas."},
	}
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

	// Insert tests and link them in ielts table
	for _, userID := range userIDs {
		for i := 0; i < 3; i++ { // Each user gets 3 sets of IELTS tests
			testIDs := make(map[int]int64) // Test type ID to test ID mapping
			for testType, feedbacks := range feedbackMapping {
				score := 5.0 + rand.Float64()*(9.0-5.0)
				feedback := feedbacks[rand.Intn(len(feedbacks))]
				timeSpent := time.Duration(rand.Intn(30)+60) * time.Minute
				timePassed := time.Now().Add(-time.Duration(rand.Intn(30*24*60)) * time.Minute)
				res, err := testStmt.Exec(userID, testType, score, feedback, timeSpent, timePassed)
				if err != nil {
					return fmt.Errorf("%s: insert test error: %w", op, err)
				}
				testId, err := res.LastInsertId()
				if err != nil {
					return fmt.Errorf("%s: getting last insert ID for test error: %w", op, err)
				}
				testIDs[testType] = testId
			}
			// Insert a record into ielts table using the captured test IDs
			_, err = ieltsStmt.Exec(userID, testIDs[1], testIDs[2], testIDs[3], testIDs[4])
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
