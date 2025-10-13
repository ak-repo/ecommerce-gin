package orderinterface

import (
	orderdto "github.com/ak-repo/ecommerce-gin/internals/customer/order/order_dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetOrderByID(ctx *gin.Context)
	GetOrderByCustomerID(ctx *gin.Context)

	// cancellation

	CancelOrder(ctx *gin.Context)
	CancellationResponse(ctx *gin.Context)
}

type Service interface {
	GetOrderByID(id uint) (*orderdto.OrderDetailResponse, error)

	GetOrderByCustomerID(userID uint) ([]orderdto.CustomerOrder, error)

	// cancellation
	CancelOrder(req *orderdto.CreateCancelRequest, userID uint) error
	CancellationResponse(orderID uint) (*orderdto.CancelRequestStatus, error)
}

type Repository interface {
	GetOrderByID(orderID uint) (*models.Order, error)
	GetOrderByCustomerID(userID uint) ([]models.Order, error)

	// cancellation
	CancelOrder(cancel *models.OrderCancelRequest) error
	CancellationResponse(orderID uint) (*models.OrderCancelRequest, error)
}
