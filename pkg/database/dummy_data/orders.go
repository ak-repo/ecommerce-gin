package dummydata

import (
	"time"

	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

func SeedOrders(db *gorm.DB) {
	orders := []models.Order{
		{
			UserID:            1,
			OrderDate:         time.Now().AddDate(0, 0, -2),
			Status:            "Delivered",
			TotalAmount:       84.00 * 2, // Example
			ShippingAddressID: 1,         // assuming address exists
			OrderItems: []models.OrderItem{
				{ProductID: 1, Quantity: 1, UnitPrice: 84.00},
				{ProductID: 2, Quantity: 1, UnitPrice: 84.00},
			},
		},
		{
			UserID:            1,
			OrderDate:         time.Now().AddDate(0, 0, -1),
			Status:            "Shipped",
			TotalAmount:       84.00 * 3,
			ShippingAddressID: 1,
			OrderItems: []models.OrderItem{
				{ProductID: 3, Quantity: 1, UnitPrice: 84.00},
				{ProductID: 4, Quantity: 2, UnitPrice: 84.00},
			},
		},
		{
			UserID:            1,
			OrderDate:         time.Now(),
			Status:            "Pending",
			TotalAmount:       107.00 * 1,
			ShippingAddressID: 1,
			OrderItems: []models.OrderItem{
				{ProductID: 6, Quantity: 1, UnitPrice: 107.00},
			},
		},
	}

	for _, order := range orders {
		if err := db.Create(&order).Error; err != nil {
			panic(err)
		}
	}
}
