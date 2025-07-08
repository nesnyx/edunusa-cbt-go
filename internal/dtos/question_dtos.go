package dtos

import (
	"cbt/internal/models"
	"encoding/json"
)

type QuestionRequest struct {
	QuestionText string              `json:"question_text"`
	QuestionType models.QuestionType `json:"question_type"`
	Points       float64             `json:"points"`
	QuestionBank string              `json:"question_bank"`
	Metadata     json.RawMessage     `json:"metadata"`
}

type QuestionBankRequest struct {
	BankName    string `json:"bank_name"`
	Description string `json:"description"`
	Subject     string `json:"subject"`
}

type QuestionBankResponse struct {
	ID          string          `json:"id"`
	BankName    string          `json:"bank_name"`
	Description string          `json:"description"`
	SubjectName string          `json:"subject_name"`
	TeacherNik  string          `json:"teacher_nik"`
	Teacher     json.RawMessage `json:"teacher"`
}

type QuestionResponse struct {
	ID           string          `json:"id"`
	QuestionText string          `json:"question_text"`
	QuestionType string          `json:"question_type"`
	Points       float64         `json:"points"`
	Metadata     json.RawMessage `json:"metadata"`
	BankName     string          `json:"bank_name,omitempty"`
	CreatedByNIK string          `json:"created_by_nik,omitempty"`
}
