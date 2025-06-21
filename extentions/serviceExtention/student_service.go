package serviceextention

import (
	"cbt/extentions/dtos"
	"cbt/extentions/models"
	repositoryextention "cbt/extentions/repositoryExtention"
	"fmt"
)

type StudentServiceInterface interface {
	InsertStudent(req *dtos.InsertStudentRequest) (*models.Student, error)
	FindAll() ([]models.Student, error)
}

type service struct {
	studentRepository repositoryextention.StudentRepositoryInterface
}

func NewStudentService(studentRepo repositoryextention.StudentRepositoryInterface) *service {
	return &service{studentRepo}
}

func (s *service) InsertStudent(req *dtos.InsertStudentRequest) (*models.Student, error) {

	newStudent := &models.Student{
		NIS:      req.NIS,
		Password: req.Password,
		Profile:  req.Profile,
	}
	err := s.studentRepository.CreateStudent(newStudent)
	if err != nil {
		return nil, fmt.Errorf("failed to create student: %w", err)
	}
	return newStudent, nil
}

func (s *service) FindAll() ([]models.Student, error) {
	students, err := s.studentRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return students, nil
}
