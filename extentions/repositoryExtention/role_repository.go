package repositoryextention

import (
	"cbt/extentions/models"

	"gorm.io/gorm"
)

type RoleRepositoryInterface interface {
	Create(role *models.Role) (*models.Role, error)
	GetAll() ([]models.Role, error)
	GetByID(id string) (*models.Role, error)
	GetByRoleName(name string) (*models.Role, error)
	Update(role *models.Role) (*models.Role, error)
	Delete(id string) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *models.Role) (*models.Role, error) {
	err := r.db.Create(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) GetAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) GetByID(id string) (*models.Role, error) {
	var role *models.Role
	err := r.db.First(&role, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) Update(role *models.Role) (*models.Role, error) {
	err := r.db.Save(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) Delete(id string) error {
	err := r.db.Delete(&models.Role{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) GetByRoleName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("role_name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
