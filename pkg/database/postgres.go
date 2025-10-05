package db

import (
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// New DB
func NewDB(dsn string) (*Database, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Address{}, &models.Order{}, &models.OrderItem{}, &models.Cart{}, &models.CartItem{}, &models.Payment{}); err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
