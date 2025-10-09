package admindashhandler

import (
	"net/http"

	"github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_management/service"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminDashboardHandler struct {
	Service service.AdminDashboardService
}

func NewAdminDashboardHandler(service service.AdminDashboardService) *AdminDashboardHandler {
	return &AdminDashboardHandler{Service: service}

}

func (h *AdminDashboardHandler) AdminDashboardShow(ctx *gin.Context) {

	dashborad, err := h.Service.DashboardService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "no dashbord data found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/dashboard/dashboard.html", dashborad)
}
