package custcheckout

import (
	custcart "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart"
	custprofile "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile"
)

// POST

type CheckoutRequest struct {
	UserID      uint   `json:"user_id" form:"-"` // (set from auth middleware)
	AddressID   uint   `json:"address_id" binding:"required"`
	PaymentMode string `json:"payment_mode" binding:"required,oneof=COD ONLINE"`
}

// GET
type CheckoutSummaryResponse struct {
	Items        []custcart.CartItemDTO `json:"items"`
	TotalItems   int                    `json:"total_items"`
	SubTotal     float64                `json:"subtotal"`
	ShippingFee  float64                `json:"shipping_fee"`
	GrandTotal   float64                `json:"grand_total"`
	Address      custprofile.AddressDTO `json:"address"`
	PaymentModes []string               `json:"payment_modes"`
}


//
// CheckoutResponse is returned after confirming checkout (POST).
// If payment = COD → order created.
// If payment = ONLINE → payment link/session returned.
type CheckoutResponse struct {
	OrderID     uint    `json:"order_id"`
	PaymentMode string  `json:"payment_mode"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"` // Pending, Paid, etc.


}

