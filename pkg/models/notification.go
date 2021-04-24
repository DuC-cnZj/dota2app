package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID int `json:"id" gorm:"primaryKey;"`

	UserID int        `json:"user_id" gorm:"not null;"`
	Data   string     `json:"data" gorm:"type:text;not null;"`
	ReadAt *time.Time `json:"read_at"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
