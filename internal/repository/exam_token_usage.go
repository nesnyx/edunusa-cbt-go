package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type ExamTokenUsageRepositoryInterface interface {
	Create(tokenUsage *models.ExamTokenUsage) (*models.ExamTokenUsage, error)
	Delete(id string) (bool, error)
	GetByStudent(id string) (*models.ExamTokenUsage, error)
}

type examTokenRepo struct {
	db *gorm.DB
}

func NewExamTokenUsageRepository(db *gorm.DB) *examTokenRepo {
	return &examTokenRepo{db}
}

func (r *examTokenRepo) Create(tokenUsage *models.ExamTokenUsage) (*models.ExamTokenUsage, error) {
	err := r.db.Create(&tokenUsage).Error
	if err != nil {
		return nil, err
	}
	return tokenUsage, nil
}

func (r *examTokenRepo) Delete(id string) (bool, error) {
	var tokenUsage *models.ExamTokenUsage
	err := r.db.Delete(&tokenUsage, "id = ?", id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *examTokenRepo) GetByStudent(id string) (*models.ExamTokenUsage, error) {
	var tokenUsage *models.ExamTokenUsage
	err := r.db.First(&tokenUsage, "student_id = ?", id).Error
	if err != nil {
		return tokenUsage, err
	}
	return tokenUsage, nil
}
