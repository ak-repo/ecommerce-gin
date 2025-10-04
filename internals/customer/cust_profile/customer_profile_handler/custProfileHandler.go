package customerprofilehandler

import (
	"errors"
	"net/http"

	custprofile "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile"
	customerprofileinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile/customer_profile_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CustomerProfileHandler struct {
	CustomerProfileService customerprofileinterface.ServiceInterface
}

func NewCustomerProfileHandler(customerProfileService customerprofileinterface.ServiceInterface) customerprofileinterface.HandlerInterface {
	return &CustomerProfileHandler{CustomerProfileService: customerProfileService}
}

// User profile GET cust/profile
func (h *CustomerProfileHandler) CustomerProfileHandler(ctx *gin.Context) {

	userIDAny, exists := ctx.Get("userID")
	userUID, ok := userIDAny.(uint)

	if !exists || !ok || userUID == 0 {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "no user id found", errors.New("requet have no user id token"))
		return
	}

	profile, err := h.CustomerProfileService.CustomerProfileService(userUID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "DB issue", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "user profile found", map[string]interface{}{
		"user_id": userUID,
		"Profile": profile,
	})

}

// GET cust/profile/address -> shop user address form
func (h *CustomerProfileHandler) GetCustomerAddress(ctx *gin.Context) {
	userIDAny, _ := ctx.Get("userID")
	userUID := userIDAny.(uint)

	address, err := h.CustomerProfileService.CustomerAddressService(userUID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "DB issue", err)
		return
	}
	utils.RenderSuccess(ctx, http.StatusOK, "customer", "user address found", map[string]interface{}{
		"user_id": userUID,
		"address": address,
	})

}

// POST cust/address/update -> add or update user address
func (h *CustomerProfileHandler) CustomerAddressUpdateHandler(ctx *gin.Context) {
	userIDAny, _ := ctx.Get("userID")
	userUID := userIDAny.(uint)

	var address custprofile.AddressDTO
	if err := ctx.ShouldBind(&address); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "binding isssue", err)
		return
	}

	if err := h.CustomerProfileService.CustomerAddressUpdateService(&address, userUID); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "address update failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "address updated", map[string]interface{}{
		"user_id": userUID,
	})
}
