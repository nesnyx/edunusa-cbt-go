package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type TeacherRepositoryInterface interface {
	Create(teacher *models.Teacher) (*models.Teacher, error)
	GetAll() ([]models.Teacher, error)
	GetByID(id string) (*models.Teacher, error)
	Delete(id string) (int64, error)
}
type teacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *teacherRepository {
	return &teacherRepository{db}
}

func (r *teacherRepository) Create(teacher *models.Teacher) (*models.Teacher, error) {
	err := r.db.Create(&teacher).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (r *teacherRepository) GetAll() ([]models.Teacher, error) {
	var teacher []models.Teacher
	err := r.db.Find(&teacher).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (r *teacherRepository) GetByID(id string) (*models.Teacher, error) {
	var teacher *models.Teacher
	err := r.db.First(&teacher, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (r *teacherRepository) Delete(id string) (int64, error) {
	var teacher models.Teacher
	result := r.db.Delete(teacher, id)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
