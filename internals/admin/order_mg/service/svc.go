package orderservice

import (
	"errors"

	orderdto "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/dto"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/order_interface"
)

type service struct {
	OrderRepo orderinterface.Repository
}

func NewOrderServiceMG(repo orderinterface.Repository) orderinterface.Service {
	return &service{OrderRepo: repo}
}

func (s *service) GetAllOrders() ([]orderdto.AllOrderResponse, error) {
	data, err := s.OrderRepo.GetAllOrders()
	if data == nil || err != nil {
		return nil, errors.New("orders not found")
	}

	var orders []orderdto.AllOrderResponse
	for _, val := range data {
		// each products
		var items []orderdto.OrderItemDTO
		for _, i := range val.OrderItems {
			item := orderdto.OrderItemDTO{
				ProductID: i.ProductID,
				Quantity:  i.Quantity,
				UnitPrice: i.UnitPrice,
				Subtotal:  i.UnitPrice * float64(i.Quantity),
				ImageURL:  i.Product.ImageURL,
			}
			items = append(items, item)

		}
		order := orderdto.AllOrderResponse{
			OrderID:     val.ID,
			UserID:      val.UserID,
			OrderDate:   val.CreatedAt,
			Status:      val.Status,
			TotalAmount: val.TotalAmount,
			UserEmail:   val.User.Email,
			Address: orderdto.OrderAddressDTO{
				ID:      val.ShippingAddress.ID,
				Street:  val.ShippingAddress.AddressLine,
				City:    val.ShippingAddress.City,
				State:   val.ShippingAddress.State,
				ZipCode: val.ShippingAddress.PostalCode,
				Country: val.ShippingAddress.Country,
			},
			Items: items,
			Payment: orderdto.OrderPaymentDTO{
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
		UserEmail:   data.User.Email,
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

func (s *service) UpdateStatus(req *orderdto.AdminUpdateOrderStatusRequest) error {

	validStatuses := map[string]bool{
		"pending": true, "accepted": true, "confirmed": true,
		"shipped": true, "delivered": true, "cancelled": true,
		"refunded": true, "completed": true,
	}
	if !validStatuses[req.Status] {
		return errors.New("invalid status")
	}
	return s.OrderRepo.UpdateStatus(req)

}

func (s *service) GetAllCancels() ([]orderdto.AdminCancelRequestResponse, error) {

	data, err := s.OrderRepo.GetAllCancels()
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

func (s *service) AcceptCancel(reqID uint) error {
	orderID, err := s.OrderRepo.AcceptCancel(reqID)
	if err != nil {
		return err
	}

	order := orderdto.AdminUpdateOrderStatusRequest{
		OrderID: orderID,
		Status:  "cancelled",
	}
	if err := s.OrderRepo.UpdateStatus(&order); err != nil {
		return err
	}

	return nil

}

func (s *service) RejectCancel(reqID uint) error {
	orderID, err := s.OrderRepo.RejectCancel(reqID)
	if err != nil {
		return err
	}

	order := orderdto.AdminUpdateOrderStatusRequest{
		OrderID: orderID,
		Status:  "confirmed",
	}
	if err := s.OrderRepo.UpdateStatus(&order); err != nil {
		return err
	}
	return nil

}
