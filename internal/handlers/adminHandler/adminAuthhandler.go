package adminhandler

import (
	"net/http"

	"github.com/ak-repo/ecommerce-gin/internal/models"
	adminservice "github.com/ak-repo/ecommerce-gin/internal/services/adminService"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminAuthService adminservice.AdminAuthService
}

func NewAdminHandler(adminAuthService adminservice.AdminAuthService) *AdminHandler {
	return &AdminHandler{adminAuthService: adminAuthService}
}

func (h *AdminHandler) AdminLoginFormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/admin/auth/adminLogin.html", gin.H{
		"title": "Admin Login",
	})
}

func (h *AdminHandler) AdminLoginHandler(ctx *gin.Context) {

	input := models.InputUser{}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusUnauthorized, "pages/admin/auth/adminLogin.html", gin.H{
			"error": err,
			"title": "Admin Login Error",
		})
		return
	}

	res, err := h.adminAuthService.AdminLoginService(&input)
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

// GET admin/dashboard
func (h *AdminHandler) AdminDashboardForm(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/admin/dashbord/dashbord.html", gin.H{})
}
