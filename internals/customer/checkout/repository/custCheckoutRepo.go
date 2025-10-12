package checkoutrepository

import (
	checkoutinterface "github.com/ak-repo/ecommerce-gin/internals/customer/checkout/checkout_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewCheckoutRepo(db *gorm.DB) checkoutinterface.Repository {
	return &repository{DB: db}
}

// order creation
func (r *repository) OrderCreation(order *models.Order) error {
	return r.DB.Create(&order).Error
}

// order items add
func (r *repository) OrderItemsCreation(orderItems []models.OrderItem) error {
	return r.DB.Create(&orderItems).Error
}

// payment
func (r *repository) PaymentCreation(payment *models.Payment) error {
	return r.DB.Create(&payment).Error
}
