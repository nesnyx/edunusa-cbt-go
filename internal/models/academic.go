package models

type Role struct {
	Base
	RoleName string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"role_name"`
	HasRole  []HasRole `gorm:"foreignKey:RoleID" json:"has_role,omitempty"`
}

type Class struct {
	Base
	ClassName   string    `gorm:"type:varchar(100);not null" json:"class_name"`
	GradeLevel  string    `gorm:"type:varchar(50);null" json:"grade_level,omitempty"`
	Description string    `gorm:"type:text;null" json:"description,omitempty"`
	Subjects    []Subject `gorm:"foreignKey:ClassID" json:"subjects,omitempty"`
	Exams       []Exam    `gorm:"foreignKey:ClassID" json:"exams,omitempty"`
}

type Subject struct {
	Base
	SubjectName string `gorm:"type:varchar(255);not null" json:"subject_name"`
	Description string `gorm:"type:text;null" json:"description,omitempty"`
	ClassID     string `gorm:"type:varchar(255);not null" json:"class_id"`
	Class       Class  `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	Exams       []Exam `gorm:"foreignKey:SubjectID" json:"exams,omitempty"`
}

func (Role) TableName() string {
	return "main.roles"
}

func (Class) TableName() string {
	return "main.classes"
}

func (Subject) TableName() string {
	return "main.subjects"
}
