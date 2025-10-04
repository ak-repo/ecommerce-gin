package cartinterface

import (
	custcart "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	ShowUserCartHandler(ctx *gin.Context)
	AddtoCartHandler(ctx *gin.Context)
	UpdateCartQuantityHandler(ctx *gin.Context)
	RemoveCartItemHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	AddtoCartService(userID uint, addtoCart *custcart.AddToCartDTO) error
	CustomerCartService(userID uint) (*custcart.CartDTO, error)
	UpdateQuantityService(updatedCart *custcart.UpdateCartItemDTO) error
	RemoveCartItemService(cartItemID uint) error
}

type RepoInterface interface {
	GetorCreateCart(userID uint) (*models.Cart, error)
	GetCartItem(cartID, ProductID uint) (*models.CartItem, error)
	GetCartItemByID(cartItemID uint) (*models.CartItem, error)
	GetAllCartItems(cartID uint) ([]models.CartItem, error)
	UpdateCartItem(cartItem *models.CartItem) error
	CreateCartItem(cartItem *models.CartItem) error
	DeleteCartItem(cartItemID uint) error
}
