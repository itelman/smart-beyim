package repo

import (
	"smart-beyim/models"
	"time"
)

func (s *Sqlite) GetTestsByUserID(userID int) ([]models.Test, error) {
	var tests []models.Test

	// Example query
	rows, err := s.db.Query("SELECT id, user_id, test_type, score, feedback, time_spent, time_passed FROM tests WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var test models.Test
		err := rows.Scan(&test.ID, &test.UserID, &test.TestType, &test.Score, &test.Feedback, &test.TimeSpent, &test.TimePassed)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}

	return tests, nil
}

func (s *Sqlite) InsertTest(userID int, testType int, score float32, feedback string, timeSpent time.Time) error {
	_, err := s.db.Exec("INSERT INTO tests (user_id, test_type_id, score, feedback, time_spent) VALUES (?, ?, ?, ?, ?)", userID, testType, score, feedback, timeSpent)
	return err
}

func (s *Sqlite) GetAvgScoreByUserIDAndTestType(userID int, testType int) (float32, error) {
	var avgScore float32

	err := s.db.QueryRow("SELECT AVG(score) FROM tests WHERE user_id = ? AND test_type_id = ?", userID, testType).Scan(&avgScore)
	if err != nil {
		return 0, err
	}

	return avgScore, nil
}
