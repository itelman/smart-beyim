package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"smart-beyim/models"

	"github.com/sashabaranov/go-openai"
)

func (s *Sqlite) GetRequestByUserID(userID int) (*openai.ChatCompletionRequest, error) {
	const op = "repo.Sqlite.GetSingleRequestByUserID"
	var requestData string

	query := `SELECT request_data FROM chat_requests WHERE user_id = ? ORDER BY id DESC LIMIT 1`
	err := s.db.QueryRow(query, userID).Scan(&requestData)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNoRecord
		}
		return nil, fmt.Errorf("%s: error querying the database: %w", op, err)
	}

	var request openai.ChatCompletionRequest
	if err := json.Unmarshal([]byte(requestData), &request); err != nil {
		return nil, fmt.Errorf("%s: error unmarshalling request data: %w", op, err)
	}

	return &request, nil
}
func (s *Sqlite) SaveRequest(user_id int, requestData string) error {
	const op = "repo.Sqlite.SaveRequest"
	_, err := s.db.Exec("INSERT INTO chat_requests (user_id, request_data) VALUES (?, ?)", user_id, requestData)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Sqlite) DeleteRequestsByUserID(userID int) (int64, error) {
	const op = "repo.Sqlite.DeleteRequestsByUserID"

	query := `DELETE FROM chat_requests WHERE user_id = ?`
	result, err := s.db.Exec(query, userID)
	if err != nil {
		return 0, fmt.Errorf("%s: error executing delete: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("%s: error retrieving affected rows: %w", op, err)
	}

	return rowsAffected, nil
}

func (s *Sqlite) UpdateRequestDataByUserID(userID int, newRequestData string) error {
	const op = "repo.Sqlite.UpdateRequestDataByUserID"

	_, err := s.db.Exec("UPDATE chat_requests SET request_data = ? WHERE user_id = ?", newRequestData, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
