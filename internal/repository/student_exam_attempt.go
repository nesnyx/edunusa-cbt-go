package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type StudentExamAttemptRepositoryInterface interface {
	Create(studentExamAttempt *models.StudentExamAttempt) (*models.StudentExamAttempt, error)
	GetByStudentAndExam(studentID, examID string) (*models.StudentExamAttempt, error)
	GetByID(id string) (*models.StudentExamAttempt, error)
	DeleteByStudentAndExam(studentID, examID string) error
	DeleteByStudent(studentID string) error
	Update(id string, attempt *models.StudentExamAttempt) (*models.StudentExamAttempt, error)
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

func (r *studentExamAttemptRepo) GetByStudentAndExam(studentID, examID string) (*models.StudentExamAttempt, error) {
	var attempt models.StudentExamAttempt
	err := r.db.Where("student_id = ? AND exam_id = ?", studentID, examID).First(&attempt).Error
	return &attempt, err
}

func (r *studentExamAttemptRepo) GetByID(id string) (*models.StudentExamAttempt, error) {
	var attempt models.StudentExamAttempt
	err := r.db.Where("id = ?", id).First(&attempt).Error
	return &attempt, err
}

func (r *studentExamAttemptRepo) DeleteByStudentAndExam(studentID, examID string) error {
	return r.db.Where("student_id = ? AND exam_id = ?", studentID, examID).Delete(&models.StudentExamAttempt{}).Error
}

func (r *studentExamAttemptRepo) DeleteByStudent(studentID string) error {
	return r.db.Where("student_id = ? ", studentID).Delete(&models.StudentExamAttempt{}).Error
}
func (r *studentExamAttemptRepo) Update(id string, attempt *models.StudentExamAttempt) (*models.StudentExamAttempt, error) {
	query := "UPDATE exam_engine.student_exam_attempt SET attempt_end_time = ?, submitted_at = ?, status = ? WHERE id = ?"
	err := r.db.Exec(query, attempt.AttemptEndTime, attempt.SubmittedAt, id).Error
	if err != nil {
		return nil, err
	}
	return attempt, nil
}
