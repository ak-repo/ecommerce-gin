package checkouthandler

import (
	"net/http"

	checkoutinterface "github.com/ak-repo/ecommerce-gin/internals/customer/checkout/checkout_interface"
	checkoutdto "github.com/ak-repo/ecommerce-gin/internals/customer/checkout/dto"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	CheckoutService checkoutinterface.Service
}

func NewCheckoutHandler(service checkoutinterface.Service) checkoutinterface.Handler {
	return &Handler{
		CheckoutService: service,
	}
}

// GET cust/auth/checkout
func (h *Handler) CheckoutSummary(ctx *gin.Context) {

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}
	checkout, err := h.CheckoutService.CheckoutSummary(userID)
	if err != nil {

		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "checkout failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "checkout summary", checkout)

}

// POST cust/auth/checkout CheckoutHandler
func (h *Handler) ProcessCheckout(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}

	var req checkoutdto.CheckoutRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid inputes", err)
		return
	}
	req.UserID = userID
	response, err := h.CheckoutService.ProcessCheckout(&req)

	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "order creation failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "order created", response)

}
