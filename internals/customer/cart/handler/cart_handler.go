package carthandler

import (
	"net/http"

	cartinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cart/cart_interface"
	"github.com/ak-repo/ecommerce-gin/internals/customer/cart/dto"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service cartinterface.Service
}

func NewCartHandler(service cartinterface.Service) cartinterface.Handler {
	return &handler{service: service}
}

// GET - cust/auth/cart
func (h *handler) GetUserCart(ctx *gin.Context) {

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "user not found", err)
		return
	}

	cart, err := h.service.GetUserCart(userID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "failed to fetch cart", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "cart fetched successfully", map[string]interface{}{
		"data":       cart,
		"cart_count": len(cart.Items),
	})
}

// POST - cust/auth/cart
func (h *handler) AddItem(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "user not found", err)
		return
	}

	var req dto.AddToCartDTO
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	if err := h.service.AddItem(userID, &req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "add to cart failed", err)
		return

	}
	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "aItem added successfully", nil)

}

// PATCH - cust/auth/cart/:id
func (h *handler) UpdateQuantity(ctx *gin.Context) {

	var req dto.UpdateCartItemDTO
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	if err := h.service.UpdateQuantity(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "quantity updated", nil)

}

// DELETE - cust/auth/cart/:id
func (h *handler) RemoveItem(ctx *gin.Context) {

	var id struct {
		CartItemID uint `json:"cart_item_id" form:"cart_item_id" binding:"required"`
	}
	if err := ctx.ShouldBind(&id); err != nil || id.CartItemID == 0 {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	if err := h.service.RemoveItem(id.CartItemID); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "failed to remove cart item.", err)
		return
	}
	utils.RenderSuccess(ctx, http.StatusOK, "customer", "item removed", nil)

}
