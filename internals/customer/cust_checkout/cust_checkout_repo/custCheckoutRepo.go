package custcheckoutrepo

import (
	custcheckoutinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_checkout/cust_checkout_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type CustomerCheckoutRepo struct {
	DB *gorm.DB
}

func NewCustomerCheckoutRepo(db *gorm.DB) custcheckoutinterface.RepoInterface {
	return &CustomerCheckoutRepo{DB: db}
}

// order creation
func (r *CustomerCheckoutRepo) OrderCreation(order *models.Order) error {
	return r.DB.Create(&order).Error
}

// order items add
func (r *CustomerCheckoutRepo) OrderItemsCreation(orderItems []models.OrderItem) error {
	return r.DB.Create(&orderItems).Error
}

// payment
func (r *CustomerCheckoutRepo) PaymentCreation(payment *models.Payment) error {
	return r.DB.Create(&payment).Error
}
