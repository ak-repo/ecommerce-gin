package productdto

import reviewdto "github.com/ak-repo/ecommerce-gin/internals/review/review_dto"

// Customer-facing full product details
type CustomerProductResponse struct {
	ID            uint        `json:"id"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	SKU           string      `json:"sku"`
	BasePrice     float64     `json:"base_price"`
	DiscountPrice float64     `json:"discount_price,omitempty"` // optional if 0
	Stock         int         `json:"stock"`
	ImageURL      string      `json:"image_url,omitempty"`
	Category      CategoryDTO `json:"category"`

	IsPublished bool `json:"is_published"`
	Reviews     []reviewdto.ReviewResponse
}

// Lightweight product for listing (e.g., homepage, catalog)
type CustomerProductListItem struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	SKU           string  `json:"sku"`
	BasePrice     float64 `json:"base_price"`
	DiscountPrice float64 `json:"discount_price,omitempty"`
	ImageURL      string  `json:"image_url,omitempty"`
	CategoryName  string  `json:"category_name"`
}

// Category DTO reused
type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
