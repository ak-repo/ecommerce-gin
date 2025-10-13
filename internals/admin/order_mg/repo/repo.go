package orderrepo

import (
	orderdto "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/dto"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/order_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewOrderRepositoryMG(db *gorm.DB) orderinterface.Repository {
	return &repository{DB: db}
}

func (r *repository) GetAllOrders() ([]models.Order, error) {
	orders := []models.Order{}
	if err := r.DB.
		Preload("OrderItems").
		Preload("ShippingAddress").
		Preload("Payment").
		Preload("User").
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
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

func (r *repository) UpdateStatus(order *orderdto.AdminUpdateOrderStatusRequest) error {
	return r.DB.Model(&models.Order{}).
		Where("id=?", order.OrderID).
		Update("status", order.Status).Error
}

func (r *repository) GetAllCancels() ([]models.OrderCancelRequest, error) {
	var cancelOrders []models.OrderCancelRequest
	err := r.DB.
		Preload("User").
		Preload("Order").
		Find(&cancelOrders).Error
	return cancelOrders, err
}

func (r *repository) AcceptCancel(reqID uint) (uint, error) {
	var orderID uint
	err := r.DB.Raw(`
		UPDATE order_cancel_requests
		SET status = 'APPROVED'
		WHERE id = ?
		RETURNING order_id
	`, reqID).Scan(&orderID).Error
	return orderID, err
}

func (r *repository) RejectCancel(reqID uint) (uint, error) {
	var orderID uint
	err := r.DB.Raw(`
	UPDATE order_cancel_requests 
	SET status = 'REJECTED'
	WHERE id = ?
	RETURNING order_id`, reqID).Scan(&orderID).Error
	return orderID, err
}
