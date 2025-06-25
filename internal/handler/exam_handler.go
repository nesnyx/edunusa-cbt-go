package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/middleware"
	"cbt/internal/models"
	"cbt/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type examHandler struct {
	examService service.ExamService
}

func NewExamHandler(examService service.ExamService) *examHandler {
	return &examHandler{examService}
}

func (h *examHandler) InsertNewExam(c *gin.Context) {
	var req dtos.ExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	currentUserData := c.MustGet("currentUser").(middleware.ClaimResult)
	fmt.Print("current Data : ", currentUserData)

	exam, err := h.examService.Insert(&req, currentUserData.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error created new exam: " + err.Error()})
		return
	}

	response := dtos.ExamResponse{
		ID:                 exam.ID,
		CreatedByTeacherID: exam.CreatedByTeacherID,
		AccessTokenExam:    exam.AccessToken,
		ExamTitle:          exam.ExamTitle,
	}
	c.JSON(http.StatusCreated, response)

}

func (h *examHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	exam, err := h.examService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exam Doesnt Exist"})
		return
	}

	response := dtos.ExamResponse{
		ID:                 exam.ID,
		CreatedByTeacherID: exam.CreatedByTeacherID,
		AccessTokenExam:    exam.AccessToken,
		ExamTitle:          exam.ExamTitle,
	}
	c.JSON(http.StatusOK, response)

}

func (h *examHandler) FindByTeacherID(c *gin.Context) {
	currentUser, _ := c.Get(middleware.ContextCurrentUser)
	exam, err := h.examService.FindByTeacherID(currentUser.(middleware.ClaimResult).ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exam Doesnt Exist"})
		return
	}
	response := dtos.ExamResponse{
		ID:                 exam.ID,
		CreatedByTeacherID: exam.CreatedByTeacherID,
		AccessTokenExam:    exam.AccessToken,
		ExamTitle:          exam.ExamTitle,
		Instruction:        exam.Instructions,
		DurationMinutes:    exam.DurationMinutes,
		StartDatetime:      exam.StartDatetime.Unix(),
		EndDatetime:        exam.EndDatetime.Unix(),
		Status:             models.ExamStatus(exam.Status),
	}
	c.JSON(http.StatusOK, response)
}

func (h *examHandler) DeleteExam(c *gin.Context) {
	id := c.Param("id")
	deleteExam, err := h.examService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error})
		return
	}
	c.JSON(http.StatusOK, deleteExam)
}

func (h *examHandler) UpdateExam(c *gin.Context) {
	id := c.Param("id")
	var req dtos.ExamRequestUpdate
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id harusa ada"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	exam, err := h.examService.Update(id, req.Instructions, req.ClassID, req.DurationMinutes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error created new exam: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, exam)
}
