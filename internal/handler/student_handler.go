package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type studentHandler struct {
	studentService service.StudentServiceInterface
}

func NewStudentHandler(studentService service.StudentServiceInterface) *studentHandler {
	return &studentHandler{studentService}
}

func (h *studentHandler) InsertStudent(c *gin.Context) {
	var req dtos.InsertStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	student, _, err := h.studentService.InsertStudent(&req)
	if err != nil {
		if err.Error() == "nis already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		}
		return
	}
	response := dtos.StudentResponse{
		ID:      student.ID,
		NIS:     student.NIS,
		Profile: student.Profile,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *studentHandler) FindAllStudent(c *gin.Context) {
	students, err := h.studentService.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find all students: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": students,
	})
}

func (h *studentHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.studentService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find all students: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func (h *studentHandler) FindByNIS(c *gin.Context) {
	nis := c.Param("nis")
	student, err := h.studentService.FindByNIS(nis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find all students: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}
