package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type examQuestionHandler struct {
	examQuestionService service.ExamQuestionServiceInterface
}

func NewExamQuestionHandler(examQuestionService service.ExamQuestionServiceInterface) *examQuestionHandler {
	return &examQuestionHandler{examQuestionService}
}

func (h *examQuestionHandler) CreateExamQuestion(c *gin.Context) {
	var req dtos.ExamQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	examQuestion, err := h.examQuestionService.CreateExamQuestion(req.QuestionID, req.ExamID, req.DisplayOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, examQuestion)
}

func (h *examQuestionHandler) GetByExam(c *gin.Context) {
	exam := c.Param("id")
	examQuestion, err := h.examQuestionService.GetByExam(exam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, examQuestion)

}

func (h *examQuestionHandler) GetBySubject(c *gin.Context) {
	question := c.Param("id")
	examQuestion, err := h.examQuestionService.GetByQuestion(question)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, examQuestion)
}

func (h *examQuestionHandler) Delete(c *gin.Context) {
	examQuestion := c.Param("id")
	delete, err := h.examQuestionService.Delete(examQuestion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, delete)
}
