package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base berisi field standar untuk semua model
type Base struct {
	ID        string    `gorm:"type:varchar(255);primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return
}
