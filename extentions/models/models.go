package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string    `gorm:"type:varchar(255);primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Student struct {
	Base
	NIS      string `gorm:"type:varchar(50);not null;unique" json:"nis"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	// Hash     string          `gorm:"type:varchar(255);not null" json:"hash"`
	Profile json.RawMessage `gorm:"type:json" json:"profile"`

	Roles []HasRole `gorm:"polymorphic:Owner;polymorphicValue:student" json:"roles,omitempty"`
}

type Teacher struct {
	Base
	NIK      string          `gorm:"type:varchar(50);not null;unique" json:"nik"`
	Password string          `gorm:"type:varchar(255);not null" json:"-"`
	Profile  json.RawMessage `gorm:"type:json" json:"profile"`
	// Hash     string          `gorm:"type:varchar(255);not null" json:"hash"`

	Roles []HasRole `gorm:"polymorphic:Owner;polymorphicValue:teacher" json:"roles,omitempty"`
}

type Subject struct {
	Base
	SubjectName string `gorm:"type:varchar(255);not null" json:"subject_name"`
	Description string `gorm:"type:text;null" json:"description,omitempty"`
	ClassID     string `gorm:"type:varchar(255);not null" json:"class_id"`

	Class Class `gorm:"foreignKey:ClassID" json:"class,omitempty"`
}

type Role struct {
	Base
	RoleName string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"role_name"` // cth: 'admin', 'teacher', 'student'
	HasRole  []HasRole `gorm:"foreignKey:RoleID" json:"has_role,omitempty"`
}

type Class struct {
	Base
	ClassName   string    `gorm:"type:varchar(100);not null" json:"class_name"`
	GradeLevel  string    `gorm:"type:varchar(50);null" json:"grade_level,omitempty"` // cth: '10', '11', '12'
	Description string    `gorm:"type:text;null" json:"description,omitempty"`
	Subject     []Subject `gorm:"foreignKey:ClassID" json:"-"`
}

// HasRole diubah secara signifikan
type HasRole struct {
	ID        uuid.UUID `gorm:"type:varchar(255);primary_key" json:"id"`
	RoleID    string    `gorm:"type:varchar(255);not null" json:"role_id"`
	OwnerID   string    `gorm:"type:varchar(255);not null" json:"owner_id"`  // Tipe disamakan dengan ID di Base
	OwnerType string    `gorm:"type:varchar(50);not null" json:"owner_type"` // "student" atau "teacher"
	CreatedAt time.Time
	UpdatedAt time.Time

	Role Role `gorm:"foreignKey:RoleID" json:"role"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return
}

func (Subject) TableName() string {
	return "main.subjects"
}

func (Student) TableName() string {
	return "main.students"
}
func (Class) TableName() string {
	return "main.classes"
}
func (Role) TableName() string {
	return "main.roles"
}

func (HasRole) TableName() string {
	return "main.has_roles"
}

func (Teacher) TableName() string {
	return "main.teachers"
}
