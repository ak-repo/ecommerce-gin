package custcheckoutinterface

import (
	custcheckout "github.com/ak-repo/ecommerce-gin/internals/customer/cust_checkout"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	CustomerShowCheckoutHandler(ctx *gin.Context)
	CustomerCheckoutHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	CheckoutSummaryService(userID uint) (*custcheckout.CheckoutSummaryResponse, error)
	ProcessCheckoutService(req *custcheckout.CheckoutRequest) (*custcheckout.CheckoutResponse, error)
}

type RepoInterface interface {
	OrderCreation(order *models.Order) error
	OrderItemsCreation(orderItems []models.OrderItem) error
	PaymentCreation(payment *models.Payment) error
}
