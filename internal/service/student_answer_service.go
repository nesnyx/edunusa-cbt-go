package service

import (
	"cbt/internal/dtos"
	"cbt/internal/models"
	"cbt/internal/repository"
	"errors"
	"log"

	"gorm.io/gorm"
)

type StudentAnswerServiceInterface interface {
	InsertOrUpdate(req *dtos.StudentAnswerRequest, id string) (bool, error)
	GetByStudent(id string) ([]*models.StudentAnswer, error)
	GetByQuestionAndStudentAttempt(studentId string, examQuestion string) ([]*models.StudentAnswer, error)
}

type studentAnswerService struct {
	studentAnswerRepo repository.StudentAnswerRepositoryInterface
}

func NewStudentAnswerService(studentAnswerRepo repository.StudentAnswerRepositoryInterface) *studentAnswerService {
	return &studentAnswerService{studentAnswerRepo: studentAnswerRepo}
}

func (s *studentAnswerService) InsertOrUpdate(req *dtos.StudentAnswerRequest, id string) (bool, error) {
	newAnswer := &models.StudentAnswer{
		AnswerData:           req.AnswerData,
		StudentExamAttemptID: req.StudentExamAttemptId,
		ExamQuestionID:       req.ExamQuestionID,
	}
	_, err := s.studentAnswerRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			log.Println("Jawaban tidak ditemukan, membuat data baru...")
			_, err := s.studentAnswerRepo.Create(newAnswer)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	updateAnswer, _ := s.studentAnswerRepo.Update(req.AnswerData, id)
	return updateAnswer, nil

}

func (s *studentAnswerService) GetByQuestionAndStudentAttempt(studentId string, examQuestion string) ([]*models.StudentAnswer, error) {
	studentAnswer, err := s.studentAnswerRepo.GetByQuestionAndStudentAttempt(studentId, examQuestion)
	if err != nil {
		return nil, err
	}
	return studentAnswer, nil
}

func (s *studentAnswerService) GetByStudent(id string) ([]*models.StudentAnswer, error) {
	studentAnswer, err := s.studentAnswerRepo.GetByStudent(id)
	if err != nil {
		return nil, err
	}
	return studentAnswer, nil
}
