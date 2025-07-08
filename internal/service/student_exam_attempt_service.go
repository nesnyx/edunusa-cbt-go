package service

import (
	"cbt/internal/models"
	"cbt/internal/repository"
	"time"
)

type StudentExamAttemptServiceInterface interface {
	Insert(studentId string, examId string) (*models.StudentExamAttempt, error)
}

type studentExamAttemptService struct {
	studentExamAttemptRepo repository.StudentExamAttemptRepositoryInterface
}

func NewStudentExamAttemptService(studentExamAttemptRepo repository.StudentExamAttemptRepositoryInterface) *studentExamAttemptService {
	return &studentExamAttemptService{studentExamAttemptRepo}
}

func (s *studentExamAttemptService) Insert(studentId string, examId string) (*models.StudentExamAttempt, error) {
	now := time.Now()
	results := &models.StudentExamAttempt{
		ExamID:           examId,
		StudentID:        studentId,
		AttemptStartTime: &now,
		Status:           models.AttemptStatusOngoing,
	}
	studentExamAttempt, err := s.studentExamAttemptRepo.Create(results)
	if err != nil {
		return nil, err
	}
	return studentExamAttempt, nil
}
