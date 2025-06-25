package serviceextention

import (
	"cbt/extentions/dtos"
	"cbt/extentions/models"
	repositoryextention "cbt/extentions/repositoryExtention"
)

type SubjectServiceInterface interface {
	Insert(req *dtos.SubjectRequest) (*models.Subject, error)
	FindAll() ([]models.Subject, error)
	FindByID(id string) (*models.Subject, error)
}

type subjectService struct {
	subjectRepo  repositoryextention.SubjectRepositoryInterface
	classService ClassServiceInterface
}

func NewSubjectService(subjectRepo repositoryextention.SubjectRepositoryInterface, classService ClassServiceInterface) *subjectService {
	return &subjectService{subjectRepo, classService}
}

func (s *subjectService) FindAll() ([]models.Subject, error) {
	subject, err := s.subjectRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return subject, nil
}

func (s *subjectService) Insert(req *dtos.SubjectRequest) (*models.Subject, error) {
	checking_existing_class, errExistingClass := s.classService.FindByID(req.ClassID.String())
	if errExistingClass != nil {
		return nil, errExistingClass
	}
	subject := &models.Subject{

		SubjectName: req.SubjectName,
		Description: req.Description,
		ClassID:     checking_existing_class.ID,
	}

	newSubject, err := s.subjectRepo.Create(subject)
	if err != nil {
		return nil, err
	}
	return newSubject, nil
}

func (s *subjectService) FindByID(id string) (*models.Subject, error) {
	subjects, err := s.subjectRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}
