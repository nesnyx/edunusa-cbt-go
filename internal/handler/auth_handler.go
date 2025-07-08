package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService service.AuthServiceInterface
}

func NewAuthHandler(authService service.AuthServiceInterface) *authHandler {
	return &authHandler{authService}
}

func (h *authHandler) LoginTeacher(c *gin.Context) {
	var req dtos.LoginTeacher
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	token, err := h.authService.LoginTeacher(req.NIK, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (h *authHandler) LoginStudent(c *gin.Context) {
	var req dtos.LoginStudent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	token, err := h.authService.LoginStudent(req.NIS, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
