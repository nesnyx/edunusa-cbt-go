package models

import "github.com/google/uuid"

// TeacherAssignment merepresentasikan tabel 'teacher_assignments'
type TeacherAssignment struct {
	Base
	TeacherID    uuid.UUID `gorm:"type:varchar(255);not null;index:idx_teacher_class_subject,unique" json:"teacher_id"`
	ClassID      uuid.UUID `gorm:"type:varchar(255);not null;index:idx_teacher_class_subject,unique" json:"class_id"`
	SubjectID    uuid.UUID `gorm:"type:varchar(255);not null;index:idx_teacher_class_subject,unique" json:"subject_id"`
	AcademicYear string    `gorm:"type:varchar(50);null;index:idx_teacher_class_subject,unique" json:"academic_year,omitempty"` // cth: '2024/2025'

	// Teacher User    `gorm:"foreignKey:TeacherID" json:"teacher,omitempty"`
	// Class   Class   `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	Subject Subject `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
}
