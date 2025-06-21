package dtos

import "github.com/google/uuid"

type RoleRequest struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" binding:"required"`
}

type RoleResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
