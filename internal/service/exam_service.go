package service

import (
	repositoryextention "cbt/extentions/repositoryExtention"
	"cbt/internal/dtos"
	"cbt/internal/feature"
	"cbt/internal/models"
	"cbt/internal/repository"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrExamNotFound   = errors.New("exam tidak ditemukan")
	ErrInvalidClassID = errors.New("class id is invalid or does not exist")
)

type ExamService interface {
	Insert(req *dtos.ExamRequest) (*models.Exam, error)
	FindByID(id string) (*models.Exam, error)
	FindByTeacherID(c *gin.Context) (*models.Exam, error)
	Delete(id string) (int64, error)
	Update(id string, instructions string, class_id string, duration_minutes int) (bool, error)
}

type examService struct {
	examRepository    repository.ExamRepository
	teacherRepository repositoryextention.TeacherRepositoryInterface
}

func NewExamService(examRepository repository.ExamRepository, teacherRepository repositoryextention.TeacherRepositoryInterface) *examService {
	return &examService{examRepository, teacherRepository}
}

func (e *examService) FindByID(id string) (*models.Exam, error) {
	exam, err := e.examRepository.GetExamByID(id)
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func (e *examService) FindByTeacherID(c *gin.Context) (*models.Exam, error) {
	currentUserData, exists := c.Get("currentUser")
	if !exists {
		return nil, errors.New("kesalahan server: informasi pengguna tidak ditemukan dalam konteks")
	}
	idTeacher, ok := currentUserData.(dtos.ExamRequestCreatedByTeacherID)
	if !ok {
		return nil, fmt.Errorf("kesalahan server: tipe data pengguna tidak valid dalam konteks, diharapkan dtos.ExamRequestCreatedByTeacherID, didapatkan %T", currentUserData)
	}
	exam, err := e.examRepository.GetExamByTeacherID(idTeacher.CreatedByTeacherID.String())
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil ujian dari repository: %w", err)
	}
	return exam, nil
}

func (e *examService) Insert(req *dtos.ExamRequest) (*models.Exam, error) {
	_, errTeacherExisting := e.teacherRepository.GetByID(req.CreatedByTeacherID.String())
	if errTeacherExisting != nil {
		if errors.Is(errTeacherExisting, gorm.ErrRecordNotFound) {
			return nil, errors.New("teacher with that ID not found")
		}
		return nil, errTeacherExisting
	}
	tokenExam, err := feature.GenerateSecureTokenExam()
	t := time.Now().Add(30 * time.Minute)
	var future *time.Time = &t
	if err != nil {
		return nil, err
	}
	newExamStruct := &models.Exam{
		Base: models.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ExamTitle:             req.ExamTitle,
		ClassID:               req.ClassID,
		CreatedByTeacherID:    req.CreatedByTeacherID,
		Instructions:          req.Instructions,
		SubjectID:             req.SubjectID,
		StartDatetime:         req.StartDatetime,
		EndDatetime:           req.EndDatetime,
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

func (e *examService) Delete(id string) (int64, error) {

	deleteExam, err := e.examRepository.DeleteExam(id)
	if err != nil {
		return deleteExam, err
	}
	return deleteExam, err
}

func (e *examService) Update(id string, instructions string, class_id string, duration_minutes int) (bool, error) {
	_, err := e.examRepository.GetExamByID(id)
	if err != nil {
		return false, ErrExamNotFound
	}
	updateExam, err := e.examRepository.Update(id, instructions, class_id, duration_minutes)
	if err != nil {
		return updateExam, err
	}
	return updateExam, err
}
