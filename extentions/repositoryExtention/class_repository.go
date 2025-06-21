package repositoryextention

import (
	"cbt/extentions/models"

	"gorm.io/gorm"
)

type ClassRepositoryInterface interface {
	Create(class *models.Class) (*models.Class, error)
	GetAll() ([]models.Class, error)
	GetByID(id string) (*models.Class, error)
}

type classRepo struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *classRepo {
	return &classRepo{db}
}

func (r *classRepo) Create(class *models.Class) (*models.Class, error) {
	err := r.db.Create(&class).Error
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (r *classRepo) GetAll() ([]models.Class, error) {
	var class []models.Class
	err := r.db.Find(&class).Error
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (r *classRepo) GetByID(id string) (*models.Class, error) {
	var class *models.Class
	err := r.db.First(&class, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return class, nil
}
