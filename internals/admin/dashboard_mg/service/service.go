package boardservice

import (
	boarddto "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_mg/dashboard_dto"
	boardinter "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_mg/dashboard_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
)

type service struct {
	repo boardinter.Repository
}

func NewDashboardService(repo boardinter.Repository) boardinter.Service {
	return &service{repo: repo}
}

func (s *service) DashboardOverview() (*boarddto.DashboardData, error) {

	productCount, err := s.repo.TotalCount(&models.Product{})
	if err != nil {
		return nil, err
	}
	users, err := s.repo.TotalCount(&models.User{})
	if err != nil {
		return nil, err
	}
	revenue, err := s.repo.TotalRevenue()
	if err != nil {
		return nil, err
	}

	orders, err := s.repo.Orders()
	if err != nil {
		return nil, err
	}

	var pending float64
	var completed float64
	var cancelled float64

	for _, v := range orders {

		switch v.Status {
		case "pending":
			pending++
		case "completed":
			completed++
		case "cancelled":
			cancelled++
		}

	}

	dashboardData := boarddto.DashboardData{
		TotalRevenue:   revenue,
		TotalProducts:  int(productCount),
		TotalOrders:    len(orders),
		PendingOrders:  int(pending),
		TotalCustomers: int(users),
		CurrentSection: "dashboard",
		RevenueChart: boarddto.LineChartData{
			Dates:   []string{"Oct 9", "Oct 10", "Oct 11", "Oct 12", "Oct 13", "Oct 14", "Oct 15"},
			Amounts: []float64{1200, 1500, 900, 1800, 2000, 1750, 2100},
		},
		OrdersChart: boarddto.PieChartData{
			Labels: []string{"Completed", "Pending", "Cancelled"},
			Values: []float64{completed, pending, cancelled},
		},
	}

	return &dashboardData, nil

}
