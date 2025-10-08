package orderservices

import (
	"errors"

	orderdto "github.com/ak-repo/ecommerce-gin/internals/order/order_dto"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/order/order_interface"
	"github.com/ak-repo/ecommerce-gin/models"
)

type OrderService struct {
	OrderRepo orderinterface.OrderRepoInterface
}

func NewOrderService(orderRepo orderinterface.OrderRepoInterface) orderinterface.OrderServiceInterface {
	return &OrderService{OrderRepo: orderRepo}
}

// Get All Orders
func (s *OrderService) GetAllOrdersService() ([]orderdto.AdminOrderResponse, error) {

	data, err := s.OrderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("orders not found")
	}

	var orders []orderdto.AdminOrderResponse
	for _, val := range data {
		// each products
		var items []orderdto.AdminOrderItemResp
		for _, i := range val.OrderItems {
			item := orderdto.AdminOrderItemResp{
				ProductID: i.ProductID,
				Quantity:  i.Quantity,
				UnitPrice: i.UnitPrice,
				Subtotal:  i.UnitPrice * float64(i.Quantity),
			}
			items = append(items, item)

		}
		order := orderdto.AdminOrderResponse{
			OrderID:     val.ID,
			UserID:      val.UserID,
			OrderDate:   val.CreatedAt,
			Status:      val.Status,
			TotalAmount: val.TotalAmount,
			UserEmail:   val.User.Email,
			Address: orderdto.AdminOrderAddressDTO{
				ID:      val.ShippingAddress.ID,
				Street:  val.ShippingAddress.AddressLine,
				City:    val.ShippingAddress.City,
				State:   val.ShippingAddress.State,
				ZipCode: val.ShippingAddress.PostalCode,
				Country: val.ShippingAddress.Country,
			},
			Items: items,
			Payment: orderdto.AdminOrderPaymentDTO{
				PaymentID: val.Payment.ID,
				Method:    val.Payment.PaymentMethod,
				Amount:    val.Payment.Amount,
				Status:    val.Payment.Status,
			},
		}

		orders = append(orders, order)
	}

	return orders, nil

}

// Get OrderBy ID  => for admin
func (s *OrderService) GetOrderByIDService(id uint) (*orderdto.AdminOrderResponse, error) {
	data, err := s.OrderRepo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("order not found")
	}

	order := orderdto.AdminOrderResponse{
		OrderID:     data.ID,
		UserID:      data.UserID,
		OrderDate:   data.CreatedAt,
		Status:      data.Status,
		TotalAmount: data.TotalAmount,
		UserEmail:   data.User.Email,
		Address: orderdto.AdminOrderAddressDTO{
			ID:      data.ShippingAddress.ID,
			Street:  data.ShippingAddress.AddressLine,
			City:    data.ShippingAddress.City,
			State:   data.ShippingAddress.State,
			ZipCode: data.ShippingAddress.PostalCode,
			Country: data.ShippingAddress.Country,
		},
		Payment: orderdto.AdminOrderPaymentDTO{
			PaymentID: data.Payment.ID,
			Method:    data.Payment.PaymentMethod,
			Amount:    data.Payment.Amount,
			Status:    data.Payment.Status,
		},
	}
	var items []orderdto.AdminOrderItemResp
	for _, i := range data.OrderItems {
		item := orderdto.AdminOrderItemResp{
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			UnitPrice: i.UnitPrice,
			Subtotal:  i.UnitPrice * float64(i.Quantity),
		}
		items = append(items, item)

	}
	order.Items = items

	return &order, nil
}

// UpdateOrderStatus
func (s *OrderService) UpdateOrderStatusService(req *orderdto.AdminUpdateOrderStatusRequest) error {

	validStatuses := map[string]bool{
		"pending": true, "accepted": true, "confirmed": true,
		"shipped": true, "delivered": true, "cancelled": true,
		"refunded": true, "completed": true,
	}
	if !validStatuses[req.Status] {
		return errors.New("invalid status")
	}
	return s.OrderRepo.OrderStatusUpdate(req)

}

// list customer order cancellation requeat for admin
func (s *OrderService) OrderCancellationReqListingService() ([]orderdto.AdminCancelRequestResponse, error) {

	data, err := s.OrderRepo.GetAllCancelRequest()
	if err != nil {
		return nil, err
	}

	var requests []orderdto.AdminCancelRequestResponse
	for _, o := range data {
		req := orderdto.AdminCancelRequestResponse{
			ID:          o.ID,
			OrderID:     o.OrderID,
			UserID:      o.UserID,
			Customer:    o.User.Username,
			OrderStatus: o.Order.Status,
			Status:      o.Status,
			Reason:      o.Reason,
			CreatedAt:   o.CreatedAt,
		}
		requests = append(requests, req)
	}

	return requests, nil
}

