package orderinterface

import (
	orderdto "github.com/ak-repo/ecommerce-gin/internals/order/order_dto"
	"github.com/ak-repo/ecommerce-gin/models"
)

type OrderServiceInterface interface {
	// admin
	GetAllOrdersService() ([]orderdto.AdminOrderResponse, error)
	GetOrderByIDService(id uint) (*orderdto.AdminOrderResponse, error)
}

type OrderRepoInterface interface {
	// admin
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(orderID uint) (*models.Order, error)
}
