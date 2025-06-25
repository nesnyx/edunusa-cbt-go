package serviceextention

import (
	"cbt/extentions/dtos"
	"cbt/extentions/models"
	repositoryextention "cbt/extentions/repositoryExtention"
	"errors"
	"strings"

	"gorm.io/gorm"
)

var (
	ErrRoleNotFound      = errors.New("role not found")
	ErrRoleAlreadyExists = errors.New("role with that name already exists")
	ErrRoleNameRequired  = errors.New("role name is required")
)

type RoleServiceInterface interface {
	Create(req *dtos.RoleRequest) (*models.Role, error)
	GetAll() ([]models.Role, error)
	GetByID(id string) (*models.Role, error)
	Update(req *dtos.RoleRequest) (*models.Role, error)
	Delete(id string) error
}

type roleService struct {
	repo repositoryextention.RoleRepositoryInterface
}

func NewRoleService(repo repositoryextention.RoleRepositoryInterface) RoleServiceInterface {
	return &roleService{repo: repo}
}

func (s *roleService) Create(req *dtos.RoleRequest) (*models.Role, error) {
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, ErrRoleNameRequired
	}

	_, err := s.repo.GetByRoleName(req.Name)
	if err == nil {

		return nil, ErrRoleAlreadyExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	newRole := &models.Role{
		RoleName: req.Name,
	}

	return s.repo.Create(newRole)
}

func (s *roleService) GetAll() ([]models.Role, error) {
	return s.repo.GetAll()
}

func (s *roleService) Update(req *dtos.RoleRequest) (*models.Role, error) {
	if req.ID == "" {
		return nil, errors.New("role ID is required for update")
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, ErrRoleNameRequired
	}

	_, err := s.repo.GetByID(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRoleNotFound
		}
		return nil, err
	}

	existing, err := s.repo.GetByRoleName(req.Name)
	if err == nil && existing.ID != req.ID {
		return nil, ErrRoleAlreadyExists
	}
	updateRole := &models.Role{

		RoleName: req.Name,
	}
	return s.repo.Update(updateRole)
}

func (s *roleService) Delete(id string) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}
		return err
	}
	return s.repo.Delete(id)
}

func (s *roleService) GetByID(id string) (*models.Role, error) {
	role, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRoleNotFound
		}
		return nil, err
	}
	return role, nil
}
