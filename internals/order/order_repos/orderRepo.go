package orderrepos

import (
	orderdto "github.com/ak-repo/ecommerce-gin/internals/order/order_dto"
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
	if err := r.DB.Preload("OrderItems").Preload("ShippingAddress").Preload("Payment").Find(&orders).Error; err != nil {
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
		Preload("Payment").
		Preload("OrderItems.Product").
		First(&order, "id = ?", orderID).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

// order status change
func (r *OrderRepo) OrderStatusUpdate(order *orderdto.AdminUpdateOrderStatusRequest) error {
	return r.DB.Model(&models.Order{}).Where("id=?", order.OrderID).Update("status", order.Status).Error
}

// get orders by users id
func (r *OrderRepo) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("OrderItems").Preload("Payment").Find(&orders, "user_id=?", userID).Error
	return orders, err
}
