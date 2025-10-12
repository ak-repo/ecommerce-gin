package wishlistrepository

import (
	wishlistinterface "github.com/ak-repo/ecommerce-gin/internals/customer/wishlist/wishlist_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewWishlistRepo(db *gorm.DB) wishlistinterface.Repository {

	return &repository{DB: db}
}

// create wishlist
func (r *repository) List(userID uint) (*models.Wishlist, error) {
	var wishlist models.Wishlist
	err := r.DB.
		Preload("WishlistItems").
		Preload("WishlistItems.Product").
		Where("user_id=?", userID).
		FirstOrCreate(&wishlist, &models.Wishlist{UserID: userID}).Error
	return &wishlist, err
}

// add to wishlist
func (r *repository) Add(list *models.WishlistItem) error {
	return r.DB.Create(list).Error
}

func (r *repository) Remove(listID uint) error {
	return r.DB.Delete(&models.WishlistItem{}, "id=?", listID).Error
}
