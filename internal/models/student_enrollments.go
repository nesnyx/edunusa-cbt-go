package models

import "github.com/google/uuid"

// StudentEnrollment merepresentasikan tabel 'student_enrollments'
type StudentEnrollment struct {
	Base
	StudentID    uuid.UUID `gorm:"type:varchar(255);not null;index:idx_student_class_year,unique" json:"student_id"`
	ClassID      uuid.UUID `gorm:"type:varchar(255);not null;index:idx_student_class_year,unique" json:"class_id"`
	AcademicYear string    `gorm:"type:varchar(50);null;index:idx_student_class_year,unique" json:"academic_year,omitempty"`

	Student User `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	// Class   Class `gorm:"foreignKey:ClassID" json:"class,omitempty"`
}
