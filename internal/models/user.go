package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Student struct {
	Base
	NIS      string          `gorm:"type:varchar(50);not null;unique" json:"nis"`
	Password string          `gorm:"type:varchar(255);not null" json:"-"`
	Profile  json.RawMessage `gorm:"type:json" json:"profile"`

	Roles []HasRole `gorm:"polymorphic:Owner;polymorphicValue:student" json:"roles,omitempty"`
}

type Teacher struct {
	Base
	NIK      string          `gorm:"type:varchar(50);not null;unique" json:"nik"`
	Password string          `gorm:"type:varchar(255);not null" json:"-"`
	Profile  json.RawMessage `gorm:"type:json" json:"profile"`

	Roles []HasRole `gorm:"polymorphic:Owner;polymorphicValue:teacher" json:"roles,omitempty"`
}

type HasRole struct {
	ID        uuid.UUID `gorm:"type:varchar(255);primary_key" json:"id"`
	RoleID    string    `gorm:"type:varchar(255);not null" json:"role_id"`
	OwnerID   string    `gorm:"type:varchar(255);not null" json:"owner_id"`
	OwnerType string    `gorm:"type:varchar(50);not null" json:"owner_type"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Role Role `gorm:"foreignKey:RoleID" json:"role"`
}

func (Student) TableName() string {
	return "main.students"
}

func (Teacher) TableName() string {
	return "main.teachers"
}

func (HasRole) TableName() string {
	return "main.has_roles"
}
