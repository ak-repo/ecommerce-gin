package wishlistdto

// WishlistItemDTO is used to return wishlist item details
type WishlistItemDTO struct {
	ID        uint `json:"id"`
	ProductID uint `json:"product_id"`
	Product   struct {
		Name  string `json:"name"`
		Price uint   `json:"price"`
	} `json:"product"`
}

// WishlistDTO is used to return a user's wishlist
type WishlistDTO struct {
	ID            uint              `json:"id"`
	UserID        uint              `json:"user_id"`
	WishlistItems []WishlistItemDTO `json:"wishlist_items"`
}
