package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentAnswerHandler struct {
	service service.StudentAnswerServiceInterface
}

func NewStudentAnswerHandler(service service.StudentAnswerServiceInterface) *studentAnswerHandler {
	return &studentAnswerHandler{service}
}

func (h *studentAnswerHandler) InsertOrUpdate(c *gin.Context) {
	var req dtos.StudentAnswerRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	studentAnswer, err := h.service.InsertOrUpdate(&req, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error created new exam: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, studentAnswer)

}

func (h *studentAnswerHandler) GetByQuestionAndStudentAttempt(c *gin.Context) {
	var req dtos.GetStudentAnswerByQuestionAndStudentAttempt
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	studentAnswer, err := h.service.GetByQuestionAndStudentAttempt(req.StudentID, req.ExamQuestion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error get answer: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, studentAnswer)
}
