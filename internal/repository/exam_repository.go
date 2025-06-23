package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type ExamRepository interface {
	CreateExam(exam *models.Exam) (*models.Exam, error)
	DeleteExam(id string) (int64, error)
	GetExamByID(id string) (*models.Exam, error)
	GetExamByTeacherID(id string) (*models.Exam, error)
	Update(id string, instructions string, class_id string, duration_minutes int) (bool, error)
}

type examRepository struct {
	db *gorm.DB
}

func NewExamRepository(db *gorm.DB) *examRepository {
	return &examRepository{db: db}
}

func (e *examRepository) CreateExam(exam *models.Exam) (*models.Exam, error) {
	err := e.db.Create(exam).Error
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (e *examRepository) GetExamByID(id string) (*models.Exam, error) {
	var exam *models.Exam
	err := e.db.First(&exam, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (e *examRepository) GetExamByTeacherID(id string) (*models.Exam, error) {
	var exam *models.Exam
	err := e.db.First(&exam, "created_by_teacher_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (e *examRepository) DeleteExam(id string) (int64, error) {
	var exam *models.Exam
	result := e.db.Delete(&exam, id)

	// Tangani jika ada error dari GORM (misal: koneksi putus)
	if result.Error != nil {
		return 0, result.Error
	}

	// Kembalikan jumlah baris yang dihapus dan error (yang akan nil jika sukses)
	return result.RowsAffected, nil
}

func (e *examRepository) Update(id string, instructions string, class_id string, duration_minutes int) (bool, error) {
	query := "UPDATE exam SET exam_title = ?, instructions = ?, class_id = ? , duration_minutes = ? WHERE id = ?"
	err := e.db.Exec(query, instructions, class_id, duration_minutes, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
