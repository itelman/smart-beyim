package service

import (
	"smart-beyim/models"
)

func (s *Service) GetIelts(userID int) ([]models.Ielts, error) {
	return s.repo.GetIeltsByUserID(userID)
}
