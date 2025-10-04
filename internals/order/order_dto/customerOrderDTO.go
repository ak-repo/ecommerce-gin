package orderdto

import "time"

type CustomerCreateOrderRequest struct {
	AddressID uint                      `json:"address_id" binding:"required"`
	Items     []CustomerOrderItemCreate `json:"items" binding:"required"`
}

type CustomerOrderItemCreate struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type CustomerOrderResponse struct {
	OrderID     uint                      `json:"order_id"`
	OrderDate   time.Time                 `json:"order_date"`
	Status      string                    `json:"status"`
	TotalAmount float64                   `json:"total_amount"`
	Items       []CustomerOrderItemDetail `json:"items"`
	PaymentID   uint                      `json:"payment_id,omitempty"`
}

type CustomerOrderItemDetail struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Subtotal  float64 `json:"subtotal"`
}
