package models

import (
	"time"

	"gorm.io/gorm"
)

type EmailOTP struct {
	gorm.Model
	Email     string    `gorm:"index;not null"`
	CodeHash  string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"index"`
	Used      bool
}
