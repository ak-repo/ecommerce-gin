package orderinterface

import (
	orderdto "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllOrders(ctx *gin.Context)
	GetOrderByID(ctx *gin.Context)
	GetOrderByCustomerID(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	GetAllCancels(ctx *gin.Context)
	AcceptCancel(ctx *gin.Context)
	RejectCancel(ctx *gin.Context)
}

type Service interface {
	GetAllOrders() ([]orderdto.AllOrderResponse, error)
	GetOrderByID(id uint) (*orderdto.OrderDetailResponse, error)
	GetOrderByCustomerID(userID uint) ([]orderdto.CustomerOrder, error)
	UpdateStatus(req *orderdto.AdminUpdateOrderStatusRequest) error
	GetAllCancels() ([]orderdto.AdminCancelRequestResponse, error)
	AcceptCancel(reqID uint) error
	RejectCancel(reqID uint) error
}

type Repository interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(orderID uint) (*models.Order, error)
	GetOrderByCustomerID(userID uint) ([]models.Order, error)
	UpdateStatus(order *orderdto.AdminUpdateOrderStatusRequest) error
	GetAllCancels() ([]models.OrderCancelRequest, error)
	AcceptCancel(reqID uint) (uint, error)
	RejectCancel(reqID uint) (uint, error)
}
