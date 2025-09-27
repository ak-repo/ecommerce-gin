package models

import (
	"time"

	"gorm.io/gorm"
)


type Address struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	AddressLine string `gorm:"type:text;not null"`
	City        string `gorm:"size:100;not null"`
	State       string `gorm:"size:100;not null"`
	PostalCode  string `gorm:"size:20;not null"`
	Country     string `gorm:"size:100;not null"`
}



type Order struct {
	gorm.Model
	UserID            uint      `gorm:"not null"`
	OrderDate         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Status            string    `gorm:"size:50;not null"`
	TotalAmount       float64   `gorm:"type:numeric(10,2);not null"`
	ShippingAddressID uint
	ShippingAddress   Address `gorm:"foreignKey:ShippingAddressID"`
	OrderItems        []OrderItem
	Payment           Payment
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	UnitPrice float64 `gorm:"type:numeric(10,2);not null"`
}

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
