package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base berisi field standar untuk semua model
type Base struct {
	ID        uuid.UUID `gorm:"type:varchar(255);primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return
}
