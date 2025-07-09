package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type ExamTokenUsageRepositoryInterface interface {
	Create(tokenUsage *models.ExamTokenUsage) (*models.ExamTokenUsage, error)
	Delete(id string) (bool, error)
	GetByStudentAndExam(student string, exam string) (*models.ExamTokenUsage, error)
	DeleteByStudentAndExam(studentID, examID string) error
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

func (r *examTokenRepo) GetByStudentAndExam(studentID, examID string) (*models.ExamTokenUsage, error) {
	var tokenUsage models.ExamTokenUsage
	err := r.db.Where("student_id = ? AND exam_id = ?", studentID, examID).First(&tokenUsage).Error
	return &tokenUsage, err
}

func (r *examTokenRepo) DeleteByStudentAndExam(studentID, examID string) error {
	return r.db.Where("student_id = ? AND exam_id = ?", studentID, examID).Delete(&models.ExamTokenUsage{}).Error
}
