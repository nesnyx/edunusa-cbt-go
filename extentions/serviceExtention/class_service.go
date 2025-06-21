package serviceextention

import (
	"cbt/extentions/dtos"
	"cbt/extentions/models"
	repositoryextention "cbt/extentions/repositoryExtention"
)

type ClassServiceInterface interface {
	Insert(req *dtos.ClassRequest) (*models.Class, error)
	FindAll() ([]models.Class, error)
	FindByID(id string) (*models.Class, error)
}

type classService struct {
	classRepo repositoryextention.ClassRepositoryInterface
}

func NewClassService(classRepo repositoryextention.ClassRepositoryInterface) *classService {
	return &classService{classRepo}
}

func (s *classService) Insert(req *dtos.ClassRequest) (*models.Class, error) {
	classSchema := &models.Class{
		ClassName:   req.ClassName,
		Description: req.Description,
		GradeLevel:  req.GradeLevel,
	}
	classNew, err := s.classRepo.Create(classSchema)
	if err != nil {
		return nil, err
	}
	return classNew, nil
}

func (s *classService) FindByID(id string) (*models.Class, error) {
	class, err := s.classRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (s *classService) FindAll() ([]models.Class, error) {
	class, err := s.classRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return class, nil
}
