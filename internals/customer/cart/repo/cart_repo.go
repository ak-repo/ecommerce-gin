package cartrepo

import (
	"errors"

	cartinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cart/cart_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewCartRepo(db *gorm.DB) cartinterface.Repository {
	return &repository{DB: db}
}

// full cart
func (r *repository) GetOrCreateCart(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.DB.Preload("CartItems").
		Preload("CartItems.Product").
		FirstOrCreate(&cart, models.Cart{UserID: userID}).Error
	return &cart, err
}

func (r *repository) DeleteCart(cartID uint) error {
	return r.DB.Delete(&models.Cart{}, cartID).Error
}

// get one cart item by cartID and productID matching
func (r *repository) GetCartItem(cartID, productID uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &item, err
}

func (r *repository) GetCartItemByID(cartItemID uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.DB.First(&item, cartItemID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &item, err
}

func (r *repository) GetAllCartItems(cartID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	err := r.DB.Preload("Product").Where("cart_id=?", cartID).Find(&items).Error
	return items, err
}

func (r *repository) CreateCartItem(item *models.CartItem) error {
	return r.DB.Create(item).Error
}

func (r *repository) UpdateCartItem(item *models.CartItem) error {
	return r.DB.Save(item).Error
}

func (r *repository) DeleteCartItem(cartItemID uint) error {
	return r.DB.Delete(&models.CartItem{}, cartItemID).Error
}

func (r *repository) DeleteCartItemsByCartID(cartID uint) error {
	return r.DB.Delete(&models.CartItem{}, "cart_id=?", cartID).Error
}
