package orderrepos

import (
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/order/order_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) orderinterface.OrderRepoInterface {
	return &OrderRepo{DB: db}
}

// Get All Orders
func (r *OrderRepo) GetAllOrders() ([]models.Order, error) {
	orders := []models.Order{}
	if err := r.DB.Preload("OrderItems").Preload("ShippingAddress").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrderByID
func (r *OrderRepo) GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order

	if err := r.DB.
		Preload("OrderItems").
		Preload("ShippingAddress").
		First(&order, "id = ?", orderID).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
