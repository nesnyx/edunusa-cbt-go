package service

import (
	"cbt/internal/auth"

	"cbt/internal/models"
	"cbt/internal/repository"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	RegisterUser(req models.RegisterUserRequest) (*models.User, error)
	LoginUser(req models.LoginUserRequest) (*models.User, string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *authService {
	return &authService{userRepo: userRepo}
}

func (s *authService) RegisterUser(req models.RegisterUserRequest) (*models.User, error) {

	existingUser, err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("error checking username: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	role, err := s.userRepo.GetRoleByName(req.RoleName)
	if err != nil {
		return nil, fmt.Errorf("error getting role: %w", err)
	}
	if role == nil {
		log.Printf("Warning: Role '%s' not found. Please ensure roles are seeded in the database.", req.RoleName)
		return nil, errors.New("role not found: " + req.RoleName)
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := &models.User{
		Base: models.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username:     req.Username,
		PasswordHash: hashedPassword,
		FullName:     req.FullName,
		Email:        req.Email,
		RoleID:       role.ID,
	}

	err = s.userRepo.CreateUser(newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	newUser.Role = *role
	return newUser, nil
}

func (s *authService) LoginUser(req models.LoginUserRequest) (*models.User, string, error) {
	user, err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, "", fmt.Errorf("error retrieving user: %w", err)
	}
	if user == nil {
		return nil, "", errors.New("invalid username or password")
	}

	if !auth.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, "", errors.New("invalid username or password")
	}

	if user.Role.RoleName == "" {

		roleFromDB, err := s.userRepo.GetRoleByName(user.Role.RoleName)
		if err != nil || roleFromDB == nil {
			log.Printf("Error fetching role details for user %s: %v", user.Username, err)
			return nil, "", errors.New("could not determine user role")
		}
		user.Role = *roleFromDB
	}

	token, err := auth.GenerateToken(user.ID, user.Username, user.Role.RoleName)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return user, token, nil
}
