package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID       uint      `gorm:"uniqueIndex;not null"`
	PaymentMethod string    `gorm:"size:50;not null"`
	Amount        float64   `gorm:"type:numeric(10,2);not null"`
	Status        string    `gorm:"size:50;not null"`
	TransactionID string    `gorm:"size:255"`
	PaymentDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
