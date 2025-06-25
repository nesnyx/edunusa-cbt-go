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
