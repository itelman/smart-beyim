package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"smart-beyim/models"
	"strings"

	"github.com/sashabaranov/go-openai"
)

const MaxTokenLimit = 15706
const startPromt = "you are like my teacher in IELTS tell my where my problem"

func (s *Service) GetMessages(userID int) ([]openai.ChatCompletionMessage, error) {
	req, err := s.repo.GetRequestByUserID(userID)
	if err != nil {
		return nil, err
	}

	return req.Messages, nil
}

func (s *Service) StartMessage(userID int) error {
	// TODO: get all test results to promt
	req := s.ai.Start(startPromt)

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
		if errors.Is(err, models.ErrNoRecord) {
			if err = s.StartMessage(userID); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	// Check if the request exceeds the token limit and adjust if necessary
	if tokenCount(req.Messages)+len(content) > MaxTokenLimit { // MaxTokenLimit is a constant, e.g., 16384 for GPT-3
		req.Messages = trimMessages(req.Messages, len(content))
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

// trimMessages removes messages from the beginning until the total token count + new content is under the limit
func trimMessages(messages []openai.ChatCompletionMessage, newContentLen int) []openai.ChatCompletionMessage {
	for len(messages) > 0 {
		if tokenCount(messages)+newContentLen <= MaxTokenLimit {
			break
		}
		messages = messages[1:] // remove the oldest message
	}
	return messages
}

// tokenCount estimates the number of tokens in the messages
func tokenCount(messages []openai.ChatCompletionMessage) int {
	count := 0
	for _, msg := range messages {
		count += len(msg.Content) // Simplified token counting
	}
	return count
}

func (s *Service) GetReview(userID int) (string, error) {
	ieltsList, err := s.GetIelts(userID)
	if err != nil {
		return "", err
	}
	return formatIeltsResults(ieltsList), nil
}

func formatIeltsResults(ieltsList []models.Ielts) string {
	var resultBuilder strings.Builder

	for _, ielts := range ieltsList {
		resultBuilder.WriteString(fmt.Sprintf("IELTS ID: %d, User ID: %d, Time Passed: %s\n", ielts.ID, ielts.UserID, ielts.TimePassed))
		resultBuilder.WriteString(fmt.Sprintf("Reading: Score %f, Feedback: %s, Time Spent: %s, Time Passed: %s\n",
			ielts.Reading.Score, ielts.Reading.Feedback, ielts.Reading.TimeSpent, ielts.Reading.TimePassed))
		resultBuilder.WriteString(fmt.Sprintf("Listening: Score %f, Feedback: %s, Time Spent: %s, Time Passed: %s\n",
			ielts.Listening.Score, ielts.Listening.Feedback, ielts.Listening.TimeSpent, ielts.Listening.TimePassed))
		resultBuilder.WriteString(fmt.Sprintf("Writing: Score %f, Feedback: %s, Time Spent: %s, Time Passed: %s\n",
			ielts.Writing.Score, ielts.Writing.Feedback, ielts.Writing.TimeSpent, ielts.Writing.TimePassed))
		resultBuilder.WriteString(fmt.Sprintf("Speaking: Score %f, Feedback: %s, Time Spent: %s, Time Passed: %s\n",
			ielts.Speaking.Score, ielts.Speaking.Feedback, ielts.Speaking.TimeSpent, ielts.Speaking.TimePassed))
		resultBuilder.WriteString("\n") // Separate entries with a newline for clarity
	}

	return resultBuilder.String()
}
