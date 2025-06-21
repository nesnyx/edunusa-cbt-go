package models

type Class struct {
	Base
	ClassName          string              `gorm:"type:varchar(100);not null" json:"class_name"`
	GradeLevel         string              `gorm:"type:varchar(50);null" json:"grade_level,omitempty"` // cth: '10', '11', '12'
	Description        string              `gorm:"type:text;null" json:"description,omitempty"`
	TeacherAssignments []TeacherAssignment `gorm:"foreignKey:ClassID" json:"-"`
	StudentEnrollments []StudentEnrollment `gorm:"foreignKey:ClassID" json:"-"`
	Exams              []Exam              `gorm:"foreignKey:ClassID" json:"-"`
}
