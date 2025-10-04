package orderservices

import (
	"errors"

	orderdto "github.com/ak-repo/ecommerce-gin/internals/order/order_dto"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/order/order_interface"
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

// Get OrderBy ID
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
		Address: orderdto.AdminOrderAddressDTO{
			ID:      data.ShippingAddress.ID,
			Street:  data.ShippingAddress.AddressLine,
			City:    data.ShippingAddress.City,
			State:   data.ShippingAddress.State,
			ZipCode: data.ShippingAddress.PostalCode,
			Country: data.ShippingAddress.Country,
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
