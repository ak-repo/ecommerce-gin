package models

import (
	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	UserID        uint `gorm:"uniqueIndex;not null"`
	User          User `gorm:"foreignKey:UserID"`
	WishlistItems []WishlistItem
}

type WishlistItem struct {
	gorm.Model
	WishlistID uint    `gorm:"not null"`
	ProductID  uint    `gorm:"not null"`
	Product    Product `gorm:"foreignKey:ProductID"`
}
