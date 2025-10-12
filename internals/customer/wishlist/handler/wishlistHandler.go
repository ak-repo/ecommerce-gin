package wishlisthandler

import (
	"net/http"
	"strconv"

	wishlistinterface "github.com/ak-repo/ecommerce-gin/internals/customer/wishlist/wishlist_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	WishlistService wishlistinterface.Service
}

func NewWishlistHandler(service wishlistinterface.Service) wishlistinterface.Handler {
	return &handler{WishlistService: service}
}

// GET - cust/auth/wishlist
func (h *handler) List(ctx *gin.Context) {

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}

	wishlist, err := h.WishlistService.List(userID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "wishlist not found", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "wishlist", map[string]interface{}{
		"data": wishlist,
	})
}

// POST - cust/auth/wishlist/:id
func (h *handler) Add(ctx *gin.Context) {
	id := ctx.Param("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "product id not valid", err)
		return
	}
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}

	if err := h.WishlistService.Add(userID, uint(productID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "add to wishlist failed", err)
		return

	}

	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "item added into wishlist", nil)
}

// DELETE - cust/auth/wishlist/:id
func (h *handler) Remove(ctx *gin.Context) {
	id := ctx.Param("id")
	wItemID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "product id not valid", err)
		return
	}

	if err := h.WishlistService.Remove(uint(wItemID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "remove from wishlist failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "item removed", nil)
}
