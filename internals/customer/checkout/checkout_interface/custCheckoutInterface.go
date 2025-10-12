package checkoutinterface

import (
	checkoutdto "github.com/ak-repo/ecommerce-gin/internals/customer/checkout/dto"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CheckoutSummary(ctx *gin.Context)
	ProcessCheckout(ctx *gin.Context)
}

type Service interface {
	CheckoutSummary(userID uint) (*checkoutdto.CheckoutSummaryResponse, error)
	ProcessCheckout(req *checkoutdto.CheckoutRequest) (*checkoutdto.CheckoutResponse, error)
}

type Repository interface {
	OrderCreation(order *models.Order) error
	OrderItemsCreation(orderItems []models.OrderItem) error
	PaymentCreation(payment *models.Payment) error
}
