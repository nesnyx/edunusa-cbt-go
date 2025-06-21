package database

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Define Model
type Users struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	Username     string
	PasswordHash string
	Profile      json.RawMessage
	RoleID       Roles `gorm:"foreignKey:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Roles struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	RoleName string
}

type Subjects struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	SubjectName string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Classes struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	ClassName   string
	GradeLevel  string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TeacherAssignments struct {
	gorm.Model
	ID        uint  `gorm:"primaryKey;autoIncrement"`
	TeacherID Users `gorm:"foreignKey:ID"`
}