func (s *OrderService) AcceptCancellationReqService(reqID uint) error {
	orderID, err := s.OrderRepo.AcceptOrderCancellationReq(reqID)
	if err != nil {
		return err
	}

	status := orderdto.AdminUpdateOrderStatusRequest{
		OrderID: orderID,
		Status:  "cancelled",
	}
	if err := s.OrderRepo.OrderStatusUpdate(&status); err != nil {
		return err
	}

	return nil

}

// reject order  cancellation req
func (s *OrderService) RejectCancellationReqService(reqID uint) error {
	orderID, err := s.OrderRepo.RejectOrderCancellationReq(reqID)
	if err != nil {
		return err
	}

	status := orderdto.AdminUpdateOrderStatusRequest{
		OrderID: orderID,
		Status:  "confirmed",
	}
	if err := s.OrderRepo.OrderStatusUpdate(&status); err != nil {
		return err
	}
	return nil

}

// ----------------------------------------------------------------for customers----------------------------------------------------------
// GetCustomerOrders
func (s *OrderService) GetCustomerOrdersService(userID uint) (*orderdto.CustomerOrderListResponse, error) {

	data, err := s.OrderRepo.GetOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("no data found in this user id")
	}

	var orders []orderdto.CustomerOrderSummary
	for _, i := range data {
		order := orderdto.CustomerOrderSummary{
			OrderID:     i.ID,
			OrderDate:   i.OrderDate,
			Status:      i.Status,
			TotalAmount: i.TotalAmount,
			PaymentMode: i.Payment.PaymentMethod,
		}

		orders = append(orders, order)

	}

	return &orderdto.CustomerOrderListResponse{Orders: orders}, nil
}

func (s *OrderService) GetCustomerOrderbyOrderIDService(orderID uint) (*orderdto.CustomerOrderDetailResponse, error) {
	data, err := s.OrderRepo.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("order not found")
	}

	order := orderdto.CustomerOrderDetailResponse{
		OrderID:     data.ID,
		OrderDate:   data.OrderDate,
		Status:      data.Status,
		TotalAmount: data.TotalAmount,
		Address: orderdto.CustomerOrderAddressDTO{
			Street:  data.ShippingAddress.AddressLine,
			City:    data.ShippingAddress.City,
			State:   data.ShippingAddress.State,
			ZipCode: data.ShippingAddress.PostalCode,
			Country: data.ShippingAddress.Country,
		},
		Payment: orderdto.CustomerOrderPaymentDTO{
			Method: data.Payment.PaymentMethod,
			Amount: data.Payment.Amount,
			Status: data.Payment.Status,
		},
	}

	if data.CancelRequest != nil {
		order.CancelReq = &orderdto.CustomerCancelRequestResponse{
			ID:      data.CancelRequest.ID,
			OrderID: data.CancelRequest.OrderID,
			Status:  data.CancelRequest.Status,
			Reason:  data.CancelRequest.Reason,
		}
	}

	var items []orderdto.CustomerOrderItemResp
	for _, i := range data.OrderItems {
		item := orderdto.CustomerOrderItemResp{
			ProductID: i.ProductID,
			Name:      i.Product.Title,
			Quantity:  i.Quantity,
			Price:     i.UnitPrice,
			Subtotal:  float64(i.Quantity) * i.UnitPrice,
			ImageURL:  i.Product.ImageURL,
		}
		items = append(items, item)
	}
	order.Items = items

	return &order, nil
}

// order cancellation service for customer
func (s *OrderService) CancelOrderByCustomerService(req *orderdto.CreateCancelRequest, userID uint) error {
	cancelOrder := models.OrderCancelRequest{
		OrderID: req.OrderID,
		UserID:  userID,
		Reason:  req.Reason,
	}

	return s.OrderRepo.CancellationByCustomer(&cancelOrder)
}

// order cancellation req-response for customer
func (s *OrderService) CancellationResponseForCustomerService(orderID uint) (*orderdto.CustomerCancelRequestResponse, error) {
	data, err := s.OrderRepo.CancellationResponseToCustomer(orderID)
	if err != nil {
		return nil, err
	}
	response := orderdto.CustomerCancelRequestResponse{
		ID:      data.ID,
		Reason:  data.Reason,
		Status:  data.Status,
		OrderID: data.OrderID,
	}
	return &response, nil
}
