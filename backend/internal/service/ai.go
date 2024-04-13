package service

import (
	"encoding/json"

	"github.com/sashabaranov/go-openai"
)

func (s *Service) GetMessages(userID int) ([]openai.ChatCompletionMessage, error) {
	req, err := s.repo.GetRequestByUserID(userID)
	if err != nil {
		return nil, err
	}

	return req.Messages, nil
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

	reqString, err := serializeRequest(req)
	if err != nil {
		return "", err
	}
	if err = s.repo.UpdateRequestDataByUserID(userID, reqString); err != nil {
		return "", err
	}

	return answer, nil
}
