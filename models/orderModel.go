package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID            uint
	User              User      `gorm:"foreignKey:UserID"`
	OrderDate         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Status            string    `gorm:"size:50;not null"`
	TotalAmount       float64   `gorm:"type:numeric(10,2);not null"`
	ShippingAddressID uint
	ShippingAddress   Address     `gorm:"foreignKey:ShippingAddressID"`
	OrderItems        []OrderItem `gorm:"foreignKey:OrderID"`
	Payment           Payment
	CancelRequest     *OrderCancelRequest `gorm:"foreignKey:OrderID" json:"cancel_request,omitempty"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	UnitPrice float64 `gorm:"type:numeric(10,2);not null"`
}

type OrderCancelRequest struct {
	gorm.Model
	OrderID uint   `gorm:"not null" json:"order_id"`
	Order   Order  `gorm:"foreignKey:OrderID"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
	Reason  string `json:"reason"`
	Status  string `gorm:"type:varchar(20);default:'PENDING'" json:"status"`
}
