package boardinter

import (
	boarddto "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_mg/dashboard_dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	DashboardOverview(ctx *gin.Context)
}

type Service interface {
	DashboardOverview() (*boarddto.DashboardData, error)
}

type Repository interface {
	TotalCount(model interface{}) (int64, error)
	TotalRevenue() (float64, error)
	Orders() ([]models.Order, error)
}
