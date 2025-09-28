package adminhandler

import (
	"net/http"

	"github.com/ak-repo/ecommerce-gin/internal/models"
	adminservice "github.com/ak-repo/ecommerce-gin/internal/services/adminService"
	"github.com/gin-gonic/gin"
)

type AdminAuthHandler struct {
	adminAuthService adminservice.AdminAuthService
}

func NewAdminAuthHandler(adminAuthService adminservice.AdminAuthService) *AdminAuthHandler {
	return &AdminAuthHandler{adminAuthService: adminAuthService}
}

func (h *AdminAuthHandler) AdminLoginFormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/admin/auth/adminLogin.html", gin.H{
		"title": "Admin Login",
	})
}

func (h *AdminAuthHandler) AdminLoginHandler(ctx *gin.Context) {

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

// GET admin/dashboard
func (h *AdminAuthHandler) AdminDashboardForm(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/admin/dashbord/dashbord.html", gin.H{})
}
