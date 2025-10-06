package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint       `gorm:"not null"`          // FK to User
	User      User       `gorm:"foreignKey:UserID"` // Association
	CartItems []CartItem `gorm:"constraint:OnDelete:CASCADE;foreignKey:CartID"`
}

// Each CartItem links a product to a cart
type CartItem struct {
	gorm.Model
	CartID    uint    `gorm:"not null;index;constraint:OnDelete:CASCADE;"`
	ProductID uint    `gorm:"not null"` // FK to Product
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null;default:1"`
	Price     float64 `gorm:"type:numeric(10,2)"` // store price at time of adding
	Subtotal  float64 `gorm:"type:numeric(10,2)"` // Quantity * Price
}
