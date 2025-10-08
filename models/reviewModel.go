package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ProductID uint    `gorm:"not null"`
	UserID    uint    `gorm:"not null"`
	Rating    uint    `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment   string  `gorm:"type:text"`
	Status    string  `gorm:"type:varchar(20);default:'PENDING';check:status IN ('PENDING','APPROVED','REJECTED')"`
	Product   Product `gorm:"foreignKey:ProductID"`
	User      User    `gorm:"foreignKey:UserID"`
}
