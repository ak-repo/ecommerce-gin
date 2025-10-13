package dashboardservice

import (
	dashboarddto "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_management/dashboard_dto"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/order_interface"
	usersinterface "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/user_interface"
	productinterface "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_interface"
)

type adminDashboardService struct {
	Orders   orderinterface.Repository
	Products productinterface.Repository
	Users    usersinterface.Repository
}

func NewDashboardService(order orderinterface.Repository, product productinterface.Repository, user usersinterface.Repository) adminDashboardService {
	return adminDashboardService{Orders: order, Users: user, Products: product}
}

func (s *adminDashboardService) DashboardOverview() (*dashboarddto.DashboardData, error) {

	// orders
	orders, err := s.Orders.GetAllOrders()
	if err != nil {
		return nil, err
	}
	products, err := s.Products.GetAllProducts()
	if err != nil {
		return nil, err
	}
	users, err := s.Users.GetAllUsers()
	if err != nil {
		return nil, err
	}

	dashboardData := dashboarddto.DashboardData{
		TotalProducts:  len(products),
		TotalOrders:    len(orders),
		TotalCustomers: len(users),
		RevenueChart: dashboarddto.RevenueChartData{
			Dates:   []string{"2025-10-01", "2025-10-02", "2025-10-03", "2025-10-04", "2025-10-05", "2025-10-06", "2025-10-07"},
			Amounts: []float64{1200, 1500, 900, 1800, 1300, 1700, 1600},
		},
	}
	// orders
	var totalRevenue float64

	var recentOrders []dashboarddto.OrderSummary
	for _, order := range orders {
		// recent ordseer
		if len(recentOrders) < 6 {
			item := dashboarddto.OrderSummary{
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
	var topProducts []dashboarddto.ProductSummary
	for _, product := range products {
		if len(topProducts) < 6 {
			item := dashboarddto.ProductSummary{
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
