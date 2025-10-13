package profilehandler

import (
	"net/http"

	profiledto "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/profile_dto"
	profileinterface "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/profile_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ProfileService profileinterface.Service
}

func NewProfileHandlerMG(profileService profileinterface.Service) profileinterface.Handler {
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

	role, _ := ctx.Get("role")
	if role == "admin" {
		ctx.HTML(http.StatusOK, "pages/profile/profile.html", gin.H{
			"Profile": profile,
		})
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "user profile found", map[string]interface{}{
		"Profile": profile,
	})

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

	role, _ := ctx.Get("role")
	if role == "admin" {
		ctx.HTML(http.StatusOK, "pages/profile/address.html", gin.H{
			"Address": address,
		})
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "user address", map[string]interface{}{
		"address": address,
	})

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

	role, _ := ctx.Get("role")
	if role == "admin" {
		ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/profile")
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "address updated", nil)
}
