package repositoryextention

import (
	"cbt/extentions/models"

	"gorm.io/gorm"
)

type RoleType string

const (
	StudentRole RoleType = "959ab6b5-2a47-4724-935a-fb7156bc5495"
	TeacherRole RoleType = "db12c729-d7a9-43ef-a8ce-bbbec1dba408"
)

type HasRoleRepositoryInterface interface {
	Create(hasRole *models.HasRole) (*models.HasRole, error)
	GetByID(id string) (*models.HasRole, error)
	Delete(id string) (int64, error)
}

type hasRoleRepo struct {
	db *gorm.DB
}

func NewHasRoleRepository(db *gorm.DB) *hasRoleRepo {
	return &hasRoleRepo{db}
}

func (r *hasRoleRepo) Create(hasRole *models.HasRole) (*models.HasRole, error) {
	err := r.db.Create(&hasRole).Error
	if err != nil {
		return nil, err
	}
	return hasRole, nil
}

func (r *hasRoleRepo) GetByID(id string) (*models.HasRole, error) {
	var hasRole *models.HasRole
	err := r.db.First(&hasRole, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return hasRole, nil
}

func (r *hasRoleRepo) Delete(id string) (int64, error) {
	var hasRole *models.HasRole
	result := r.db.Delete(hasRole, id)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
