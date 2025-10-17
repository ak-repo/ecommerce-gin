package productdto

import "time"

// Create Product Request DTO
type CreateProductRequest struct {
	Title         string  `form:"title" json:"title" binding:"required"`
	Description   string  `form:"description" json:"description" binding:"required"`
	SKU           string  `form:"sku" json:"sku" binding:"required"`
	BasePrice     float64 `form:"base_price" json:"base_price" binding:"required,gt=0"`
	DiscountPrice float64 `form:"discount_price,omitempty" json:"discount_price,omitempty"`
	Stock         int     `form:"stock" json:"stock" binding:"required,gte=0"`
	ImageURL      string  `form:"image_url,omitempty" json:"image_url,omitempty"`
	CategoryID    uint    `form:"category_id" json:"category_id" binding:"required"`

	IsActive    bool `form:"-" json:"is_active"`
	IsPublished bool `form:"-" json:"is_published"`
}

// Update Product Request DTO (PATCH / PUT)
type UpdateProductRequest struct {
	ID            *uint    `json:"id" form:"-"`
	Title         *string  `json:"title,omitempty" form:"title"`
	Description   *string  `json:"description,omitempty" form:"description"`
	SKU           *string  `json:"sku,omitempty" form:"sku"`
	BasePrice     *float64 `json:"base_price,omitempty" form:"base_price"`
	DiscountPrice *float64 `json:"discount_price,omitempty" form:"discount_price"`
	Stock         *int     `json:"stock,omitempty" form:"stock"`
	ImageURL      *string  `json:"image_url,omitempty" form:"image_url"`
	CategoryID    *uint    `json:"category_id,omitempty" form:"category_id"`

	IsActive    *bool `json:"is_active,omitempty" form:"-"`
	IsPublished *bool `json:"is_published,omitempty" form:"-"`
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

	IsActive    bool      `json:"is_active"`
	IsPublished bool      `json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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

type ProductPagination struct {
	Page       int    `json:"page"`        // current page
	Limit      int    `json:"limit"`       // items per page
	Total      int64  `json:"total"`       // total items in DB
	TotalPages int    `json:"total_pages"` // total pages
	Query      string `json:"query"`
}
