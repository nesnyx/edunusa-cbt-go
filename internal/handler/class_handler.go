package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type classHandler struct {
	classService service.ClassServiceInterface
}

func NewClassHandler(classService service.ClassServiceInterface) *classHandler {
	return &classHandler{classService}
}

func (h *classHandler) Insert(c *gin.Context) {
	var req dtos.ClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	class, err := h.classService.Insert(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := dtos.ClassResponse{
		ID:          class.Base.ID,
		GradeLevel:  class.GradeLevel,
		Description: class.Description,
		ClassName:   class.ClassName,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *classHandler) FindAll(c *gin.Context) {
	class, err := h.classService.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": class})
}

func (h *classHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	class, err := h.classService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
}
