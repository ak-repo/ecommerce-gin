package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"size:255;uniqueIndex;not null"  json:"email"`
	PasswordHash string `gorm:"size:255;not null" json:"-"`
	IsActive     bool   `gorm:"default:true"`
	Role         string `gorm:"size:100;not null" json:"role"`
	Addresses    []Address 
	Orders       []Order
	Cart         Cart
	Wishlist     Wishlist
}

type InputUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
