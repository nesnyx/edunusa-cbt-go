package service

import (
	"cbt/internal/dtos"
	"cbt/internal/models"
	"cbt/internal/repository"
)

type QuestionBankServiceInterface interface {
	CreateQuestionBank(req *dtos.QuestionBankRequest, teacher string) (*models.QuestionBank, error)
	GetQuestionBankByTeacher(teacher string) (*models.QuestionBank, error)
	DeleteQuestionBank(id string) (bool, error)
	UpdateQuestionBank(req *dtos.QuestionBankRequest, teacher string) (bool, error)
}

type questionBankService struct {
	questionBankRepository repository.QuestionBankRepositoryInterface
}

func NewQuestionBankService(questionBankRepository repository.QuestionBankRepositoryInterface) *questionBankService {
	return &questionBankService{questionBankRepository}
}

func (s *questionBankService) CreateQuestionBank(req *dtos.QuestionBankRequest, teacher string) (*models.QuestionBank, error) {
	newQuestionBank := &models.QuestionBank{
		BankName:           req.BankName,
		Description:        req.Description,
		CreatedByTeacherID: teacher,
		SubjectID:          req.Subject,
	}
	questionBank, err := s.questionBankRepository.Create(newQuestionBank)
	if err != nil {
		return nil, err
	}
	return questionBank, nil
}

func (s *questionBankService) DeleteQuestionBank(id string) (bool, error) {
	return s.questionBankRepository.Delete(id)
}

func (s *questionBankService) GetQuestionBankByTeacher(teacher string) (*models.QuestionBank, error) {
	return s.questionBankRepository.GetByTeacher(teacher)
}

func (s *questionBankService) UpdateQuestionBank(req *dtos.QuestionBankRequest, teacher string) (bool, error) {
	newQuestionBank := &models.QuestionBank{
		BankName:           req.BankName,
		Description:        req.Description,
		CreatedByTeacherID: teacher,
		SubjectID:          req.Subject,
	}
	questionBank, err := s.questionBankRepository.Update(newQuestionBank)
	if err != nil {
		return false, err
	}
	return questionBank, nil
}
