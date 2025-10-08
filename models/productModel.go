package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string    `gorm:"size:255;uniqueIndex;not null"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Title         string  `gorm:"size:255;not null"`
	Description   string  `gorm:"type:text"`
	SKU           string  `gorm:"size:100;uniqueIndex"`
	BasePrice     float64 `gorm:"type:numeric(10,2);not null"`
	DiscountPrice float64 `gorm:"type:numeric(10,2);default:0"`
	Stock         int     `gorm:"not null;default:0"`
	ImageURL      string  `gorm:"type:text"`

	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	IsActive    bool `gorm:"default:true"`
	IsPublished bool `gorm:"default:false"`

	// --- Reviews ---
	Reviews     []Review `gorm:"foreignKey:ProductID"`
	AvgRating   float64  `gorm:"-:migration"` // calculated, ignored by GORM migrations
	ReviewCount int      `gorm:"-:migration"` // calculated
}
