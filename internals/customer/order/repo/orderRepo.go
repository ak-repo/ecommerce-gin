package orderrepo

import (
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/customer/order/order_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) orderinterface.Repository {
	return &repository{DB: db}
}

func (r *repository) GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order

	if err := r.DB.
		Preload("OrderItems").
		Preload("ShippingAddress").
		Preload("Payment").
		Preload("OrderItems.Product").
		Preload("CancelRequest").
		Preload("User").
		First(&order, "id = ?", orderID).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *repository) GetOrderByCustomerID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("Payment").
		Find(&orders, "user_id=?", userID).Error
	return orders, err
}

func (r *repository) CancelOrder(cancel *models.OrderCancelRequest) error {
	return r.DB.Create(cancel).Error
}

func (r *repository) CancellationResponse(orderID uint) (*models.OrderCancelRequest, error) {

	var order models.OrderCancelRequest
	err := r.DB.Where("order_id=?", orderID).First(&order).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &order, err
}
