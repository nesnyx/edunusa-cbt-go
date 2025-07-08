package models

import (
	"time"
)

// StudentExamAttempt merepresentasikan tabel 'student_exam_attempts'
type StudentExamAttempt struct {
	Base
	StudentID         string                   `gorm:"type:varchar(255);not null;index:idx_student_exam,unique" json:"student_id"`
	ExamID            string                   `gorm:"type:varchar(255);not null;index:idx_student_exam,unique" json:"exam_id"`
	AttemptStartTime  *time.Time               `gorm:"null" json:"attempt_start_time,omitempty"`
	AttemptEndTime    *time.Time               `gorm:"null" json:"attempt_end_time,omitempty"`
	Score             *float64                 `gorm:"null" json:"score,omitempty"`
	Status            StudentExamAttemptStatus `gorm:"type:varchar(50);not null;default:'not_started'" json:"status"`
	SubmittedAt       *time.Time               `gorm:"null" json:"submitted_at,omitempty"`
	GradedByTeacherID *string                  `gorm:"type:varchar(255);null" json:"graded_by_teacher_id,omitempty"`
	TeacherFeedback   string                   `gorm:"type:text;null" json:"teacher_feedback,omitempty"`

	Student         Student         `gorm:"foreignKey:StudentID;references:ID" json:"student,omitempty"`
	Exam            Exam            `gorm:"foreignKey:ExamID;references:ID" json:"exam,omitempty"`
	GradedByTeacher *Teacher        `gorm:"foreignKey:GradedByTeacherID;references:ID" json:"graded_by_teacher,omitempty"`
	StudentAnswers  []StudentAnswer `gorm:"foreignKey:StudentExamAttemptID" json:"student_answers,omitempty"`
}

type StudentAnswer struct {
	Base
	StudentExamAttemptID   string             `gorm:"type:varchar(255);not null" json:"student_exam_attempt_id"`
	ExamQuestionID         string             `gorm:"type:varchar(255);not null" json:"exam_question_id"`
	AnswerData             string             `gorm:"null" json:"answer_data,omitempty"`
	MarksAwarded           *float64           `gorm:"null" json:"marks_awarded,omitempty"`
	IsCorrect              *bool              `gorm:"null" json:"is_correct,omitempty"`
	TeacherCommentOnAnswer string             `gorm:"type:text;null" json:"teacher_comment_on_answer,omitempty"`
	StudentExamAttempt     StudentExamAttempt `gorm:"foreignKey:StudentExamAttemptID;references:ID" json:"-"`
	ExamQuestion           ExamQuestion       `gorm:"foreignKey:ExamQuestionID;references:ID" json:"-"`
}

func (StudentExamAttempt) TableName() string {
	return "exam_engine.student_exam_attempt"
}

func (StudentAnswer) TableName() string {
	return "exam_engine.student_answer"
}
