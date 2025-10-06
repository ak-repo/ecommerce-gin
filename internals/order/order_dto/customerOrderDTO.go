package orderdto

type CheckoutRequest struct {
	UserID      uint   `json:"user_id"` // from auth middleware
	AddressID   uint   `json:"address_id" binding:"required"`
	PaymentMode string `json:"payment_mode" binding:"required,oneof=COD ONLINE"`
	CouponCode  string `json:"coupon_code,omitempty"` // optional
}

type CheckoutItem struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	Subtotal  float64 `json:"subtotal"`
}

type CheckoutSummary struct {
	Subtotal     float64         `json:"subtotal"`
	ShippingFee  float64         `json:"shipping_fee"`
	Discount     float64         `json:"discount"`
	GrandTotal   float64         `json:"grand_total"`
	PaymentModes []string        `json:"payment_modes"`
	Items        []CheckoutItem  `json:"items"`
	Address      AddressResponse `json:"address"`
}

type AddressResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Pincode   string `json:"pincode"`
	IsDefault bool   `json:"is_default"`
}
