package productdto

import "time"

// Customer-facing full product details
type CustomerProductResponse struct {
	ID              uint              `json:"id"`
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	SKU             string            `json:"sku"`
	BasePrice       float64           `json:"base_price"`
	DiscountPrice   float64           `json:"discount_price,omitempty"` // optional if 0
	Stock           int               `json:"stock"`
	ImageURL        string            `json:"image_url,omitempty"`
	Category        CategoryDTO       `json:"category"`
	Reviews         []Reviews         `json:"reviews"`
	SimilarProducts []SimilarProducts `json:"similar_products"`
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

type Reviews struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Rating    uint      `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type FilterParams struct {
	Category string  `form:"category"`
	MinPrice float64 `form:"min_price"`
	MaxPrice float64 `form:"max_price"`
	Search   string  `form:"search"`
	Sort     string  `form:"sort"`
	Page     int     `form:"page"`
	Limit    int     `form:"limit"`
}

// {
//   "category": "electronics",
//   "min_price": 1000,
//   "max_price": 3000,
//   "search": "headphones",
//   "sort": "price_desc",
//   "page": 1,
//   "limit": 10
// }

type SimilarProducts struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	BasePrice     float64 `json:"base_price"`
	DiscountPrice float64 `json:"discount_price,omitempty"`
	ImageURL      string  `json:"image_url,omitempty"`
	CategoryName  string  `json:"category_name"`
}
