package service

import (
	"encoding/json"
	AI "smart-beyim/internal/AI"
	"smart-beyim/internal/repo"

	"github.com/sashabaranov/go-openai"
)

type Service struct {
	repo *repo.Sqlite
	ai   *AI.ChatGPT
}

func NewService(repo *repo.Sqlite, ai *AI.ChatGPT) *Service {
	return &Service{
		repo: repo,
		ai:   ai,
	}
}
func (s *Service) StartMessage(userID int) error {
	// TODO: get all test results to promt
	req := s.ai.Start("start messaging")

	reqString, err := serializeRequest(req)
	if err != nil {
		return err
	}
	if err = s.repo.SaveRequest(userID, reqString); err != nil {
		return err
	}
	return nil
}

func serializeRequest(req *openai.ChatCompletionRequest) (string, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (s *Service) SendMessage(userID int, content string) (string, error) {
	req, err := s.repo.GetRequestByUserID(userID)
	if err != nil {
		return "", err
	}
	answer, err := s.ai.Message(content, req)
	if err != nil {
		return "", err
	}
	_, err = s.repo.DeleteRequestsByUserID(userID)
	if err != nil {
		return "", err
	}

	reqString, err := serializeRequest(req)
	if err != nil {
		return "", err
	}
	if err = s.repo.SaveRequest(userID, reqString); err != nil {
		return "", err
	}

	return answer, nil
}
