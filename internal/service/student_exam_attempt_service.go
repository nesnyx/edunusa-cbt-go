package service

import (
	"cbt/internal/models"
	"cbt/internal/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type StudentExamAttemptServiceInterface interface {
	StartOrContinueExam(studentID, examID string) (*models.StudentExamAttempt, error)
	FinishExam(attemptID string, studentId string) (*models.StudentExamAttempt, error)
	GetAttemptProgress(attemptID string) (*models.StudentExamAttempt, error)
}

type studentExamAttemptService struct {
	studentExamAttemptRepo repository.StudentExamAttemptRepositoryInterface
	examRepo               repository.ExamRepository
	tokenUsageRepo         repository.ExamTokenUsageRepositoryInterface
}

func NewStudentExamAttemptService(
	studentExamAttemptRepo repository.StudentExamAttemptRepositoryInterface,
	examRepo repository.ExamRepository,
	tokenUsageRepo repository.ExamTokenUsageRepositoryInterface,
) *studentExamAttemptService {
	return &studentExamAttemptService{
		studentExamAttemptRepo: studentExamAttemptRepo,
		examRepo:               examRepo,
		tokenUsageRepo:         tokenUsageRepo,
	}
}

func (s *studentExamAttemptService) StartOrContinueExam(studentID, examID string) (*models.StudentExamAttempt, error) {
	// Check if attempt already exists
	existingAttempt, err := s.studentExamAttemptRepo.GetByStudentAndExam(studentID, examID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingAttempt != nil {
		// Continue existing attempt
		if existingAttempt.Status == models.StudentExamAttemptStatus(models.ExamStatusOngoing) {
			return nil, errors.New("exam already completed")
		}

		// Check if attempt is still valid (not expired)
		exam, err := s.examRepo.GetExamByID(examID)
		if err != nil {
			return nil, err
		}

		if existingAttempt.AttemptStartTime != nil {
			expirationTime := existingAttempt.AttemptStartTime.Add(time.Duration(exam.DurationMinutes) * time.Minute)
			if time.Now().After(expirationTime) {
				// Clean up expired attempt
				s.studentExamAttemptRepo.DeleteByStudentAndExam(existingAttempt.ID, studentID)
				s.tokenUsageRepo.DeleteByStudentAndExam(studentID, examID)
				return nil, errors.New("exam session has expired")
			}
		}

		return existingAttempt, nil
	}

	// Create new attempt
	now := time.Now()
	newAttempt := &models.StudentExamAttempt{
		ExamID:           examID,
		StudentID:        studentID,
		AttemptStartTime: &now,
		Status:           models.AttemptStatusOngoing,
	}

	return s.studentExamAttemptRepo.Create(newAttempt)
}

func (s *studentExamAttemptService) FinishExam(attemptID string, studentId string) (*models.StudentExamAttempt, error) {
	attempt, err := s.studentExamAttemptRepo.GetByID(attemptID)
	if err != nil {
		return nil, err
	}

	if attempt.Status == models.StudentExamAttemptStatus(models.ExamStatusCompleted) {
		return nil, errors.New("exam already completed")
	}

	now := time.Now()
	attempt.AttemptEndTime = &now
	attempt.SubmittedAt = &now
	attempt.Status = models.StudentExamAttemptStatus(models.ExamStatusCompleted)

	return s.studentExamAttemptRepo.Update(attemptID, attempt)
}

func (s *studentExamAttemptService) GetAttemptProgress(attemptID string) (*models.StudentExamAttempt, error) {
	attempt, err := s.studentExamAttemptRepo.GetByID(attemptID)
	if err != nil {
		return nil, err
	}

	// Check if attempt is expired
	exam, err := s.examRepo.GetExamByID(attempt.ExamID)
	if err != nil {
		return nil, err
	}

	if attempt.AttemptStartTime != nil {
		expirationTime := attempt.AttemptStartTime.Add(time.Duration(exam.DurationMinutes) * time.Minute)
		if time.Now().After(expirationTime) && attempt.Status != models.StudentExamAttemptStatus(models.ExamStatusCompleted) {

			// Auto-finish expired attempt
			now := time.Now()
			attempt.AttemptEndTime = &now
			attempt.Status = models.StudentExamAttemptStatus(models.ExamStatusCompleted)
			s.studentExamAttemptRepo.Update(attemptID, attempt)
		}
	}

	return attempt, nil
}
