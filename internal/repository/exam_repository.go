package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type ExamRepository interface {
	CreateExam(exam *models.Exam) (*models.Exam, error)
	DeleteExam(id string) (bool, error)
	GetExamByID(id string) (*models.Exam, error)
	GetExamByTeacherID(id string) ([]*models.Exam, error)
	Update(id string, exam_title string, instructions string, class_id string, subject_id string, duration_minutes int) (bool, error)
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
	err := e.db.
		Preload("Teacher").
		Preload("Subject").
		Preload("Class").
		Preload("ExamQuestions").
		Preload("ExamQuestions.Question").
		Preload("StudentAttempts").
		Preload("StudentAttempts.Student").
		First(&exam, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (e *examRepository) GetExamByTeacherID(id string) ([]*models.Exam, error) {
	var exam []*models.Exam
	err := e.db.
		Preload("Teacher").                // Load teacher yang membuat exam
		Preload("Subject").                // Load subject exam
		Preload("Class").                  // Load class exam
		Preload("ExamQuestions").          // Load semua soal di exam
		Preload("ExamQuestions.Question"). // Load detail question
		Find(&exam, "created_by_teacher_id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (e *examRepository) DeleteExam(id string) (bool, error) {
	var exam *models.Exam
	err := e.db.Where("id = ?", id).Delete(&exam).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *examRepository) Update(id string, exam_title string, instructions string, class_id string, subject_id string, duration_minutes int) (bool, error) {
	query := "UPDATE exam_engine.exam SET exam_title = ?, instructions = ?, class_id = ? ,subject_id =? , duration_minutes = ? WHERE id = ?"
	err := e.db.Exec(query, exam_title, instructions, class_id, subject_id, duration_minutes, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
