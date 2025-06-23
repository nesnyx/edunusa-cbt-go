package repositoryextention

import (
	"cbt/extentions/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentRepositoryInterface interface {
	CreateStudent(student *models.Student) (*models.Student, error)
	DeleteStudent(id uuid.UUID) (bool, error)
	GetAll() ([]models.Student, error)
	FindByID(id string) (*models.Student, error)
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

func (s *student) CreateStudent(student *models.Student) (*models.Student, error) {
	err := s.db.Create(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *student) GetAll() ([]models.Student, error) {
	var students []models.Student
	err := s.db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *student) FindByID(id string) (*models.Student, error) {
	var student *models.Student
	err := s.db.Select("id,nis,profile").First(&student, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}
