package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string    `gorm:"size:255;uniqueIndex;not null"`
	Products []Product `gorm:"foreignKey:CategoryID"` // One-to-many relationship
}

type Product struct {
	ID        string `gorm:"primaryKey;size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Product Catalog Information
	Title       string  `gorm:"size:255;not null"`
	Description string  `gorm:"type:text"`
	SKU         string  `gorm:"size:100;uniqueIndex"`
	BasePrice   float64 `gorm:"type:numeric(10,2);not null"`
	Stock       int     `gorm:"not null;default:0"` // Added quantity/stock
	ImageURL    string  `gorm:"type:text"`          // From 'image'

	// Foreign Key for Category
	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Active flag
	IsActive bool `gorm:"default:true"`
}

type UpdateProductInput struct {
	Title       *string  `form:"title" json:"title,omitempty"`
	Description *string  `form:"description" json:"description,omitempty"`
	CategoryID  *uint    `form:"category_id" json:"category_id,omitempty"`
	BasePrice   *float64 `form:"price" json:"price,omitempty"`
	Stock       *int     `form:"stock" json:"stock,omitempty"`
	ImageURL    *string  `form:"image_url" json:"image_url,omitempty"`
	IsActive    *bool    `form:"is_active" json:"is_active,omitempty"`
}
