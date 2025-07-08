package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type StudentAnswerRepositoryInterface interface {
	Create(studentAnswer *models.StudentAnswer) (*models.StudentAnswer, error)
	GetByStudent(id string) ([]*models.StudentAnswer, error)
	GetByID(id string) (*models.StudentAnswer, error)
	GetByQuestionAndStudentAttempt(studentId string, examQuestion string) ([]*models.StudentAnswer, error)
	Update(id string, answerData string) (bool, error)
}

type studentAnswerRepo struct {
	db *gorm.DB
}

func NewStudentAnswerRepository(db *gorm.DB) *studentAnswerRepo {
	return &studentAnswerRepo{db}
}

func (r *studentAnswerRepo) Create(studentAnswer *models.StudentAnswer) (*models.StudentAnswer, error) {
	err := r.db.Create(&studentAnswer).Error
	if err != nil {
		return nil, err
	}
	return studentAnswer, nil
}

func (r *studentAnswerRepo) GetByStudent(id string) ([]*models.StudentAnswer, error) {
	var studentAnswer []*models.StudentAnswer
	err := r.db.Find(&studentAnswer, "student_exam_attempt_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return studentAnswer, nil
}

func (r *studentAnswerRepo) GetByQuestionAndStudentAttempt(studentId string, examQuestion string) ([]*models.StudentAnswer, error) {
	var students []*models.StudentAnswer
	err := r.db.Find(&students, "exam_question_id = ? AND student_exam_attempt_id = ?", examQuestion, studentId).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentAnswerRepo) Update(id string, answerData string) (bool, error) {
	query := "UPDATE exam_engine.student_answer SET answer_data = ? WHERE id = ?"
	err := r.db.Exec(query, answerData, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *studentAnswerRepo) GetByID(id string) (*models.StudentAnswer, error) {
	var studentAnswer *models.StudentAnswer
	err := r.db.First(&studentAnswer, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return studentAnswer, nil
}
