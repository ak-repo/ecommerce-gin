package admindashhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminDashboardHandler struct {
}

func NewAdminDashboardHandler() *AdminDashboardHandler {
	return &AdminDashboardHandler{}

}

func (h *AdminDashboardHandler) AdminDashboardShow(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/dashboard/dashboard.html", gin.H{})
}
