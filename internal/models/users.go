package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

// User merepresentasikan tabel 'users'
type User struct {
	Base
	Username     string          `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	PasswordHash string          `gorm:"type:varchar(255);not null" json:"-"` // Jangan kirim password hash ke client
	FullName     string          `gorm:"type:varchar(255);not null" json:"full_name"`
	Profile      json.RawMessage `json:"profile"`
	Email        string          `gorm:"type:varchar(255);uniqueIndex;null" json:"email,omitempty"`
	RoleID       uuid.UUID       `gorm:"type:varchar(255);not null" json:"role_id"`
	Role         Role            `gorm:"foreignKey:RoleID" json:"role"` // Relasi Belongs To

	// Relasi untuk Teacher (jika user adalah teacher)
	TeacherAssignments    []TeacherAssignment  `gorm:"foreignKey:TeacherID" json:"teacher_assignments,omitempty"`
	QuestionBanksAuthored []QuestionBank       `gorm:"foreignKey:CreatedByTeacherID" json:"question_banks_authored,omitempty"`
	QuestionsAuthored     []Question           `gorm:"foreignKey:CreatedByTeacherID" json:"questions_authored,omitempty"`
	ExamsAuthored         []Exam               `gorm:"foreignKey:CreatedByTeacherID" json:"exams_authored,omitempty"`
	GradedAttempts        []StudentExamAttempt `gorm:"foreignKey:GradedByTeacherID" json:"graded_attempts,omitempty"`

	// Relasi untuk Student (jika user adalah student)
	StudentEnrollments     []StudentEnrollment  `gorm:"foreignKey:StudentID" json:"student_enrollments,omitempty"`
	ExamAttempts           []StudentExamAttempt `gorm:"foreignKey:StudentID" json:"exam_attempts,omitempty"`
	StudentExamTokenUsages []ExamTokenUsage     `gorm:"foreignKey:StudentID" json:"-"`
}
