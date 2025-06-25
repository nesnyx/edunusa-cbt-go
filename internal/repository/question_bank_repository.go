package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type QuestionBankRepositoryInterface interface {
	GetByTeacher(teacher string) (*models.QuestionBank, error)
	Create(questionBank *models.QuestionBank) (*models.QuestionBank, error)
	Delete(id string) (bool, error)
	Update(questionBank *models.QuestionBank) (bool, error)
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
	err := r.db.Delete(&questionBank, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *questionBankRepository) GetByTeacher(teacher string) (*models.QuestionBank, error) {
	var questionBank *models.QuestionBank
	err := r.db.Find(&questionBank, "created_by_teacher_id = ?", teacher).Error
	if err != nil {
		return nil, err
	}
	return questionBank, nil
}

func (r *questionBankRepository) Update(questionBank *models.QuestionBank) (bool, error) {
	query := "UPDATE question_bank SET bank_name = ?, description = ? WHERE created_by_teacher_id = ?"
	err := r.db.Exec(query, questionBank.BankName, questionBank.Description, questionBank.CreatedByTeacherID).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
