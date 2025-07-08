package service

import (
	"cbt/internal/dtos"
	"cbt/internal/models"
	"cbt/internal/repository"
	"errors"
)

type QuestionServiceInterface interface {
	GetByTeacher(teacher string) ([]*dtos.QuestionResponse, error)
	CreateQuestion(req *dtos.QuestionRequest, questionBank string, teacher string) (*models.Question, error)
	DeleteQuestion(id string) (bool, error)
	UpdateQuestion(req *dtos.QuestionRequest, teacher string, id string) (bool, error)
}

type questionService struct {
	questionRepository repository.QuestionRepositoryInterface
}

func NewQuestionService(questionRepository repository.QuestionRepositoryInterface) *questionService {
	return &questionService{questionRepository}
}

func (s *questionService) CreateQuestion(req *dtos.QuestionRequest, questionBank string, teacher string) (*models.Question, error) {
	newQuestion := &models.Question{
		QuestionText:       req.QuestionText,
		QuestionType:       req.QuestionType,
		QuestionBankID:     questionBank,
		Points:             req.Points,
		CreatedByTeacherID: teacher,
		Metadata:           req.Metadata,
	}
	question, err := s.questionRepository.CreateQuestion(newQuestion)
	if err != nil {
		return nil, errors.New("gagal membuat soal")
	}
	return question, nil
}

func (s *questionService) DeleteQuestion(id string) (bool, error) {
	return s.questionRepository.DeleteQuestion(id)
}

func (s *questionService) GetByTeacher(teacher string) ([]*dtos.QuestionResponse, error) {
	var results []*dtos.QuestionResponse
	question, err := s.questionRepository.GetByTeacher(teacher)
	if err != nil {
		return nil, err
	}
	for _, q := range question {
		results = append(results, &dtos.QuestionResponse{
			ID:           q.ID,
			QuestionText: q.QuestionText,
			QuestionType: string(q.QuestionType),
			Points:       q.Points,
			Metadata:     q.Metadata,
			BankName:     q.QuestionBank.BankName,
			CreatedByNIK: q.CreatedByTeacher.NIK,
		})
	}
	return results, nil
}

func (s *questionService) UpdateQuestion(req *dtos.QuestionRequest, teacher string, id string) (bool, error) {
	newQuestion := &models.Question{
		QuestionText: req.QuestionText,
		QuestionType: req.QuestionType,
		Points:       req.Points,
		Metadata:     req.Metadata,
	}
	question, err := s.questionRepository.UpdateQuestion(newQuestion, teacher, id)
	if err != nil {
		return false, err
	}
	return question, nil
}
