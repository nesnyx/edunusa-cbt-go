package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type QuestionBankRepositoryInterface interface {
	GetByTeacher(teacher string) ([]*models.QuestionBank, error)
	GetBySubject(subject string) ([]*models.QuestionBank, error)
	Create(questionBank *models.QuestionBank) (*models.QuestionBank, error)
	Delete(id string) (bool, error)
	Update(bankName string, description string, id string, teacher string) (bool, error)
}

type questionBankRepository struct {
	db *gorm.DB
}

func NewQuestionBankRepository(db *gorm.DB) *questionBankRepository {
	return &questionBankRepository{db}
}

func (r *questionBankRepository) Create(questionBank *models.QuestionBank) (*models.QuestionBank, error) {
	err := r.db.Create(&questionBank).Error
	if err != nil {
		return nil, err
	}
	return questionBank, nil
}

func (r *questionBankRepository) Delete(id string) (bool, error) {
	var questionBank *models.QuestionBank
	err := r.db.Delete(&questionBank, "id = ?", id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *questionBankRepository) GetByTeacher(teacher string) ([]*models.QuestionBank, error) {
	var questionBank []*models.QuestionBank
	err := r.db.
		Preload("Subject").
		Preload("CreatedByTeacher").
		Find(&questionBank, "created_by_teacher_id = ?", teacher).Error
	if err != nil {
		return nil, err
	}
	return questionBank, nil
}

func (r *questionBankRepository) Update(bankName string, description string, id string, teacher string) (bool, error) {
	query := "UPDATE exam_engine.question_bank SET bank_name = ?, description = ? WHERE created_by_teacher_id = ? AND id = ?"
	err := r.db.Exec(query, bankName, description, teacher, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *questionBankRepository) GetBySubject(subject string) ([]*models.QuestionBank, error) {
	var questionBank []*models.QuestionBank
	err := r.db.Find(&questionBank, "subject_id = ?", subject).Error
	if err != nil {
		return nil, err
	}
	return questionBank, nil
}
