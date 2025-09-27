package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"size:255;uniqueIndex;not null"`
	PasswordHash string `gorm:"size:255;not null" `
	// Username     string `gorm:"size:255;not null" `
	IsActive bool   `gorm:"default:true"`
	Role     string `gorm:"size:100;not null"` // condition need to given
	// Addresses    []Address
	// Orders       []Order
	// Cart         Cart
	// Wishlist     Wishlist
}

type InputUser struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8"`
}
