package profilehandler

import (
	"net/http"

	profiledto "github.com/ak-repo/ecommerce-gin/internals/customer/profile/profile_dto"
	profileinterface "github.com/ak-repo/ecommerce-gin/internals/customer/profile/profile_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ProfileService profileinterface.Service
}

func NewProfileHandler(profileService profileinterface.Service) profileinterface.Handler {
	return &handler{ProfileService: profileService}
}

func (h *handler) GetProfile(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}

	profile, err := h.ProfileService.GetProfile(userID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "DB issue", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "user profile found", profile)

}

func (h *handler) GetAddress(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}

	address, err := h.ProfileService.GetAddress(userID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "address not found", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "user address", address)

}

func (h *handler) UpdateAddress(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}

	var address profiledto.AddressDTO
	if err := ctx.ShouldBind(&address); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "binding isssue", err)
		return
	}

	if err := h.ProfileService.UpdateAddress(&address, userID); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "address update failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "address updated", nil)
}
