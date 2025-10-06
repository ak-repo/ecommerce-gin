package cartrepo

import (
	"errors"

	cartinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart/cart_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) cartinterface.RepoInterface {
	return &cartRepository{DB: db}
}

// create or get cart
func (r *cartRepository) GetorCreateCart(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.DB.Preload("CartItems").FirstOrCreate(&cart, models.Cart{UserID: userID}).Error
	return &cart, err
}

// delete cart
func (r *cartRepository) DeleteCart(cartID uint) error {
	return r.DB.Delete(&models.Cart{}, "id=?", cartID).Error

}

// get one cart item
func (r *cartRepository) GetCartItem(cartID, ProductID uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.DB.Where("cart_id=? AND product_id=?", cartID, ProductID).First(&item).Error
	return &item, err

}

// get cart item by id
func (r *cartRepository) GetCartItemByID(cartItemID uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.DB.First(&item, cartItemID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &item, err
}

// get all cart items
func (r *cartRepository) GetAllCartItems(cartID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	err := r.DB.Preload("Product").Where("cart_id=?", cartID).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

// update cart item
func (r *cartRepository) UpdateCartItem(cartItem *models.CartItem) error {

	return r.DB.Save(&cartItem).Error

}

// Create cart item
func (r *cartRepository) CreateCartItem(cartItem *models.CartItem) error {
	return r.DB.Create(&cartItem).Error
}

// Delete cart item by cartItem id
func (r *cartRepository) DeleteCartItem(cartItemID uint) error {

	return r.DB.Delete(&models.CartItem{}, cartItemID).Error
}

// Delete cart item by cartItem id
func (r *cartRepository) DeleteCartItemBycartID(cartID uint) error {

	return r.DB.Delete(&models.CartItem{}, "cart_id=?", cartID).Error
}
