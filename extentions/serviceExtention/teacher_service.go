package serviceextention

import (
	"cbt/extentions/dtos"
	"cbt/extentions/models"
	repositoryextention "cbt/extentions/repositoryExtention"

	"github.com/google/uuid"
)

type TeacherServiceInterface interface {
	Insert(req *dtos.TeacherRequest) (*models.Teacher, *models.HasRole, error)
	FindAll() ([]models.Teacher, error)
	FindByID(id string) (*models.Teacher, error)
	Delete(id string) (int64, error)
}

type teacherService struct {
	teacherRepo repositoryextention.TeacherRepositoryInterface
	hasRoleRepo repositoryextention.HasRoleRepositoryInterface
}

func NewTeacherService(teacherRepo repositoryextention.TeacherRepositoryInterface, hasRoleRepo repositoryextention.HasRoleRepositoryInterface) *teacherService {
	return &teacherService{teacherRepo, hasRoleRepo}
}

func (s *teacherService) Insert(req *dtos.TeacherRequest) (*models.Teacher, *models.HasRole, error) {
	newTeacher := &models.Teacher{
		NIK:      req.NIK,
		Profile:  req.Profile,
		Password: req.Password,
	}
	teacher, err := s.teacherRepo.Create(newTeacher)
	if err != nil {
		return nil, nil, err
	}
	setHasRole := &models.HasRole{
		ID:        uuid.New(),
		RoleID:    string(repositoryextention.TeacherRole),
		OwnerID:   teacher.ID,
		OwnerType: "teacher",
	}

	hasRoleTeacher, errHasRole := s.hasRoleRepo.Create(setHasRole)
	if errHasRole != nil {
		return nil, nil, errHasRole
	}

	return teacher, hasRoleTeacher, nil
}

func (s *teacherService) FindAll() ([]models.Teacher, error) {
	teacher, err := s.teacherRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (s *teacherService) FindByID(id string) (*models.Teacher, error) {
	teacher, err := s.teacherRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (s *teacherService) Delete(id string) (int64, error) {
	deleteTeacher, err := s.teacherRepo.Delete(id)
	if err != nil {
		return deleteTeacher, err
	}
	return deleteTeacher, nil
}
