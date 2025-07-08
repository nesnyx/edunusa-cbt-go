package repository

import (
	"cbt/internal/models"

	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	LoginStudent(nis string) (*models.Student, error)
	LoginTeacher(nik string) (*models.Teacher, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) LoginStudent(nis string) (*models.Student, error) {
	var student *models.Student
	err := r.db.First(&student, "nis = ?", nis).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (r *authRepository) LoginTeacher(nik string) (*models.Teacher, error) {
	var teacher *models.Teacher
	err := r.db.First(&teacher, "nik = ?", nik).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}
