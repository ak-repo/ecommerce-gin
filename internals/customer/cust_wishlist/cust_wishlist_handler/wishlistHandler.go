package custwishlisthandler

import (
	"errors"
	"net/http"
	"strconv"

	custwishlistinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_wishlist/cust_wishlist_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	WishlistService custwishlistinterface.ServiceInterface
}

func NewWishlistHandler(service custwishlistinterface.ServiceInterface) custwishlistinterface.HandlerInterface {
	return &WishlistHandler{WishlistService: service}
}

// GET - cust/auth/wishlist
func (h *WishlistHandler) ListCustomerWishlistHandler(ctx *gin.Context) {

	id, exists := ctx.Get("userID")
	if !exists || id.(uint) == 0 {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "unauthorised", errors.New("user id is missing or session ended"))
		return
	}
	wishlist, err := h.WishlistService.ListAllWishlistService(id.(uint))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "wishlist not found", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "wishlist", map[string]interface{}{
		"data": wishlist,
	})
}

// POST - cust/auth/wishlist/:id
func (h *WishlistHandler) AddToWishlistHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "product id not valid", err)
		return
	}
	userID, exists := ctx.Get("userID")
	if !exists || userID.(uint) == 0 {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "unauthorised", errors.New("user id is missing or session ended"))
		return
	}

	if err := h.WishlistService.AddToWishlistService(userID.(uint), uint(productID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "add to wishlist failed", err)
		return

	}

	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "item added into wishlist", nil)
}

// DELETE - cust/auth/wishlist/:id
func (h *WishlistHandler) RemoveFromWishlistHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	wItemID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "product id not valid", err)
		return
	}

	if err := h.WishlistService.RemoveWishlistService(uint(wItemID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "remove from wishlist failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "item removed", nil)
}
