package carthandler

import (
	"net/http"
	"strconv"

	"github.com/ak-repo/ecommerce-gin/internal/dto"
	cartservice "github.com/ak-repo/ecommerce-gin/internal/services/cartService"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService cartservice.CartService
}

func NewCartHandler(cartService cartservice.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

// show user cart
func (h *CartHandler) ShowUserCartHandler(ctx *gin.Context) {

	email, _ := ctx.Get("email")
	emailStr := email.(string)

	cart, err := h.cartService.UserCartService(emailStr)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/notification/error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	msg, _ := ctx.Cookie("flash")
	ctx.SetCookie("flash", "", -1, "/", "localhost", false, true)
	ctx.HTML(http.StatusOK, "pages/user/cart.html", gin.H{
		"Cart":      cart,
		"CartCount": len(cart.Items),
		"User":      email,
		"Message":   msg,
	})
}

// add to cart
func (h *CartHandler) AddtoCartHandler(ctx *gin.Context) {

	email, exists := ctx.Get("email")
	if !exists {
		ctx.HTML(http.StatusBadRequest, "pages/notification/error.html", gin.H{
			"Error": "invalid sesssion, try again",
		})
		return
	}

	var inputCart dto.AddToCartDTO
	if err := ctx.ShouldBind(&inputCart); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/notification/error.html", gin.H{
			"User":  email,
			"Error": "invalid input",
		})
		return
	}

	emailStr := email.(string)
	if err := h.cartService.AddtoCartService(emailStr, &inputCart); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/notification/error.html", gin.H{
			"User":  email,
			"Error": "add to cart failed",
		})
		return

	}

	ctx.SetCookie("flash", "Add to cart succesfully", 3600, "/", "localhost", false, true)

	ctx.Redirect(http.StatusSeeOther, "/product/"+strconv.FormatUint(uint64(inputCart.ProductID), 10))

}

func (h *CartHandler) UpdateCartQuantityHandler(ctx *gin.Context) {

	email, exists := ctx.Get("email")
	if !exists {
		ctx.HTML(http.StatusBadRequest, "ppages/notification/error.html", gin.H{
			"Error": "invalid sesssion, try again",
		})
		return
	}

	var inputCart dto.UpdateCartItemDTO
	if err := ctx.ShouldBind(&inputCart); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/notification/error.html", gin.H{
			"User":  email,
			"Error": "invalid input",
		})
		return
	}

	if err := h.cartService.UpdateQuantityService(&inputCart); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/notification/error.html", gin.H{
			"User":  email,
			"Error": err.Error(),
		})
		return
	}

	ctx.SetCookie("flash", "cart updated", 10, "/", "localhost", false, true)
	ctx.Redirect(http.StatusSeeOther, "/user/cart")

}

// remove cart item
func (h *CartHandler) RemoveCartItemHandler(ctx *gin.Context) {

	email, exists := ctx.Get("email")
	if !exists {
		ctx.HTML(http.StatusBadRequest, "ppages/notification/error.html", gin.H{
			"Error": "invalid sesssion, try again",
		})
		return
	}

	cartItemIDStr := ctx.PostForm("cart_item_id")
	cartIemID, _ := strconv.ParseUint(cartItemIDStr, 10, 64)

	if err := h.cartService.RemoveCartItemService(uint(cartIemID)); err != nil {
		ctx.HTML(http.StatusInternalServerError, "ppages/notification/error.html", gin.H{
			"Error": "can't remove item, try again",
			"User":  email,
		})
		return
	}
	ctx.SetCookie("flash", "Removed item from cart", 10, "/", "localhost", false, true)

	ctx.Redirect(http.StatusSeeOther, "/user/cart")

}
