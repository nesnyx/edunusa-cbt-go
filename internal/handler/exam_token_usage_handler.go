package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/middleware"
	"cbt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type examTokenUsageHandler struct {
	examTokenUsageService service.ExamTokenUsageServiceInterface
}

func NewExamTokenUsageHandler(examTokenUsageService service.ExamTokenUsageServiceInterface) *examTokenUsageHandler {
	return &examTokenUsageHandler{examTokenUsageService}
}

func (h *examTokenUsageHandler) Insert(c *gin.Context) {
	var req dtos.ExamTokenUsageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	currentUserData := c.MustGet("currentUser").(middleware.ClaimResult)
	examTokenUsage, err := h.examTokenUsageService.Create(req.TokenValueUsed, req.ExamID, currentUserData.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error created new exam: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, examTokenUsage)

}

func (h *examTokenUsageHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	delete, err := h.examTokenUsageService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error delete exam: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": delete})

}

func (h *examTokenUsageHandler) FindByStudent(c *gin.Context) {
	id := c.Param("id")
	exam, err := h.examTokenUsageService.FindByStudent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error delete exam: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": exam})
}
