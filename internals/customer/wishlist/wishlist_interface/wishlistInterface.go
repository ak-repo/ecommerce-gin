package wishlistinterface

import (
	wishlistdto "github.com/ak-repo/ecommerce-gin/internals/customer/wishlist/wishlist_dto"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	List(ctx *gin.Context)
	Add(ctx *gin.Context)
	Remove(ctx *gin.Context)
}

type Service interface {
	List(userID uint) (*wishlistdto.WishlistDTO, error)
	Add(userID, productID uint) error
	Remove(wishlistItemID uint) error
}

type Repository interface {
	List(userID uint) (*models.Wishlist, error)
	Add(item *models.WishlistItem) error
	Remove(wishlistItemID uint) error
}
