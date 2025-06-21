package repositoryextention

import (
	"cbt/extentions/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentRepositoryInterface interface {
	CreateStudent(student *models.Student) error
	DeleteStudent(id uuid.UUID) (bool, error)
	GetAll() ([]models.Student, error)
}

type student struct {
	db *gorm.DB
}

func NewStudentRepository(database *gorm.DB) *student { // Atau kembalikan *student jika tidak pakai interface
	return &student{database}
}

// DeleteStudent implements StudentRepositoryInterface.
func (s *student) DeleteStudent(id uuid.UUID) (bool, error) {
	panic("unimplemented")
}

func (s *student) CreateStudent(student *models.Student) error {
	return s.db.Create(&student).Error
}

func (s *student) GetAll() ([]models.Student, error) {
	var students []models.Student
	err := s.db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
