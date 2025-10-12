package cartinterface

import (
	"github.com/ak-repo/ecommerce-gin/internals/customer/cart/dto"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetUserCart(ctx *gin.Context)
	AddItem(ctx *gin.Context)
	UpdateQuantity(ctx *gin.Context)
	RemoveItem(ctx *gin.Context)
}

type Service interface {
	AddItem(userID uint, input *dto.AddToCartDTO) error
	GetUserCart(userID uint) (*dto.CartDTO, error)
	UpdateQuantity(input *dto.UpdateCartItemDTO) error
	RemoveItem(cartItemID uint) error
	DeleteCart(cartID uint) error
}

type Repository interface {
	GetOrCreateCart(userID uint) (*models.Cart, error)
	GetCartItem(cartID, productID uint) (*models.CartItem, error)
	GetCartItemByID(cartItemID uint) (*models.CartItem, error)
	GetAllCartItems(cartID uint) ([]models.CartItem, error)
	UpdateCartItem(cartItem *models.CartItem) error
	CreateCartItem(cartItem *models.CartItem) error
	DeleteCartItem(cartItemID uint) error
	DeleteCart(cartID uint) error
	DeleteCartItemsByCartID(cartID uint) error
}
