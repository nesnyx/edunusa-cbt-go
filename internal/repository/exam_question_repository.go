package repository

import "gorm.io/gorm"

type ExamQuestionRepository interface {
}

type examQuestionRepo struct {
	db *gorm.DB
}

func NewExamQuestionRepository(db *gorm.DB) *examQuestionRepo {
	return &examQuestionRepo{db}
}
