package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string  `gorm:"size:255;not null"`
	Description   string  `gorm:"type:text"`
	Price         float64 `gorm:"type:numeric(10,2);not null"`
	StockQuantity int     `gorm:"not null"`
	SKU           string  `gorm:"size:100;uniqueIndex;not null"`

	CategoryID uint     `gorm:"not null"`                                       // foreign key
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // GORM relation
}

type Category struct {
	gorm.Model
	Name     string    `gorm:"size:255;uniqueIndex;not null"` // Category name must be unique
	Products []Product // GORM can use this for reverse relationship (optional)
}
