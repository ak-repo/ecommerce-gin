package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID            uint      `gorm:"not null"`
	OrderDate         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Status            string    `gorm:"size:50;not null"`
	TotalAmount       float64   `gorm:"type:numeric(10,2);not null"`
	ShippingAddressID uint
	ShippingAddress   Address     `gorm:"foreignKey:ShippingAddressID"`
	OrderItems        []OrderItem `gorm:"foreignKey:OrderID"`
	Payment           Payment
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	UnitPrice float64 `gorm:"type:numeric(10,2);not null"`
}
