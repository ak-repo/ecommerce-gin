package orderdto

import "time"

// 1 Response for showing all orders of a customer
type CustomerOrderListResponse struct {
	Orders []CustomerOrderSummary `json:"orders"`
}
type CustomerOrderSummary struct {
	OrderID     uint      `json:"order_id"`
	OrderDate   time.Time `json:"order_date"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
	PaymentMode string    `json:"payment_mode"`
}





// 2. Show individual order (detailed view)
//
type CustomerOrderDetailResponse struct {
	OrderID     uint                    `json:"order_id"`
	OrderDate   time.Time               `json:"order_date"`
	Status      string                  `json:"status"`
	TotalAmount float64                 `json:"total_amount"`
	Address     CustomerOrderAddressDTO `json:"shipping_address"`
	Items       []CustomerOrderItemResp `json:"items"`
	Payment     CustomerOrderPaymentDTO `json:"payment"`
}

// Each product inside this specific order
type CustomerOrderItemResp struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Subtotal  float64 `json:"subtotal"`
	ImageURL  string  `json:"image_url,omitempty"`
}

// Shipping address used for this order
type CustomerOrderAddressDTO struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

// Payment info for this order
type CustomerOrderPaymentDTO struct {
	Method string  `json:"method"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
