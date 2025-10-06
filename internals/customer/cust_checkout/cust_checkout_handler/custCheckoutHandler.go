package custcheckouthandler

import (
	"errors"
	"net/http"

	custcheckout "github.com/ak-repo/ecommerce-gin/internals/customer/cust_checkout"
	custcheckoutinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_checkout/cust_checkout_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CustomerCheckoutHandler struct {
	CustomerCheckoutService custcheckoutinterface.ServiceInterface
}

func NewCustomerCheckoutHandler(service custcheckoutinterface.ServiceInterface) custcheckoutinterface.HandlerInterface {
	return &CustomerCheckoutHandler{
		CustomerCheckoutService: service,
	}
}

// GET cust/auth/checkout  => ShowCheckoutPageHandler
func (h *CustomerCheckoutHandler) CustomerShowCheckoutHandler(ctx *gin.Context) {

	id, exists := ctx.Get("userID")
	userID := id.(uint)
	if !exists || userID == 0 {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "user not found", errors.New("user id not valid"))
		return
	}

	checkout, err := h.CustomerCheckoutService.CheckoutSummaryService(userID)
	if err != nil {

		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "checkout failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "checkout summary", map[string]interface{}{
		"data": checkout,
	})

}

// POST cust/auth/checkout CheckoutHandler

func (h *CustomerCheckoutHandler) CustomerCheckoutHandler(ctx *gin.Context) {
	id, exists := ctx.Get("userID")
	userID := id.(uint)
	if !exists || userID == 0 {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "user not found", errors.New("user id not valid"))
		return
	}

	var req custcheckout.CheckoutRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid inputes", err)
		return
	}
	req.UserID = userID
	response, err := h.CustomerCheckoutService.ProcessCheckoutService(&req)

	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "order creation failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "order created", map[string]interface{}{
		"data": response,
	})

}
