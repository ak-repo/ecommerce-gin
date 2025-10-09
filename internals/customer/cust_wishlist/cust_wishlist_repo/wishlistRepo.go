package custwishlistrepo

import (
	custwishlistinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_wishlist/cust_wishlist_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type WishlistRepo struct {
	DB *gorm.DB
}

func NewWishlistRepo(db *gorm.DB) custwishlistinterface.RepoInterface {

	return &WishlistRepo{DB: db}
}

// create wishlist
func (r *WishlistRepo) GetOrCreateWishlist(userID uint) (*models.Wishlist, error) {
	var wishlist models.Wishlist
	err := r.DB.
		Preload("WishlistItems").
		Preload("WishlistItems.Product").
		Where("user_id=?", userID).
		FirstOrCreate(&wishlist, &models.Wishlist{UserID: userID}).Error
	return &wishlist, err
}

// add to wishlist
func (r *WishlistRepo) AddToWishlistItem(list *models.WishlistItem) error {
	return r.DB.Create(list).Error
}

func (r *WishlistRepo) DeleteWishlistItem(listID uint) error {
	return r.DB.Delete(&models.WishlistItem{}, "id=?", listID).Error
}
