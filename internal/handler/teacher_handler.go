package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type teacherHandler struct {
	teacherService service.TeacherServiceInterface
}

func NewTeacherHandler(teacherService service.TeacherServiceInterface) *teacherHandler {
	return &teacherHandler{teacherService}
}

func (h *teacherHandler) Insert(c *gin.Context) {
	var req dtos.TeacherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	teacher, _, err := h.teacherService.Insert(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create teacher: " + err.Error()})
		return
	}
	response := dtos.TeacherResponse{
		ID:      teacher.ID,
		NIK:     teacher.NIK,
		Profile: teacher.Profile,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *teacherHandler) FindAll(c *gin.Context) {
	teacher, err := h.teacherService.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func (h *teacherHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	teacher, err := h.teacherService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func (h *teacherHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	deleteTeacher, err := h.teacherService.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, deleteTeacher)
}
