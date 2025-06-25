package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/middleware"
	"cbt/internal/service"

	"github.com/gin-gonic/gin"
)

type questionHandler struct {
	questionService service.QuestionServiceInterface
}

func NewQuestionHandler(questionService service.QuestionServiceInterface) *questionHandler {
	return &questionHandler{questionService}
}

func (h *questionHandler) CreateQuestion(c *gin.Context) {
	var req dtos.QuestionRequest
	currentUser, _ := c.Get(middleware.ContextCurrentUser)
	teacherID := currentUser.(middleware.ClaimResult).ID
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	question, err := h.questionService.CreateQuestion(&req, req.QuestionBank, teacherID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, question)
}

func (h *questionHandler) DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	_, err := h.questionService.DeleteQuestion(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "soal berhasil dihapus"})
}

func (h *questionHandler) UpdateQuestion(c *gin.Context) {
	var req dtos.QuestionRequest
	id := c.Param("id")
	currentUser, _ := c.Get(middleware.ContextCurrentUser)
	teacherID := currentUser.(middleware.ClaimResult).ID
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	question, err := h.questionService.UpdateQuestion(&req, teacherID, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, question)
}

func (h *questionHandler) GetByTeacher(c *gin.Context) {
	currentUser, _ := c.Get(middleware.ContextCurrentUser)
	teacherID := currentUser.(middleware.ClaimResult).ID
	question, err := h.questionService.GetByTeacher(teacherID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, question)
}
