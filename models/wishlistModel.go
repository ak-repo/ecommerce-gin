package models

import (
	"gorm.io/gorm"
)

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
