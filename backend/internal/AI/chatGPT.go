package ai

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	client *openai.Client
}

func New(key string) *ChatGPT {
	c := openai.NewClient(key)
	return &ChatGPT{
		client: c,
	}
}

func (ai *ChatGPT) Start(startPromt string) *openai.ChatCompletionRequest {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: startPromt,
			},
		},
		MaxTokens: 1024,
	}

	return &req
}

func (ai *ChatGPT) Message(content string, req *openai.ChatCompletionRequest) (string, error) {

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
	resp, err := ai.client.CreateChatCompletion(context.Background(), *req)

	if err != nil {
		return "", err
	}

	req.Messages = append(req.Messages, resp.Choices[0].Message)

	return resp.Choices[0].Message.Content, nil
}
