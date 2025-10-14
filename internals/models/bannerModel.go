package models

import (
	"gorm.io/gorm"
)

// -----------------
// Banner Model
// -----------------
type Banner struct {
	gorm.Model
	Title       string `gorm:"size:255;not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	ImageURL    string `gorm:"size:512;not null" json:"image_url"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
}
