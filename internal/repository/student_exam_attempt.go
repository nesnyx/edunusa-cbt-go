package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type StudentExamAttemptRepositoryInterface interface {
	Create(studentExamAttempt *models.StudentExamAttempt) (*models.StudentExamAttempt, error)
}

type studentExamAttemptRepo struct {
	db *gorm.DB
}

func NewStudentExamAttemptRepository(db *gorm.DB) *studentExamAttemptRepo {
	return &studentExamAttemptRepo{db}
}

func (r *studentExamAttemptRepo) Create(studentExamAttempt *models.StudentExamAttempt) (*models.StudentExamAttempt, error) {
	err := r.db.Create(&studentExamAttempt).Error
	if err != nil {
		return nil, err
	}
	return studentExamAttempt, nil
}
