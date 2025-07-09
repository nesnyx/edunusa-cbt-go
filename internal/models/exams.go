package models

import (
	"time"
)

type AttemptStatusCompleted string

const (
	ExamStatusDraft     AttemptStatusCompleted = "draft"
	ExamStatusPublished AttemptStatusCompleted = "published"
	ExamStatusOngoing   AttemptStatusCompleted = "ongoing"
	ExamStatusCompleted AttemptStatusCompleted = "completed"
	ExamStatusArchived  AttemptStatusCompleted = "archived"
)

// Exam merepresentasikan tabel 'exams' (dengan penambahan)
type Exam struct {
	Base
	ExamTitle             string                 `gorm:"type:varchar(255);not null" json:"exam_title"`
	SubjectID             string                 `gorm:"type:varchar(255);not null" json:"subject_id"`
	ClassID               string                 `gorm:"type:varchar(255);not null" json:"class_id"`
	CreatedByTeacherID    string                 `gorm:"type:varchar(255);not null" json:"created_by_teacher_id"`
	Instructions          string                 `gorm:"type:text;null" json:"instructions,omitempty"`
	StartDatetime         time.Time              `gorm:"not null" json:"start_datetime"`
	EndDatetime           time.Time              `gorm:"not null" json:"end_datetime"`
	DurationMinutes       int                    `gorm:"not null" json:"duration_minutes"`
	AccessToken           string                 `gorm:"type:varchar(50);uniqueIndex;not null" json:"access_token"`
	AccessTokenValidUntil *time.Time             `gorm:"null" json:"access_token_valid_until,omitempty"` // TAMBAHAN: Pointer agar bisa null
	Status                AttemptStatusCompleted `gorm:"type:varchar(50);not null;default:'draft'" json:"status"`
	RandomizeQuestions    bool                   `gorm:"default:false" json:"randomize_questions"`
	TotalPoints           float64                `gorm:"null" json:"total_points,omitempty"`
	PassingScore          float64                `gorm:"null" json:"passing_score,omitempty"`
	ShowAnswersAfterExam  bool                   `gorm:"default:false" json:"show_answers_after_exam"`

	// Teacher         models.Teacher       `gorm:"foreignKey:CreatedByTeacherID" json:"created_by_teacher,omitempty"`
	Teacher         Teacher              `gorm:"foreignKey:CreatedByTeacherID;references:ID" json:"teacher,omitempty"`
	Subject         Subject              `gorm:"foreignKey:SubjectID;references:ID" json:"subject,omitempty"`
	Class           Class                `gorm:"foreignKey:ClassID;references:ID" json:"class,omitempty"`
	ExamQuestions   []ExamQuestion       `gorm:"foreignKey:ExamID" json:"exam_questions,omitempty"`
	StudentAttempts []StudentExamAttempt `gorm:"foreignKey:ExamID" json:"student_attempts,omitempty"`
	TokenUsages     []ExamTokenUsage     `gorm:"foreignKey:ExamID" json:"-"` // Relasi ke log penggunaan token
}

// ExamTokenUsage merepresentasikan tabel 'exam_token_usages' (BARU)
type ExamTokenUsage struct {
	Base
	StudentID      string    `gorm:"type:varchar(255);not null" json:"student_id"`
	ExamID         string    `gorm:"type:varchar(255);not null" json:"exam_id"`
	TokenValueUsed string    `gorm:"type:varchar(50);not null" json:"token_value_used"`
	UsageTimestamp time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"usage_timestamp"`

	Student Student `gorm:"foreignKey:StudentID;references:ID" json:"student,omitempty"`
	Exam    Exam    `gorm:"foreignKey:ExamID;references:ID" json:"exam,omitempty"`
}

func (Exam) TableName() string {
	return "exam_engine.exam"
}

func (ExamTokenUsage) TableName() string {
	return "exam_engine.exam_token_usage"
}
