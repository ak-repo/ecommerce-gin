package orderrepository

import (
	"strconv"

	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type OrderRepo interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id string) (*models.Order, error)
}

type orderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{DB: db}
}

// All orders
func (r *orderRepo) GetAllOrders() ([]models.Order, error) {
	orders := []models.Order{}
	if err := r.DB.Preload("OrderItems").Preload("ShippingAddress").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// Get order by order_ID
func (r *orderRepo) GetOrderByID(id string) (*models.Order, error) {
	var order models.Order

	idUint, _ := strconv.ParseUint(id, 10, 32)
	if err := r.DB.
		Preload("OrderItems").
		Preload("ShippingAddress").
		First(&order, "id = ?", uint(idUint)).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
