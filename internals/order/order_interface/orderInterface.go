package orderinterface

import (
	orderdto "github.com/ak-repo/ecommerce-gin/internals/order/order_dto"
	"github.com/ak-repo/ecommerce-gin/models"
)

type OrderServiceInterface interface {
	GetAllOrdersService() ([]orderdto.AdminOrderResponse, error)
	GetOrderByIDService(id uint) (*orderdto.AdminOrderResponse, error)
	UpdateOrderStatusService(order *orderdto.AdminUpdateOrderStatusRequest) error

	GetCustomerOrdersService(userID uint) (*orderdto.CustomerOrderListResponse, error)
	GetCustomerOrderbyOrderIDService(orderID uint) (*orderdto.CustomerOrderDetailResponse, error)
}

type OrderRepoInterface interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(orderID uint) (*models.Order, error)
	OrderStatusUpdate(order *orderdto.AdminUpdateOrderStatusRequest) error

	GetOrdersByUserID(userID uint) ([]models.Order, error)
}
