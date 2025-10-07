package profilehandler

import (
	"errors"
	"net/http"
	"strconv"

	adminprofile "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile"
	profileinterface "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile/profile_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminProfileHandler struct {
	ProfileService profileinterface.ServiceInterface
}

func NewAdminProfileHandler(profileService profileinterface.ServiceInterface) profileinterface.HandlerInterface {
	return &AdminProfileHandler{ProfileService: profileService}
}

// GET admin/profile
func (h *AdminProfileHandler) AdminProfileHandler(ctx *gin.Context) {

	id, exists := ctx.Get("userID")
	if !exists || id == "" {
		utils.RenderError(ctx, http.StatusUnauthorized, "admin", "unauthorised request", errors.New("no user id found"))
		return
	}
	profile, err := h.ProfileService.AdminProfileService(id.(uint))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "profile service failed", err)
		return

	}
	ctx.HTML(http.StatusOK, "pages/profile/profile.html", gin.H{
		"Profile": profile,
		"Error":   nil,
	})
}

// GET admin/address -> shop user address form
func (h *AdminProfileHandler) ShowAddressFormHandler(ctx *gin.Context) {
	addressID := ctx.Param("id")
	res := adminprofile.AddressDTO{}
	if addressID == "0" {
		ctx.HTML(http.StatusOK, "pages/profile/address.html", gin.H{
			"Address": res,
		})
		return

	}
	userID, exists := ctx.Get("userID")
	if !exists || userID == "" {
		utils.RenderError(ctx, http.StatusUnauthorized, "admin", "unauthorised request", errors.New("no user id found"))
		return
	}

	id, err := strconv.ParseUint(addressID, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "invalid id", err)
		return
	}
	address, err := h.ProfileService.GetAdminAddressService(uint(id))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "address is not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/profile/address.html", gin.H{
		"Address": address,
	})
}

// POST admin/address/update/:id
func (h *AdminProfileHandler) UpdateAddressHandler(ctx *gin.Context) {
	addressID := ctx.Param("id")
	var req adminprofile.AddressDTO
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", errors.New("no user id found"))
		return
	}
	id, err := strconv.ParseUint(addressID, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "invalid id", err)
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists || userID == "" {
		utils.RenderError(ctx, http.StatusUnauthorized, "admin", "unauthorised request", errors.New("no user id found"))
		return
	}

	if err := h.ProfileService.AdminAddressUpdateService(userID.(uint), uint(id), &req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "address is not found", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/profile")

}
