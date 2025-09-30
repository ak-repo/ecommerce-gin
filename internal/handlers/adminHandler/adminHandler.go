package adminhandler

import (
	"net/http"
	"strconv"

	"github.com/ak-repo/ecommerce-gin/internal/dto"
	adminservice "github.com/ak-repo/ecommerce-gin/internal/services/adminService"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService adminservice.AdminService
}

func NewAdminHandler(adminService adminservice.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

func (h *AdminHandler) AdminLoginFormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/admin/auth/adminLogin.html", gin.H{
		"title": "Admin Login",
	})
}

func (h *AdminHandler) AdminLoginHandler(ctx *gin.Context) {

	var input dto.LoginRequest
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusUnauthorized, "pages/admin/auth/adminLogin.html", gin.H{
			"error": err,
			"title": "Admin Login Error",
		})
		return
	}

	res, err := h.adminService.AdminLoginService(input.Email, input.Password)
	if err != nil {
		ctx.HTML(http.StatusUnauthorized, "pages/admin/auth/adminLogin.html", gin.H{
			"error": err,
			"title": "Admin Login Error",
		})
		return

	}
	// Set secure cookies
	ctx.SetCookie("accessToken", res.AccessToken, int(res.AccessExp), "/", "localhost", true, true)
	ctx.SetCookie("refreshToken", res.RefreshToken, int(res.RefreshExp), "/", "localhost", true, true)

	ctx.Redirect(http.StatusSeeOther, "/admin/dashboard")

}

// GET Logout
func (h *AdminHandler) AdminLogout(ctx *gin.Context) {

	ctx.SetCookie("accessToken", "", 0, "/", "localhost", true, true)
	ctx.SetCookie("refreshToken", "", 0, "/", "localhost", true, true)

	ctx.Redirect(http.StatusSeeOther, "/admin/login")

}

// ----------------------------------------------------------
// GET admin/dashboard
func (h *AdminHandler) AdminDashboardForm(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/admin/dashbord/dashbord.html", gin.H{})
}

// GET admin/profile
func (h *AdminHandler) AdminProfileHandler(ctx *gin.Context) {

	email, exists := ctx.Get("email")
	if !exists || email == "" {
		ctx.HTML(http.StatusBadRequest, "pages/admin/profile/profile.html", gin.H{
			"Error": "no email on token",
		})
		return
	}
	emailstr := email.(string)
	profile, err := h.adminService.AdminProfileService(emailstr)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/admin/profile/profile.html", gin.H{
			"Error": err.Error(),
		})
		return

	}
	ctx.HTML(http.StatusOK, "pages/admin/profile/profile.html", gin.H{
		"Profile": profile,
		"Error":   nil,
	})
}

// GET user/address -> shop user address form
func (h *AdminHandler) ShowAddressFormHandler(ctx *gin.Context) {
	addressID := ctx.Param("address_id")
	email, _ := ctx.Get("email")
	emailStr := email.(string)

	address := dto.AddressDTO{}
	if addressID == "0" {
		ctx.HTML(http.StatusOK, "pages/admin/profile/address.html", gin.H{
			"Address": address,
			"User":    email,
		})
		return

	}

	profile, err := h.adminService.AdminProfileService(emailStr)
	if err != nil {
		ctx.HTML(http.StatusOK, "pages/admin/profile/address.html", gin.H{
			"Address": address,
			"User":    email,
			"Error":   err.Error(),
		})
		return
	}

	addID, _ := strconv.ParseUint(addressID, 10, 64)
	if profile.Address.ID == 0 || uint(addID) != profile.Address.ID {
		ctx.HTML(http.StatusOK, "pages/admin/profile/address.html", gin.H{
			"Address": nil,
			"User":    email,
		})
		return

	}

	ctx.HTML(http.StatusOK, "pages/admin/profile/address.html", gin.H{
		"User":    email,
		"Address": profile.Address,
	})
}

func (h *AdminHandler) UpdateAddressHandler(ctx *gin.Context) {
	addressID := ctx.Param("address_id")
	email, _ := ctx.Get("email")
	emailStr := email.(string)

	var address dto.AddressDTO
	if err := ctx.ShouldBind(&address); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/admin/profile/address.html", gin.H{
			"Error":   err.Error(),
			"Address": address,
		})
		return
	}

	if err := h.adminService.AdminAddressUpdateService(emailStr, addressID, &address); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/admin/profile/address.html", gin.H{
			"Error":   err.Error(),
			"Address": address,
		})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/profile")

}
