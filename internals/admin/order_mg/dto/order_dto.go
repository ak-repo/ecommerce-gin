package orderdto

import "time"

type AllOrderResponse struct {
	OrderID     uint            `json:"order_id"`
	UserID      uint            `json:"user_id"`
	UserEmail   string          `json:"user_email,omitempty"`
	OrderDate   time.Time       `json:"order_date"`
	Status      string          `json:"status"`
	TotalAmount float64         `json:"total_amount"`
	Address     OrderAddressDTO `json:"shipping_address"`
	Items       []OrderItemDTO  `json:"items"`
	Payment     OrderPaymentDTO `json:"payment"`
}

type CustomerOrder struct {
	OrderID     uint      `json:"order_id"`
	OrderDate   time.Time `json:"order_date"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
	PaymentMode string    `json:"payment_mode"`
}

// Detailed single order view
type OrderDetailResponse struct {
	UserID      uint            `json:"user_id"`
	OrderID     uint            `json:"order_id"`
	OrderDate   time.Time       `json:"order_date"`
	Status      string          `json:"status"`
	TotalAmount float64         `json:"total_amount"`
	Address     OrderAddressDTO `json:"shipping_address"`
	Items       []OrderItemDTO  `json:"items"`
	Payment     OrderPaymentDTO `json:"payment"`
}

// Shared Order Item
type OrderItemDTO struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name,omitempty"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price,omitempty"` // optional for customer
	Subtotal  float64 `json:"subtotal"`
	ImageURL  string  `json:"image_url,omitempty"`
}

// Shared Address
type OrderAddressDTO struct {
	ID      uint   `json:"id,omitempty"` // present only for admin or when required
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

// Shared Payment info
type OrderPaymentDTO struct {
	PaymentID uint    `json:"payment_id,omitempty"`
	Method    string  `json:"method"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
}

// Request: Admin updates order status
type AdminUpdateOrderStatusRequest struct {
	OrderID uint   `json:"order_id" binding:"required"`
	Status  string `json:"status" binding:"required,oneof=pending confirmed shipped delivered cancelled refunded"`
}

// Admin cancel request review
type AdminCancelRequestResponse struct {
	ID          uint      `json:"id"`
	OrderID     uint      `json:"order_id"`
	UserID      uint      `json:"user_id"`
	Customer    string    `json:"customer"`
	OrderStatus string    `json:"order_status"`
	Reason      string    `json:"reason"`
	Status      string    `json:"status"` // PENDING / APPROVED / REJECTED
	CreatedAt   time.Time `json:"created_at"`
}

type AdminCancelRequestListResponse struct {
	CancelRequests []AdminCancelRequestResponse `json:"cancel_requests"`
}
