package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type ExamQuestionRepository interface {
	CreateExamQuestion(examQuestion *models.ExamQuestion) (*models.ExamQuestion, error)
	GetByQuestion(id string) (*models.ExamQuestion, error)
	GetByExam(exam string) (*models.ExamQuestion, error)
	DeleteExamQuestion(id string) (bool, error)
}

type examQuestionRepo struct {
	db *gorm.DB
}

func NewExamQuestionRepository(db *gorm.DB) *examQuestionRepo {
	return &examQuestionRepo{db}
}

func (s *examQuestionRepo) CreateExamQuestion(examQuestion *models.ExamQuestion) (*models.ExamQuestion, error) {
	err := s.db.Create(&examQuestion).Error
	if err != nil {
		return nil, err
	}
	return examQuestion, nil
}

func (s *examQuestionRepo) GetByExam(exam string) (*models.ExamQuestion, error) {
	var examQuestion *models.ExamQuestion
	err := s.db.First(&examQuestion, "exam_id = ?", exam).Error
	if err != nil {
		return nil, err
	}
	return examQuestion, nil

}

func (s *examQuestionRepo) DeleteExamQuestion(id string) (bool, error) {
	var examQuestion *models.ExamQuestion
	err := s.db.Delete(&examQuestion, "id = ?", examQuestion).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *examQuestionRepo) GetByQuestion(id string) (*models.ExamQuestion, error) {
	var examQuestion *models.ExamQuestion
	err := s.db.
		Preload("Question").
		Preload("Exam").
		First(&examQuestion, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return examQuestion, nil
}
