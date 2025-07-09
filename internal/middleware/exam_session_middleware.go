// middleware/exam_session_middleware.go
package middleware

import (
	"cbt/internal/models"
	"cbt/internal/repository"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExamSessionStatus string

const (
	ExamSessionActive   ExamSessionStatus = "active"
	ExamSessionExpired  ExamSessionStatus = "expired"
	ExamSessionFinished ExamSessionStatus = "finished"
	ExamSessionNew      ExamSessionStatus = "new"
)

type ExamSessionInfo struct {
	Status           ExamSessionStatus `json:"status"`
	ExamID           string            `json:"exam_id"`
	StudentID        string            `json:"student_id"`
	StartTime        time.Time         `json:"start_time"`
	EndTime          time.Time         `json:"end_time"`
	DurationMinutes  int               `json:"duration_minutes"`
	RemainingMinutes int               `json:"remaining_minutes"`
	CanContinue      bool              `json:"can_continue"`
	AttemptID        string            `json:"attempt_id,omitempty"`
}

func EnhancedExamSessionMiddleware(db *gorm.DB) gin.HandlerFunc {
	tokenUsageRepo := repository.NewExamTokenUsageRepository(db)
	examRepo := repository.NewExamRepository(db)
	attemptRepo := repository.NewStudentExamAttemptRepository(db)

	return func(c *gin.Context) {
		currentUser, exists := c.Get(ContextCurrentUser)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		user := currentUser.(ClaimResult)
		if user.Role != "student" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "only students can access exam sessions"})
			return
		}

		studentID := user.ID
		examID := c.Query("examId")
		if examID == "" {
			examID = c.PostForm("examId")
		}
		if examID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "examId is required"})
			return
		}

		// Validate exam session
		sessionInfo, err := validateExamSession(studentID, examID, tokenUsageRepo, examRepo, attemptRepo)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Set session info to context
		c.Set("examSession", sessionInfo)

		// Handle different session statuses
		switch sessionInfo.Status {
		case ExamSessionExpired:
			c.AbortWithStatusJSON(http.StatusGone, gin.H{
				"error":   "exam session has expired",
				"details": sessionInfo,
			})
			return
		case ExamSessionFinished:
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error":   "exam already finished",
				"details": sessionInfo,
			})
			return
		case ExamSessionActive, ExamSessionNew:
			c.Next()
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid session status"})
			return
		}
	}
}

func validateExamSession(
	studentID, examID string,
	tokenUsageRepo repository.ExamTokenUsageRepositoryInterface,
	examRepo repository.ExamRepository,
	attemptRepo repository.StudentExamAttemptRepositoryInterface,
) (*ExamSessionInfo, error) {
	// 1. Get exam details
	exam, err := examRepo.GetExamByID(examID)
	if err != nil {
		return nil, errors.New("exam not found")
	}

	// 2. Check if exam is available (within start/end datetime)
	now := time.Now()
	if now.Before(exam.StartDatetime) {
		return nil, errors.New("exam has not started yet")
	}
	if now.After(exam.EndDatetime) {
		return nil, errors.New("exam has already ended")
	}

	// 3. Check existing token usage
	tokenUsage, err := tokenUsageRepo.GetByStudentAndExam(studentID, examID)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// New session - create token usage
		return handleNewExamSession(studentID, examID, exam, tokenUsageRepo, attemptRepo)
	} else if err != nil {
		return nil, errors.New("failed to check token usage")
	}

	// 4. Existing session - validate
	return handleExistingExamSession(studentID, examID, exam, tokenUsage, attemptRepo)
}

