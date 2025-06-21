package models

import (
	"cbt/extentions/models"
	"encoding/json"

	"github.com/google/uuid"
)

// Question merepresentasikan tabel 'questions'
type Question struct {
	Base
	QuestionBankID     uuid.UUID       `gorm:"type:varchar(255);null" json:"question_bank_id,omitempty"` // Bisa null jika soal dibuat langsung untuk ujian
	QuestionText       string          `gorm:"type:text;not null" json:"question_text"`
	QuestionType       QuestionType    `gorm:"type:varchar(50);not null" json:"question_type"` // Gunakan tipe QuestionType
	Points             float64         `gorm:"default:1.0" json:"points"`
	Metadata           json.RawMessage `gorm:"null" json:"metadata,omitempty"` // Untuk opsi MCQ, kunci jawaban, dll.
	CreatedByTeacherID uuid.UUID       `gorm:"type:varchar(255);not null" json:"created_by_teacher_id"`
	QuestionBank       QuestionBank    `gorm:"foreignKey:QuestionBankID;references:ID" json:"question_bank,omitempty"`
	CreatedByTeacher   models.Teacher  `gorm:"foreignKey:CreatedByTeacherID;references:ID" json:"created_by_teacher,omitempty"`
	ExamQuestions      []ExamQuestion  `gorm:"foreignKey:QuestionID" json:"-"` // Relasi Many-to-Many dengan Exam melalui ExamQuestion
}

func (Question) TableName() string {
	return "exam_engine.question"
}
