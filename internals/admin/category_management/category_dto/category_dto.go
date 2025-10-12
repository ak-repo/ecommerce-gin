package categorydto

type CategoryDTO struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ProductCount int    `json:"product_count"`
}

type CategoryDeatiledResponse struct {
	ID           uint              `json:"id"`
	Name         string            `json:"name"`
	ProductCount int               `json:"product_count"`
	Products     []ProductListItem `json:"products"`
}
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