func handleNewExamSession(
	studentID, examID string,
	exam *models.Exam,
	tokenUsageRepo repository.ExamTokenUsageRepositoryInterface,
	attemptRepo repository.StudentExamAttemptRepositoryInterface,
) (*ExamSessionInfo, error) {
	// Check if student already has an attempt (prevent duplicate)
	existingAttempt, err := attemptRepo.GetByStudentAndExam(studentID, examID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("failed to check existing attempt")
	}

	if existingAttempt != nil {
		// Student already has an attempt, check if it's finished
		if existingAttempt.Status == models.AttemptStatusSubmitted {
			return &ExamSessionInfo{
				Status:      ExamSessionFinished,
				ExamID:      examID,
				StudentID:   studentID,
				CanContinue: false,
				AttemptID:   existingAttempt.ID,
			}, nil
		}

		// Check if the attempt is expired
		if existingAttempt.AttemptStartTime != nil {
			expirationTime := existingAttempt.AttemptStartTime.Add(time.Duration(exam.DurationMinutes) * time.Minute)
			if time.Now().After(expirationTime) {
				// Clean up expired attempt and token
				attemptRepo.DeleteByStudent(existingAttempt.StudentID)
				return &ExamSessionInfo{
					Status:      ExamSessionExpired,
					ExamID:      examID,
					StudentID:   studentID,
					CanContinue: false,
				}, nil
			}
		}
	}

	// Create new token usage
	newTokenUsage := &models.ExamTokenUsage{
		ExamID:         examID,
		StudentID:      studentID,
		TokenValueUsed: exam.AccessToken,
		UsageTimestamp: time.Now(),
	}

	_, err = tokenUsageRepo.Create(newTokenUsage)
	if err != nil {
		return nil, errors.New("failed to create token usage")
	}

	return &ExamSessionInfo{
		Status:           ExamSessionNew,
		ExamID:           examID,
		StudentID:        studentID,
		StartTime:        newTokenUsage.UsageTimestamp,
		EndTime:          newTokenUsage.UsageTimestamp.Add(time.Duration(exam.DurationMinutes) * time.Minute),
		DurationMinutes:  exam.DurationMinutes,
		RemainingMinutes: exam.DurationMinutes,
		CanContinue:      true,
	}, nil
}

func handleExistingExamSession(
	studentID, examID string,
	exam *models.Exam,
	tokenUsage *models.ExamTokenUsage,
	attemptRepo repository.StudentExamAttemptRepositoryInterface,
) (*ExamSessionInfo, error) {
	// Calculate session expiration
	sessionStart := tokenUsage.UsageTimestamp
	sessionEnd := sessionStart.Add(time.Duration(exam.DurationMinutes) * time.Minute)
	now := time.Now()

	// Check if session is expired
	if now.After(sessionEnd) {
		// Clean up expired session
		attemptRepo.DeleteByStudentAndExam(studentID, examID)
		// Note: You might want to also clean up the token usage or mark it as expired
		return &ExamSessionInfo{
			Status:      ExamSessionExpired,
			ExamID:      examID,
			StudentID:   studentID,
			StartTime:   sessionStart,
			EndTime:     sessionEnd,
			CanContinue: false,
		}, nil
	}

	// Check existing attempt
	attempt, err := attemptRepo.GetByStudentAndExam(studentID, examID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("failed to check existing attempt")
	}

	if attempt != nil && attempt.Status == models.AttemptStatusSubmitted {
		return &ExamSessionInfo{
			Status:      ExamSessionFinished,
			ExamID:      examID,
			StudentID:   studentID,
			StartTime:   sessionStart,
			EndTime:     sessionEnd,
			CanContinue: false,
			AttemptID:   attempt.ID,
		}, nil
	}

	// Session is active
	remainingMinutes := int(sessionEnd.Sub(now).Minutes())
	if remainingMinutes < 0 {
		remainingMinutes = 0
	}

	return &ExamSessionInfo{
		Status:           ExamSessionActive,
		ExamID:           examID,
		StudentID:        studentID,
		StartTime:        sessionStart,
		EndTime:          sessionEnd,
		DurationMinutes:  exam.DurationMinutes,
		RemainingMinutes: remainingMinutes,
		CanContinue:      true,
		AttemptID: func() string {
			if attempt != nil {
				return attempt.ID
			}
			return ""
		}(),
	}, nil
}
