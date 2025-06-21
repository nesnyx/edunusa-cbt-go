package handlerextention

import (
	"cbt/extentions/dtos"
	serviceextention "cbt/extentions/serviceExtention"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentHandler struct {
	studentService serviceextention.StudentServiceInterface
}

func NewStudentHandler(studentService serviceextention.StudentServiceInterface) *studentHandler {
	return &studentHandler{studentService}
}

func (h *studentHandler) InsertStudent(c *gin.Context) {
	var req dtos.InsertStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	student, err := h.studentService.InsertStudent(&req)
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
