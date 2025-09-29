package models

import (
	"time"

	"gorm.io/gorm"
)



type Cart struct {
	gorm.Model
	UserID    uint `gorm:"uniqueIndex;not null"`
	CartItems []CartItem
}

type CartItem struct {
	gorm.Model
	CartID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"not null"`
}

type Wishlist struct {
	gorm.Model
	UserID        uint `gorm:"uniqueIndex;not null"`
	WishlistItems []WishlistItem
}

type WishlistItem struct {
	gorm.Model
	WishlistID uint `gorm:"not null"`
	ProductID  uint `gorm:"not null"`
}

type Payment struct {
	gorm.Model
	OrderID       uint      `gorm:"uniqueIndex;not null"`
	PaymentMethod string    `gorm:"size:50;not null"`
	Amount        float64   `gorm:"type:numeric(10,2);not null"`
	Status        string    `gorm:"size:50;not null"`
	TransactionID string    `gorm:"size:255"`
	PaymentDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
