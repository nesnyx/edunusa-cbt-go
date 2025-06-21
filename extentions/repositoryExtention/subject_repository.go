package repositoryextention

import (
	"cbt/extentions/models"

	"gorm.io/gorm"
)

type SubjectRepositoryInterface interface {
	Create(subject *models.Subject) (*models.Subject, error)
	GetAll() ([]models.Subject, error)
	GetByID(id string) (*models.Subject, error)
}

type subject struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *subject {
	return &subject{db}
}

func (s *subject) Create(subject *models.Subject) (*models.Subject, error) {
	err := s.db.Create(&subject).Error
	if err != nil {
		return nil, err
	}
	return subject, nil
}

func (s *subject) GetAll() ([]models.Subject, error) {
	var subject []models.Subject
	err := s.db.Find(&subject).Error
	if err != nil {
		return nil, err
	}
	return subject, nil
}
func (s *subject) GetByID(id string) (*models.Subject, error) {
	var subject *models.Subject
	err := s.db.First(&subject, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return subject, nil
}
