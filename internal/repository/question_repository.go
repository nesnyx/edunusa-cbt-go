package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type QuestionRepositoryInterface interface {
	GetByTeacher(teacher string) (*models.Question, error)
	CreateQuestion(question *models.Question) (*models.Question, error)
	DeleteQuestion(id string) (bool, error)
	UpdateQuestion(question *models.Question, teacher string, id string) (bool, error)
	// GetByID(id string) (*models.Question, error)
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *questionRepository {
	return &questionRepository{db}
}

func (r *questionRepository) CreateQuestion(question *models.Question) (*models.Question, error) {
	err := r.db.Create(&question).Error
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (r *questionRepository) DeleteQuestion(id string) (bool, error) {
	var question *models.Question
	err := r.db.Delete(&question, id).Error
	if err != nil {
		return false, err
	}
	return true, nil

}

func (r *questionRepository) GetByTeacher(teacher string) (*models.Question, error) {
	var question *models.Question
	err := r.db.Find(&question, "created_by_teacher_id = ?", teacher).Error
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (r *questionRepository) UpdateQuestion(question *models.Question, teacher string, id string) (bool, error) {
	query := "UPDATE question SET question_text = ?, question_type = ?, point = ? , metadata = ? WHERE created_by_teacher_id = ? AND id = ?"
	err := r.db.Exec(query, question.QuestionText, question.QuestionType, question.Points, question.Metadata, teacher, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
