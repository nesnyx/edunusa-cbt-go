package repository

import "cbt/internal/models"

type QuestionRepositoryInterface interface {
	CreateQuestion(question *models.Question) (*models.Question, error)
}
