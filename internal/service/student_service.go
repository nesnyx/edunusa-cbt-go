package service

import (
	"cbt/internal/dtos"
	"cbt/internal/models"
	"cbt/internal/repository"
	"fmt"

	"github.com/google/uuid"
)

type StudentServiceInterface interface {
	InsertStudent(req *dtos.InsertStudentRequest) (*models.Student, *models.HasRole, error)
	FindAll() ([]models.Student, error)
	FindByID(id string) (*models.Student, error)
	FindByNIS(nis string) (*models.Student, error)
}

type service struct {
	studentRepository repository.StudentRepositoryInterface
	hasRoleRepo       repository.HasRoleRepositoryInterface
}

func NewStudentService(studentRepo repository.StudentRepositoryInterface, hasRoleRepo repository.HasRoleRepositoryInterface) *service {
	return &service{studentRepo, hasRoleRepo}
}

func (s *service) FindByID(id string) (*models.Student, error) {
	student, err := s.studentRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *service) InsertStudent(req *dtos.InsertStudentRequest) (*models.Student, *models.HasRole, error) {

	newStudent := &models.Student{
		NIS:      req.NIS,
		Password: req.Password,
		Profile:  req.Profile,
	}
	student, err := s.studentRepository.CreateStudent(newStudent)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create student: %w", err)
	}
	setHasRole := &models.HasRole{
		ID:        uuid.New(),
		RoleID:    string(repository.StudentRole),
		OwnerID:   student.ID,
		OwnerType: "student",
	}
	hasRoleTeacher, errHasRole := s.hasRoleRepo.Create(setHasRole)
	if errHasRole != nil {
		return nil, nil, errHasRole
	}
	return student, hasRoleTeacher, nil
}

func (s *service) FindAll() ([]models.Student, error) {
	students, err := s.studentRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *service) FindByNIS(nis string) (*models.Student, error) {
	student, err := s.studentRepository.FindByNIS(nis)
	if err != nil {
		return nil, err
	}
	return student, nil
}
