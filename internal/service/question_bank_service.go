package service

import (
	"cbt/internal/dtos"
	"cbt/internal/models"
	"cbt/internal/repository"
)

type QuestionBankServiceInterface interface {
	CreateQuestionBank(req *dtos.QuestionBankRequest, teacher string) (*models.QuestionBank, error)
	GetBySubject(subjet string) ([]*models.QuestionBank, error)
	GetQuestionBankByTeacher(teacher string) ([]*dtos.QuestionBankResponse, error)
	DeleteQuestionBank(id string) (bool, error)
	UpdateQuestionBank(bankName string, description string, id string, teacher string) (bool, error)
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

func (s *questionBankService) GetQuestionBankByTeacher(teacher string) ([]*dtos.QuestionBankResponse, error) {
	var results []*dtos.QuestionBankResponse
	question, err := s.questionBankRepository.GetByTeacher(teacher)
	if err != nil {
		return nil, err
	}
	for _, q := range question {
		results = append(results, &dtos.QuestionBankResponse{
			ID:          q.ID,
			BankName:    q.BankName,
			Description: q.Description,
			TeacherNik:  q.CreatedByTeacher.NIK,
			Teacher:     q.CreatedByTeacher.Profile,
			SubjectName: q.Subject.SubjectName,
		})
	}
	return results, nil
}

func (s *questionBankService) UpdateQuestionBank(bankName string, description string, id string, teacher string) (bool, error) {
	questionBank, err := s.questionBankRepository.Update(bankName, description, id, teacher)
	if err != nil {
		return false, err
	}
	return questionBank, nil
}

func (s *questionBankService) GetBySubject(subjet string) ([]*models.QuestionBank, error) {
	exam, err := s.questionBankRepository.GetBySubject(subjet)
	if err != nil {
		return nil, err
	}
	return exam, nil
}
