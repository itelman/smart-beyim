package service

import (
	AI "smart-beyim/internal/AI"
	"smart-beyim/internal/repo"
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
