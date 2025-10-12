package dashboardhandler

import (
	"net/http"

	dashboardinterface "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_management/dashboard_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type adminDashboardHandler struct {
	DashboardService dashboardinterface.Service
}

func NewAdminDashboardHandler(service dashboardinterface.Service) dashboardinterface.Handler {
	return &adminDashboardHandler{DashboardService: service}

}

func (h *adminDashboardHandler) DashboardOverview(ctx *gin.Context) {

	dashborad, err := h.DashboardService.DashboardOverview()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "no dashbord data found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/dashboard/dashboard.html", dashborad)
}
