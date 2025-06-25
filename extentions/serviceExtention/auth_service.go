package serviceextention

import (
	repositoryextention "cbt/extentions/repositoryExtention"
	"cbt/pkg/utils"
	"errors"
)

var (
	ErrStudentNotFound      = errors.New("student tidak ditemukan")
	ErrPasswordDoesnotMatch = errors.New("password tidak sesuai")
	ErrTeacherNotFound      = errors.New("teacher tidak ditemukan")
)

type AuthServiceInterface interface {
	LoginTeacher(nik string, password string) (string, error)
	LoginStudent(nis string, password string) (string, error)
}

type authService struct {
	authRepo repositoryextention.AuthRepositoryInterface
}

func NewAuthService(authRepo repositoryextention.AuthRepositoryInterface) *authService {
	return &authService{authRepo}
}

func (s *authService) LoginTeacher(nik string, password string) (string, error) {
	getTeacher, err := s.authRepo.LoginTeacher(nik)
	if err != nil {
		return "", ErrTeacherNotFound
	}
	if getTeacher.Password != password {
		return "", ErrPasswordDoesnotMatch
	}
	generatedToken, err := utils.GenerateJWT(getTeacher.Base.ID, "teacher")
	if err != nil {
		return "", err
	}
	return generatedToken, nil

}

func (s *authService) LoginStudent(nis string, password string) (string, error) {
	getStudent, err := s.authRepo.LoginStudent(nis)
	if err != nil {
		return "", ErrStudentNotFound
	}
	if getStudent.Password != password {
		return "", ErrPasswordDoesnotMatch
	}
	generatedToken, err := utils.GenerateJWT(getStudent.Base.ID, "student")
	if err != nil {
		return "", err
	}
	return generatedToken, nil
}
