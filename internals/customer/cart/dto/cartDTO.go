package dto

// Add product to cart
type AddToCartDTO struct {
	ProductID uint `json:"product_id" form:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" form:"quantity" binding:"required,min=1"`
}

// Update cart item quantity
type UpdateCartItemDTO struct {
	CartItemID uint `json:"cart_item_id" form:"cart_item_id" binding:"required"`
	Quantity   int  `json:"quantity" form:"quantity" binding:"required,min=1"`
}

// Cart item response
type CartItemDTO struct {
	CartItemID      uint    `json:"cart_item_id"`
	CartID          uint    `json:"cart_id"`
	ProductID       uint    `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductImageURL string  `json:"product_image_url"`
	Price           float64 `json:"price"`
	Quantity        int     `json:"quantity"`
	Subtotal        float64 `json:"subtotal"`
}

// Full cart response
type CartDTO struct {
	CartID   uint          `json:"cart_id"`
	UserID   uint          `json:"user_id"`
	Items    []CartItemDTO `json:"items"`
	Total    float64       `json:"total"`
	Currency string        `json:"currency"`
}
