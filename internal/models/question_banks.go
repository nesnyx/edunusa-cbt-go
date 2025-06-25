package models

import (
	models "cbt/extentions/models"
)

// QuestionBank merepresentasikan tabel 'question_banks'
type QuestionBank struct {
	Base
	BankName           string         `gorm:"type:varchar(255);not null" json:"bank_name"`
	SubjectID          string         `gorm:"type:varchar(255);not null" json:"subject_id"`
	CreatedByTeacherID string         `gorm:"type:varchar(255);not null" json:"created_by_teacher_id"`
	Description        string         `gorm:"type:text;null" json:"description,omitempty"`
	Subject            models.Subject `gorm:"foreignKey:SubjectID;references:ID" json:"subject,omitempty"`
	CreatedByTeacher   models.Teacher `gorm:"foreignKey:CreatedByTeacherID;references:ID" json:"created_by_teacher,omitempty"`
	Questions          []Question     `gorm:"foreignKey:QuestionBankID" json:"questions,omitempty"` // Relasi Has Many
}

// Mendefinisikan tipe untuk QuestionType agar lebih terkontrol
type QuestionType string

const (
	MultipleChoice   QuestionType = "multiple_choice"
	MultipleResponse QuestionType = "multiple_response"
	Essay            QuestionType = "essay"
	TrueFalse        QuestionType = "true_false"
	FillInTheBlanks  QuestionType = "fill_in_the_blanks"
	Matching         QuestionType = "matching"
)

func (QuestionBank) TableName() string {
	return "exam_engine.question_bank"
}
