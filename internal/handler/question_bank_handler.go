package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/middleware"
	"cbt/internal/service"

	"github.com/gin-gonic/gin"
)

type questionBankHandler struct {
	questionBankService service.QuestionBankServiceInterface
}

func NewQuestionBankHandler(questionBankService service.QuestionBankServiceInterface) *questionBankHandler {
	return &questionBankHandler{questionBankService}
}

func (h *questionBankHandler) CreateQuestionBank(c *gin.Context) {
	var req dtos.QuestionBankRequest
	currentUser, _ := c.Get(middleware.ContextCurrentUser)
	teacherID := currentUser.(middleware.ClaimResult).ID
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	questionBank, err := h.questionBankService.CreateQuestionBank(&req, teacherID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, questionBank)
}

func (h *questionBankHandler) DeleteQuestionBank(c *gin.Context) {
	id := c.Param("id")
	_, err := h.questionBankService.DeleteQuestionBank(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "soal berhasil dihapus"})
}

func (h *questionBankHandler) GetQuestionBankByTeacher(c *gin.Context) {
	currentUser, _ := c.Get(middleware.ContextCurrentUser)
	teacherID := currentUser.(middleware.ClaimResult).ID
	questionBank, err := h.questionBankService.GetQuestionBankByTeacher(teacherID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": questionBank})
}

func (h *questionBankHandler) GetBySubject(c *gin.Context) {
	id := c.Param("id")
	exam, err := h.questionBankService.GetBySubject(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": exam})

}

func (h *questionBankHandler) Update(c *gin.Context) {
	var req dtos.QuestionBankRequest
	currentUser, _ := c.Get(middleware.ContextCurrentUser)
	id := c.Param("id")
	teacherID := currentUser.(middleware.ClaimResult).ID
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	exam, err := h.questionBankService.UpdateQuestionBank(req.BankName, req.Description, id, teacherID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": exam})
}
