package carthandler

import (
	"errors"
	"net/http"

	custcart "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart"
	cartinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart/cart_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService cartinterface.ServiceInterface
}

func NewCartHandler(cartService cartinterface.ServiceInterface) cartinterface.HandlerInterface {
	return &CartHandler{cartService: cartService}
}

// GET - cust/auth/cart => get user cart and it items
func (h *CartHandler) ShowUserCartHandler(ctx *gin.Context) {

	userIDAny, exists := ctx.Get("userID")
	userUID, ok := userIDAny.(uint)
	if !exists || !ok {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found or not valid", errors.New("id not found"))
		return

	}

	cart, err := h.cartService.CustomerCartService(userUID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "user cart db issue", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "customer cart fetch successfully", map[string]interface{}{
		"user_id":    cart.UserID,
		"data":       cart,
		"cart_count": len(cart.Items),
	})
}

// POST - cust/auth/cart  => add new cart item
func (h *CartHandler) AddtoCartHandler(ctx *gin.Context) {
	userIDAny, exists := ctx.Get("userID")
	if !exists {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid sesssion, try again", errors.New("id not found"))
		return
	}

	var inputCart custcart.AddToCartDTO
	if err := ctx.ShouldBind(&inputCart); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	userUID := userIDAny.(uint)
	if err := h.cartService.AddtoCartService(userUID, &inputCart); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "add to cart failed", err)
		return

	}
	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "add to cart successful", map[string]interface{}{
		"customer_id": userUID,
	})

}

// PATCH - cust/auth/cart/update/:id  => update quantity
func (h *CartHandler) UpdateCartQuantityHandler(ctx *gin.Context) {

	var inputCart custcart.UpdateCartItemDTO
	if err := ctx.ShouldBind(&inputCart); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	if err := h.cartService.UpdateQuantityService(&inputCart); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "update cart successful", map[string]interface{}{
		"cart_id":  inputCart.CartItemID,
		"quantity": inputCart.Quantity,
	})

}

// DELETE - cust/auth/cart/remove/:id => remove cart item
func (h *CartHandler) RemoveCartItemHandler(ctx *gin.Context) {

	// cartItemIDStr := ctx.PostForm("cart_item_id")
	// cartIemID, _ := strconv.ParseUint(cartItemIDStr, 10, 64)
	// log.Println(cartIemID)
	var id struct {
		CartItemID uint `json:"cart_item_id" form:"cart_item_id" binding:"required"`
	}
	if err := ctx.ShouldBind(&id); err != nil || id.CartItemID == 0 {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	if err := h.cartService.RemoveCartItemService(id.CartItemID); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "cart item remove failed", err)
		return
	}
	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "removed cart item", map[string]interface{}{
		"cart_id": id,
	})

}
