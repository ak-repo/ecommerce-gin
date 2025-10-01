package dto

// Add product to cart request (works for JSON + HTML forms)
type AddToCartDTO struct {
	ProductID uint `json:"product_id" form:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" form:"quantity" binding:"required,min=1"`
}

// Update quantity in cart
type UpdateCartItemDTO struct {
	CartItemID uint `json:"cart_item_id" form:"cart_item_id" binding:"required"`
	Quantity   int  `json:"quantity" form:"quantity" binding:"required,min=1"`
	
}

// Response for each item in the cart
type CartItemDTO struct {
	CartItemID uint    `json:"cart_item_id"`
	CartID     uint    `json:"cart_id"`
	ProductID  uint    `json:"product_id"`
	Product    string  `json:"product"` // Product Name
	ProductImageURL string `json:"product_image_url"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	Subtotal   float64 `json:"subtotal"`
}

// Response for the whole cart
type CartDTO struct {
	CartID   uint          `json:"cart_id"`
	UserID   uint          `json:"user_id"`
	Items    []CartItemDTO `json:"items"`
	Total    float64       `json:"total"`
	Currency string        `json:"currency"`
}
