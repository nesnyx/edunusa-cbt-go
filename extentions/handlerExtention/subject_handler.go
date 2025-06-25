package handlerextention

import (
	"cbt/extentions/dtos"
	serviceextention "cbt/extentions/serviceExtention"
	"net/http"

	"github.com/gin-gonic/gin"
)

type subjectHandler struct {
	subjectService serviceextention.SubjectServiceInterface
}

func NewSubjectHandler(subjectService serviceextention.SubjectServiceInterface) *subjectHandler {
	return &subjectHandler{subjectService}
}

func (h *subjectHandler) InsertSubject(c *gin.Context) {
	var req dtos.SubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	subject, err := h.subjectService.Insert(&req)
	if err != nil {
		if err.Error() == "subject already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subject: " + err.Error()})
		}
		return
	}
	response := dtos.SubjectResponse{
		ID:          subject.Base.ID,
		SubjectName: subject.SubjectName,
	}
	c.JSON(http.StatusCreated, response)

}

func (h *subjectHandler) FindAll(c *gin.Context) {
	subjects, err := h.subjectService.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subjects)
}

func (h *subjectHandler) FindbyID(c *gin.Context) {
	id := c.Param("id")
	subjects, err := h.subjectService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subjects)
}
