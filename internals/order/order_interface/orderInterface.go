package orderinterface

import (
	orderdto "github.com/ak-repo/ecommerce-gin/internals/order/order_dto"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllOrders(ctx *gin.Context)
	GetOrderByIDForCustomer(ctx *gin.Context)
	GetOrderByIDForAdmin(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	GetOrderByCustomerIDForCustomer(ctx *gin.Context)
	// GetOrderByCustomerIDForAdmin(ctx *gin.Context)

	// cancellation
	GetAllCancels(ctx *gin.Context)
	AcceptCancel(ctx *gin.Context)
	RejectCancel(ctx *gin.Context)
	CancelOrder(ctx *gin.Context)
	CancellationResponse(ctx *gin.Context)
}

type Service interface {
	GetAllOrders() ([]orderdto.AllOrderResponse, error)
	GetOrderByID(id uint) (*orderdto.OrderDetailResponse, error)
	UpdateStatus(order *orderdto.AdminUpdateOrderStatusRequest) error

	GetOrderByCustomerID(userID uint) ([]orderdto.CustomerOrder, error)

	// cancellation
	CancelOrder(req *orderdto.CreateCancelRequest, userID uint) error
	CancellationResponse(orderID uint) (*orderdto.CancelRequestStatus, error)
	GetAllCancels() ([]orderdto.AdminCancelRequestResponse, error)
	AcceptCancel(reqID uint) error
	RejectCancel(reqID uint) error
}

type Repository interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(orderID uint) (*models.Order, error)
	UpdateStatus(order *orderdto.AdminUpdateOrderStatusRequest) error
	GetOrderByCustomerID(userID uint) ([]models.Order, error)

	// cancellation
	CancelOrder(cancel *models.OrderCancelRequest) error
	CancellationResponse(orderID uint) (*models.OrderCancelRequest, error)
	GetAllCancels() ([]models.OrderCancelRequest, error)
	AcceptCancel(reqID uint) (uint, error)
	RejectCancel(reqID uint) (uint, error)
}
