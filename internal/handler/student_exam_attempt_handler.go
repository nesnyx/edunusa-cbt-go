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

func (h *studentExamAttemptHandler) StartExamination(c *gin.Context) {
	currentUser, _ := c.Get(middleware.ContextCurrentUser)
	examId := c.Query("examId")
	token := c.Query("token")
	if err := c.ShouldBindQuery(examId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	studenExamAttempt, err := h.studentExamAttemptService.Insert(currentUser.(middleware.ClaimResult).ID, examId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error created new exam: " + err.Error()})
		return
	}
	examTokenUsage, err := h.examTokenUsage.Create(token, examId, currentUser.(middleware.ClaimResult).ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error created new exam: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": map[string]string{
			"token":   examTokenUsage.TokenValueUsed,
			"student": string(studenExamAttempt.Student.Profile),
		},
		"msg": "oke",
	})
}
