package models

import "github.com/google/uuid"

// ExamQuestion merepresentasikan tabel 'exam_questions' (junction table)
type ExamQuestion struct {
	Base
	ExamID         uuid.UUID       `gorm:"type:varchar(255);not null;index:idx_exam_question,unique" json:"exam_id"`
	QuestionID     uuid.UUID       `gorm:"type:varchar(255);not null;index:idx_exam_question,unique" json:"question_id"`
	DisplayOrder   int             `gorm:"not null" json:"display_order"`
	PointsOverride float64         `gorm:"null" json:"points_override,omitempty"` // Jika poin soal ini beda dari default
	Exam           Exam            `gorm:"foreignKey:ExamID;references:ID" json:"exam,omitempty"`
	Question       Question        `gorm:"foreignKey:QuestionID;references:ID" json:"question,omitempty"`
	StudentAnswers []StudentAnswer `gorm:"foreignKey:ExamQuestionID" json:"-"`
}

// Mendefinisikan tipe untuk StudentExamAttemptStatus
type StudentExamAttemptStatus string

const (
	AttemptStatusNotStarted        StudentExamAttemptStatus = "not_started"
	AttemptStatusOngoing           StudentExamAttemptStatus = "ongoing"
	AttemptStatusSubmitted         StudentExamAttemptStatus = "submitted"
	AttemptStatusGradingInProgress StudentExamAttemptStatus = "grading_in_progress"
	AttemptStatusGraded            StudentExamAttemptStatus = "graded"
	AttemptStatusCancelled         StudentExamAttemptStatus = "cancelled"
)

func (ExamQuestion) TableName() string {
	return "exam_engine.exam_questions"
}
