package service

import (
	"cbt/internal/models"
	"cbt/internal/repository"
)

type ExamQuestionServiceInterface interface {
	CreateExamQuestion(questionId string, examId string, displayOrder int) (*models.ExamQuestion, error)
	GetByExam(exam string) (*models.ExamQuestion, error)
	GetByQuestion(id string) (*models.ExamQuestion, error)
	Delete(id string) (bool, error)
}

type examQuestionService struct {
	examQuestionRepo repository.ExamQuestionRepository
}

func NewExamQuestionService(examQuestionRepo repository.ExamQuestionRepository) *examQuestionService {
	return &examQuestionService{examQuestionRepo: examQuestionRepo}
}

func (s *examQuestionService) CreateExamQuestion(questionId string, examId string, displayOrder int) (*models.ExamQuestion, error) {
	examQuestion := &models.ExamQuestion{
		ExamID:       examId,
		QuestionID:   questionId,
		DisplayOrder: displayOrder,
	}
	return s.examQuestionRepo.CreateExamQuestion(examQuestion)
}

func (s *examQuestionService) GetByExam(exam string) (*models.ExamQuestion, error) {
	return s.examQuestionRepo.GetByExam(exam)
}

func (s *examQuestionService) Delete(id string) (bool, error) {
	exam, err := s.examQuestionRepo.DeleteExamQuestion(id)
	if err != nil {
		return exam, err
	}
	return exam, nil
}

func (s *examQuestionService) GetByQuestion(id string) (*models.ExamQuestion, error) {
	exam, err := s.examQuestionRepo.GetByQuestion(id)
	if err != nil {
		return exam, err
	}
	return exam, nil
}
