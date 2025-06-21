package dtos

import "github.com/google/uuid"

type SubjectRequest struct {
	SubjectName string    `json:"subject_name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	ClassID     uuid.UUID `json:"class_id" binding:"required"`
}

type SubjectResponse struct {
	ID          string `json:"id" binding:"required"`
	SubjectName string `json:"name" binding:"required"`
}
