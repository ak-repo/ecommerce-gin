package orderrepos

import (
	"errors"

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

// GetOrderByID
func (r *OrderRepo) GetOrderByID(orderID uint) (*models.Order, error) {
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

// order cancelled by customer
func (r *OrderRepo) CancellationByCustomer(cancel *models.OrderCancelRequest) error {
	return r.DB.Create(cancel).Error
}

// order cancellation response to customer
func (r *OrderRepo) CancellationResponseToCustomer(orderID uint) (*models.OrderCancelRequest, error) {

	var order models.OrderCancelRequest
	err := r.DB.Where("order_id=?", orderID).First(&order).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("no order cancel request for this order")
	}
	return &order, err
}

// GET all cancel request
func (r *OrderRepo) GetAllCancelRequest() ([]models.OrderCancelRequest, error) {
	var cancelOrders []models.OrderCancelRequest
	err := r.DB.
		Preload("User").
		Preload("Order").
		Find(&cancelOrders).Error
	return cancelOrders, err
}

// order cancellation accept
func (r *OrderRepo) AcceptOrderCancellationReq(reqID uint) (uint, error) {
	var orderID uint
	err := r.DB.Raw(`
		UPDATE order_cancel_requests
		SET status = 'APPROVED'
		WHERE id = ?
		RETURNING order_id
	`, reqID).Scan(&orderID).Error
	return orderID, err
}

// order cancellation reject

func (r *OrderRepo) RejectOrderCancellationReq(reqID uint) (uint, error) {
	var orderID uint
	err := r.DB.Raw(`
	UPDATE order_cancel_requests 
	SET status = 'REJECTED'
	WHERE id = ?
	RETURNING order_id`, reqID).Scan(&orderID).Error
	return orderID, err
}
