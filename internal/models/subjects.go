package models


type Subject struct {
	Base
	SubjectName string `gorm:"type:varchar(255);not null" json:"subject_name"`
	Description string `gorm:"type:text;null" json:"description,omitempty"`

	TeacherAssignments []TeacherAssignment `gorm:"foreignKey:SubjectID" json:"-"`
	QuestionBanks      []QuestionBank      `gorm:"foreignKey:SubjectID" json:"-"`
	Exams              []Exam              `gorm:"foreignKey:SubjectID" json:"-"`
}

