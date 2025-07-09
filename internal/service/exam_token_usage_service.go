package service

import (
	"cbt/internal/models"
	"cbt/internal/repository"
)

type ExamTokenUsageServiceInterface interface {
	Create(token string, exam string, student string) (*models.ExamTokenUsage, error)
	Delete(id string) (bool, error)
	FindByStudentAndExam(student string, examId string) (*models.ExamTokenUsage, error)
}

type examTokenService struct {
	examTokenUsageRepo repository.ExamTokenUsageRepositoryInterface
}

func NewExamTokenUsage(examTokenUsageRepo repository.ExamTokenUsageRepositoryInterface) *examTokenService {
	return &examTokenService{examTokenUsageRepo: examTokenUsageRepo}
}

func (s *examTokenService) Create(token string, exam string, student string) (*models.ExamTokenUsage, error) {

	tokenUsage := &models.ExamTokenUsage{
		StudentID:      student,
		ExamID:         exam,
		TokenValueUsed: token,
	}
	newTokenUsage, err := s.examTokenUsageRepo.Create(tokenUsage)
	if err != nil {
		return nil, err
	}
	return newTokenUsage, nil

}

func (s *examTokenService) Delete(id string) (bool, error) {
	delete, err := s.examTokenUsageRepo.Delete(id)
	if err != nil {
		return false, err
	}
	return delete, nil
}

func (s *examTokenService) FindByStudentAndExam(student string, examId string) (*models.ExamTokenUsage, error) {
	exam, err := s.examTokenUsageRepo.GetByStudentAndExam(student, examId)
	if err != nil {
		return nil, err
	}
	return exam, nil
}
