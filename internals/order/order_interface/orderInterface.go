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

	// cancellation
	CancelOrderByCustomerService(req *orderdto.CreateCancelRequest, userID uint) error
	CancellationResponseForCustomerService(orderID uint) (*orderdto.CustomerCancelRequestResponse, error)
	OrderCancellationReqListingService() ([]orderdto.AdminCancelRequestResponse, error)
	AcceptCancellationReqService(reqID uint) error
	RejectCancellationReqService(reqID uint) error
}

type OrderRepoInterface interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(orderID uint) (*models.Order, error)
	OrderStatusUpdate(order *orderdto.AdminUpdateOrderStatusRequest) error
	GetOrdersByUserID(userID uint) ([]models.Order, error)

	// cancellation
	CancellationByCustomer(cancel *models.OrderCancelRequest) error
	CancellationResponseToCustomer(orderID uint) (*models.OrderCancelRequest, error)
	GetAllCancelRequest() ([]models.OrderCancelRequest, error)
	AcceptOrderCancellationReq(reqID uint) (uint, error)
	RejectOrderCancellationReq(reqID uint) (uint, error)
}
