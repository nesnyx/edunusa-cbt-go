package handler

import (
	"cbt/internal/middleware"
	"cbt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentExamAttemptHandler struct {
	studentExamAttemptService service.StudentExamAttemptServiceInterface
	examTokenUsage            service.ExamTokenUsageServiceInterface
}

func NewStudentExamAttemptHandler(studentExamAttemptService service.StudentExamAttemptServiceInterface, examTokenUsage service.ExamTokenUsageServiceInterface) *studentExamAttemptHandler {
	return &studentExamAttemptHandler{studentExamAttemptService: studentExamAttemptService, examTokenUsage: examTokenUsage}
}

// handler/student_exam_attempt_handler.go - Enhanced version
func (h *studentExamAttemptHandler) StartExamination(c *gin.Context) {
	sessionInfo, exists := c.Get("examSession")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "exam session not found"})
		return
	}

	session := sessionInfo.(*middleware.ExamSessionInfo)

	// Start or continue exam attempt
	attempt, err := h.studentExamAttemptService.StartOrContinueExam(session.StudentID, session.ExamID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "exam session ready",
		"session": &session,
		"attempt": attempt,
	})
}

func (h *studentExamAttemptHandler) GetExamProgress(c *gin.Context) {
	attemptID := c.Param("attemptId")

	attempt, err := h.studentExamAttemptService.GetAttemptProgress(attemptID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attempt": attempt,
	})
}

func (h *studentExamAttemptHandler) FinishExam(c *gin.Context) {
	attemptID := c.Param("attemptId")
	currentUser, _ := c.Get(middleware.ContextCurrentUser)
	studentID := currentUser.(middleware.ClaimResult).ID
	attempt, err := h.studentExamAttemptService.FinishExam(attemptID, studentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "exam finished successfully",
		"attempt": attempt,
	})
}
