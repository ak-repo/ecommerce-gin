package service

import (
	dashboardmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_management"
	adminproductinterface "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_interface"
	adminuserinterface "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_interface"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/order/order_interface"
)

type AdminDashboardService struct {
	Orders   orderinterface.OrderRepoInterface
	Products adminproductinterface.RepoInterface
	Users    adminuserinterface.RepoInterface
}

func NewAdminDashboardService(order orderinterface.OrderRepoInterface, product adminproductinterface.RepoInterface, user adminuserinterface.RepoInterface) AdminDashboardService {
	return AdminDashboardService{Orders: order, Users: user, Products: product}
}

func (s *AdminDashboardService) DashboardService() (*dashboardmanagement.DashboardData, error) {

	// orders
	orders, err := s.Orders.GetAllOrders()
	if err != nil {
		return nil, err
	}
	products, err := s.Products.GetAllProducts()
	if err != nil {
		return nil, err
	}
	users, err := s.Users.AdminGetAllUsers()
	if err != nil {
		return nil, err
	}

	dashboardData := dashboardmanagement.DashboardData{
		TotalProducts:  len(products),
		TotalOrders:    len(orders),
		TotalCustomers: len(users),
		RevenueChart: dashboardmanagement.RevenueChartData{
			Dates:   []string{"2025-10-01", "2025-10-02", "2025-10-03", "2025-10-04", "2025-10-05", "2025-10-06", "2025-10-07"},
			Amounts: []float64{1200, 1500, 900, 1800, 1300, 1700, 1600},
		},
	}
	// orders
	var totalRevenue float64

	var recentOrders []dashboardmanagement.OrderSummary
	for _, order := range orders {
		// recent ordseer
		if len(recentOrders) < 6 {
			item := dashboardmanagement.OrderSummary{
				OrderID:      order.ID,
				CustomerName: order.User.Username,
				ItemsCount:   len(order.OrderItems),
				TotalAmount:  order.TotalAmount,
				Status:       order.Status,
			}
			recentOrders = append(recentOrders, item)
		}

		// pending
		if order.Status == "pending" {
			dashboardData.PendingOrders++
		}

		if order.Status == "completed" {
			totalRevenue += order.TotalAmount
		}
	}

	// products
	var topProducts []dashboardmanagement.ProductSummary
	for _, product := range products {
		if len(topProducts) < 6 {
			item := dashboardmanagement.ProductSummary{
				Name:     product.Title,
				Category: product.Category.Name,
				Price:    product.BasePrice,
				InStock:  product.IsActive,
			}
			topProducts = append(topProducts, item)

		}

	}

	dashboardData.TotalRevenue = totalRevenue

	dashboardData.TopProducts = topProducts
	dashboardData.RecentOrders = recentOrders
	dashboardData.CurrentSection = "dashboard"

	return &dashboardData, nil

}
