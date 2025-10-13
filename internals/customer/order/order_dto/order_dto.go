package orderdto

import "time"


type OrderItemDTO struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name,omitempty"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price,omitempty"` // optional for customer
	Subtotal  float64 `json:"subtotal"`
	ImageURL  string  `json:"image_url,omitempty"`
}

type OrderAddressDTO struct {
	ID      uint   `json:"id,omitempty"` // present only for admin or when required
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

type OrderPaymentDTO struct {
	PaymentID uint    `json:"payment_id,omitempty"`
	Method    string  `json:"method"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
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
	UserID      uint                 `json:"user_id"`
	OrderID     uint                 `json:"order_id"`
	OrderDate   time.Time            `json:"order_date"`
	Status      string               `json:"status"`
	TotalAmount float64              `json:"total_amount"`
	Address     OrderAddressDTO      `json:"shipping_address"`
	Items       []OrderItemDTO       `json:"items"`
	Payment     OrderPaymentDTO      `json:"payment"`
	CancelReq   *CancelRequestStatus `json:"cancel_request,omitempty"`
}

type CreateCancelRequest struct {
	OrderID uint   `json:"order_id" binding:"required"`
	Reason  string `json:"reason" binding:"required,min=5,max=255"`
}

type CancelRequestStatus struct {
	ID      uint   `json:"id"`
	OrderID uint   `json:"order_id"`
	Status  string `json:"status"`
	Reason  string `json:"reason"`
}
