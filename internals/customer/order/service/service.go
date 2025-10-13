package orderservice

import (
	"errors"

	orderdto "github.com/ak-repo/ecommerce-gin/internals/customer/order/order_dto"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/customer/order/order_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
)

type service struct {
	OrderRepo orderinterface.Repository
}

func NewOrderService(repo orderinterface.Repository) orderinterface.Service {
	return &service{OrderRepo: repo}
}

func (s *service) GetOrderByID(id uint) (*orderdto.OrderDetailResponse, error) {
	data, err := s.OrderRepo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("order not found")
	}

	order := orderdto.OrderDetailResponse{
		OrderID:     data.ID,
		UserID:      data.UserID,
		OrderDate:   data.CreatedAt,
		Status:      data.Status,
		TotalAmount: data.TotalAmount,
		Address: orderdto.OrderAddressDTO{
			ID:      data.ShippingAddress.ID,
			Street:  data.ShippingAddress.AddressLine,
			City:    data.ShippingAddress.City,
			State:   data.ShippingAddress.State,
			ZipCode: data.ShippingAddress.PostalCode,
			Country: data.ShippingAddress.Country,
		},
		Payment: orderdto.OrderPaymentDTO{
			PaymentID: data.Payment.ID,
			Method:    data.Payment.PaymentMethod,
			Amount:    data.Payment.Amount,
			Status:    data.Payment.Status,
		},
	}
	var items []orderdto.OrderItemDTO
	for _, i := range data.OrderItems {
		item := orderdto.OrderItemDTO{
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			UnitPrice: i.UnitPrice,
			Subtotal:  i.UnitPrice * float64(i.Quantity),
			ImageURL:  i.Product.ImageURL,
		}
		items = append(items, item)

	}
	order.Items = items

	return &order, nil
}

func (s *service) GetOrderByCustomerID(userID uint) ([]orderdto.CustomerOrder, error) {

	data, err := s.OrderRepo.GetOrderByCustomerID(userID)

	if data == nil || err != nil {
		return nil, errors.New("no orders found")
	}

	var orders []orderdto.CustomerOrder
	for _, i := range data {
		order := orderdto.CustomerOrder{
			OrderID:     i.ID,
			OrderDate:   i.OrderDate,
			Status:      i.Status,
			TotalAmount: i.TotalAmount,
			PaymentMode: i.Payment.PaymentMethod,
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func (s *service) CancelOrder(req *orderdto.CreateCancelRequest, userID uint) error {
	cancelOrder := models.OrderCancelRequest{
		OrderID: req.OrderID,
		UserID:  userID,
		Reason:  req.Reason,
	}

	return s.OrderRepo.CancelOrder(&cancelOrder)
}

func (s *service) CancellationResponse(orderID uint) (*orderdto.CancelRequestStatus, error) {
	data, err := s.OrderRepo.CancellationResponse(orderID)
	if err != nil || data == nil {
		return nil, errors.New("no order cancellation found")
	}
	response := orderdto.CancelRequestStatus{
		ID:      data.ID,
		Reason:  data.Reason,
		Status:  data.Status,
		OrderID: data.OrderID,
	}
	return &response, nil
}
