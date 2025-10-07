package orderdto

import "time"

// Request for admin updates (e.g., change status)
type AdminUpdateOrderStatusRequest struct {
	OrderID uint   `json:"order_id" binding:"required"`
	Status  string `json:"status" binding:"required,oneof=pending confirmed shipped delivered cancelled refunded"`
}

// Response for admin when viewing orders
type AdminOrderResponse struct {
	OrderID     uint                 `json:"order_id"`
	UserID      uint                 `json:"user_id"`
	UserEmail   string               `json:"user_email,omitempty"` // optional if you fetch via join
	OrderDate   time.Time            `json:"order_date"`
	Status      string               `json:"status"`
	TotalAmount float64              `json:"total_amount"`
	Address     AdminOrderAddressDTO `json:"shipping_address"`
	Items       []AdminOrderItemResp `json:"items"`
	Payment     AdminOrderPaymentDTO `json:"payment"`
}

type AdminOrderItemResp struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unit_price"`
	Subtotal  float64 `json:"subtotal"`
	ImageURL  string  `json:"image_url,omitempty"`
}

type AdminOrderAddressDTO struct {
	ID      uint   `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

type AdminOrderPaymentDTO struct {
	PaymentID uint    `json:"payment_id"`
	Method    string  `json:"method"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
}
