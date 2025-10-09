package custwishlistinterface

import (
	custwishlist "github.com/ak-repo/ecommerce-gin/internals/customer/cust_wishlist"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	ListCustomerWishlistHandler(ctx *gin.Context)
	AddToWishlistHandler(ctx *gin.Context)
	RemoveFromWishlistHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	ListAllWishlistService(userID uint) (*custwishlist.WishlistDTO, error)
	AddToWishlistService(userID, productID uint) error
	RemoveWishlistService(wishlistItemID uint) error
}

type RepoInterface interface {
	GetOrCreateWishlist(userID uint) (*models.Wishlist, error)
	DeleteWishlistItem(listID uint) error
	AddToWishlistItem(list *models.WishlistItem) error
}
