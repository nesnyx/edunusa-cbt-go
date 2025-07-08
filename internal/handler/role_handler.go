package handler

import (
	"cbt/internal/dtos"
	"cbt/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	service service.RoleServiceInterface
}

func NewRoleHandler(service service.RoleServiceInterface) *RoleHandler {
	return &RoleHandler{service}
}

func (h *RoleHandler) InsertRole(c *gin.Context) {
	var req dtos.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	role, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := dtos.RoleResponse{
		ID:   role.ID,
		Name: role.RoleName,
	}
	c.JSON(http.StatusCreated, response)

}

func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}
