package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID int `json:"id" gorm:"primaryKey;"`

	Name     string `json:"name" gorm:"type:varchar(80);not null;default:'';"`
	Email    string `json:"email" gorm:"type:varchar(80);uniqueIndex:uniq_email;not null;"`
	Password string `json:"-" gorm:"type:varchar(255);not null;default:'';"`
	Mobile   string `json:"mobile" gorm:"type:varchar(40);"`
	Avatar   string `json:"avatar" gorm:"type:text;"`
	Intro    string `json:"intro" gorm:"type:text;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
