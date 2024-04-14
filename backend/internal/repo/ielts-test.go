package repo

import (
	"database/sql"
	"fmt"
	"smart-beyim/models"
	"time"
)

func (s *Sqlite) GetIeltsByUserID(userID int) ([]models.Ielts, error) {
	const op = "storage.sqlite.GetIeltsByUserID"

	query := `
		SELECT 
			i.id,
			i.time_passed,
			i.user_id,
			r.id AS reading_id, r.user_id AS reading_user_id, r.test_type_id AS reading_test_type,
			r.score AS reading_score, r.feedback AS reading_feedback, r.time_spent AS reading_time_spent, r.time_passed AS reading_time_passed,
			l.id AS listening_id, l.user_id AS listening_user_id, l.test_type_id AS listening_test_type,
			l.score AS listening_score, l.feedback AS listening_feedback, l.time_spent AS listening_time_spent, l.time_passed AS listening_time_passed,
			w.id AS writing_id, w.user_id AS writing_user_id, w.test_type_id AS writing_test_type,
			w.score AS writing_score, w.feedback AS writing_feedback, w.time_spent AS writing_time_spent, w.time_passed AS writing_time_passed,
			s.id AS speaking_id, s.user_id AS speaking_user_id, s.test_type_id AS speaking_test_type,
			s.score AS speaking_score, s.feedback AS speaking_feedback, s.time_spent AS speaking_time_spent, s.time_passed AS speaking_time_passed
		FROM 
			ielts i
		LEFT JOIN 
			tests r ON i.reading_id = r.id
		LEFT JOIN 
			tests l ON i.listening_id = l.id
		LEFT JOIN 
			tests w ON i.writing_id = w.id
		LEFT JOIN 
			tests s ON i.speaking_id = s.id
		WHERE 
			i.user_id = ?;
	`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var ieltsList []models.Ielts

	for rows.Next() {
		var ielts models.Ielts
		var reading, listening, writing, speaking models.Test

		err := rows.Scan(
			&ielts.ID,
			&ielts.TimePassed,
			&ielts.UserID,
			&reading.ID, &reading.UserID, &reading.TestType, &reading.Score, &reading.Feedback, &reading.TimeSpent, &reading.TimePassed,
			&listening.ID, &listening.UserID, &listening.TestType, &listening.Score, &listening.Feedback, &listening.TimeSpent, &listening.TimePassed,
			&writing.ID, &writing.UserID, &writing.TestType, &writing.Score, &writing.Feedback, &writing.TimeSpent, &writing.TimePassed,
			&speaking.ID, &speaking.UserID, &speaking.TestType, &speaking.Score, &speaking.Feedback, &speaking.TimeSpent, &speaking.TimePassed,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		ielts.Reading = reading
		ielts.Listening = listening
		ielts.Writing = writing
		ielts.Speaking = speaking

		ieltsList = append(ieltsList, ielts)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return ieltsList, nil
}

func (s *Sqlite) GetUserIELTSAverages(userID int) (avgScore float32, totalTime time.Time, totalTests int, err error) {
	const op = "storage.sqlite.GetUserIELTSAverages"

	query := `
		SELECT 
			AVG(t.score) AS avg_score,
			SUM(COALESCE(STRFTIME('%s', t.time_spent), 0)) AS total_time,
			COUNT(i.id) AS total_tests
		FROM 
			ielts i
		JOIN 
			tests t ON i.reading_id = t.id OR i.listening_id = t.id OR i.writing_id = t.id OR i.speaking_id = t.id
		WHERE 
			i.user_id = ?;
	`

	row := s.db.QueryRow(query, userID)

	err = row.Scan(&avgScore, &totalTime, &totalTests)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		return
	}

	return avgScore, totalTime, totalTests, nil
}

func InsertTest(db *sql.DB, test models.Test) error {
	// Execute the INSERT statement
	_, err := db.Exec("INSERT INTO tests (id, user_id, test_type_id, score, feedback, time_spent, time_passed) VALUES (?, ?, ?, ?, ?, ?, ?)",
		test.ID, test.UserID, test.TestType, test.Score, test.Feedback, test.TimeSpent, test.TimePassed)
	if err != nil {
		return err
	}
	return nil
}
