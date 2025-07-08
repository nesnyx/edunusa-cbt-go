package dtos

import (
	"encoding/json"
)

type TeacherRequest struct {
	NIK      string          `json:"nik" binding:"required"`
	Profile  json.RawMessage `json:"profile" binding:"required"`
	Password string          `json:"password" binding:"required"`
}

type TeacherResponse struct {
	ID      string          `json:"id"`
	NIK     string          `json:"nik"`
	Profile json.RawMessage `json:"profile"`
}
