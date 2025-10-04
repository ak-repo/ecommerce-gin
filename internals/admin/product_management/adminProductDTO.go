package productmanagement

import "time"

// Create Product Request DTO
type CreateProductRequest struct {
	Title         string   `json:"title" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	SKU           string   `json:"sku" binding:"required"`
	BasePrice     float64  `json:"base_price" binding:"required,gt=0"`
	DiscountPrice float64  `json:"discount_price,omitempty"` // optional
	Stock         int      `json:"stock" binding:"required,gte=0"`
	ImageURL      string   `json:"image_url,omitempty"`
	CategoryID    uint     `json:"category_id" binding:"required"`
	
	IsActive      bool     `json:"is_active"`
	IsPublished   bool     `json:"is_published"`
}

// Update Product Request DTO (PATCH / PUT)
type UpdateProductRequest struct {
	Title         *string  `json:"title,omitempty"`
	Description   *string  `json:"description,omitempty"`
	SKU           *string  `json:"sku,omitempty"`
	BasePrice     *float64 `json:"base_price,omitempty"`
	DiscountPrice *float64 `json:"discount_price,omitempty"`
	Stock         *int     `json:"stock,omitempty"`
	ImageURL      *string  `json:"image_url,omitempty"`
	CategoryID    *uint    `json:"category_id,omitempty"`

	IsActive      *bool    `json:"is_active,omitempty"`
	IsPublished   *bool    `json:"is_published,omitempty"`
}

// Full Product Response DTO
type ProductResponse struct {
	ID            uint        `json:"id"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	SKU           string      `json:"sku"`
	BasePrice     float64     `json:"base_price"`
	DiscountPrice float64     `json:"discount_price"`
	Stock         int         `json:"stock"`
	ImageURL      string      `json:"image_url"`
	Category      CategoryDTO `json:"category"`

	IsActive      bool        `json:"is_active"`
	IsPublished   bool        `json:"is_published"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

// Lightweight product for admin/product list
type ProductListItem struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	SKU           string  `json:"sku"`
	BasePrice     float64 `json:"base_price"`
	DiscountPrice float64 `json:"discount_price"`
	Stock         int     `json:"stock"`
	IsActive      bool    `json:"is_active"`
	IsPublished   bool    `json:"is_published"`
	ImageURL      string  `json:"image_url"`
}

// Category DTO for embedding
type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
