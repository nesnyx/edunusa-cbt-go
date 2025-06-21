package dtos

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Student DTOs
type InsertStudentRequest struct {
	NIS      string          `json:"nis" binding:"required"`
	Password string          `json:"password" binding:"required"`
	Profile  json.RawMessage `json:"profile" binding:"required"`
}

type StudentResponse struct {
	ID      uuid.UUID       `json:"id"`
	NIS     string          `json:"nis" binding:"required,min=3"`
	Profile json.RawMessage `json:"profile" binding:"required"`
}
