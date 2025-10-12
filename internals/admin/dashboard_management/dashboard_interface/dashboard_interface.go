package dashboardinterface

import (
	dashboarddto "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_management/dashboard_dto"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	DashboardOverview(ctx *gin.Context)
}

type Service interface {
	DashboardOverview() (*dashboarddto.DashboardData, error)
}
