package service

import (
	"cbt/internal/dtos"
	"cbt/internal/feature"
	"cbt/internal/models"
	"cbt/internal/repository"
	"errors"
	"fmt"
	"time"
)

var (
	ErrExamNotFound   = errors.New("exam tidak ditemukan")
	ErrInvalidClassID = errors.New("class id is invalid or does not exist")
)

type ExamService interface {
	Insert(req *dtos.ExamRequest, idTeacher string) (*models.Exam, error)
	FindByID(id string) (*models.Exam, error)
	FindByTeacherID(id string) ([]*models.Exam, error)
	Delete(id string) (bool, error)
	Update(id string, exam_title string, instructions string, class_id string, subject_id string, duration_minutes int) (bool, error)
}

type examService struct {
	examRepository    repository.ExamRepository
	teacherRepository repository.TeacherRepositoryInterface
}

func NewExamService(examRepository repository.ExamRepository, teacherRepository repository.TeacherRepositoryInterface) *examService {
	return &examService{examRepository, teacherRepository}
}

func (e *examService) FindByID(id string) (*models.Exam, error) {
	exam, err := e.examRepository.GetExamByID(id)
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (e *examService) FindByTeacherID(id string) ([]*models.Exam, error) {
	exam, err := e.examRepository.GetExamByTeacherID(id)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil ujian dari repository: %w", err)
	}
	return exam, nil
}

func (e *examService) Insert(req *dtos.ExamRequest, idTeacher string) (*models.Exam, error) {
	tokenExam, err := feature.GenerateSecureTokenExam()
	t := time.Now().Add(30 * time.Minute)
	var future *time.Time = &t
	if err != nil {
		return nil, err
	}
	newExamStruct := &models.Exam{
		ExamTitle:             req.ExamTitle,
		ClassID:               req.ClassID,
		CreatedByTeacherID:    idTeacher,
		Instructions:          req.Instructions,
		SubjectID:             req.SubjectID,
		StartDatetime:         time.Unix(req.StartDatetime, 0),
		EndDatetime:           time.Unix(req.EndDatetime, 0),
		AccessToken:           tokenExam,
		AccessTokenValidUntil: future,
		DurationMinutes:       req.DurationMinutes,
	}
	newExam, err := e.examRepository.CreateExam(newExamStruct)
	if err != nil {
		return nil, err
	}
	return newExam, nil
}

func (e *examService) Delete(id string) (bool, error) {
	deleteExam, err := e.examRepository.DeleteExam(id)
	if err != nil {
		return deleteExam, err
	}
	return deleteExam, err
}

func (e *examService) Update(id string, exam_title string, instructions string, class_id string, subject_id string, duration_minutes int) (bool, error) {
	_, err := e.examRepository.GetExamByID(id)
	if err != nil {
		return false, ErrExamNotFound
	}
	updateExam, err := e.examRepository.Update(id, exam_title, instructions, class_id, subject_id, duration_minutes)
	if err != nil {
		return updateExam, err
	}
	return updateExam, err
}
