package boarddto

// Chart data structures
type LineChartData struct {
	Dates   []string  `json:"dates"`
	Amounts []float64 `json:"amounts"`
}

type PieChartData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

// Main Dashboard struct
type DashboardData struct {
	TotalRevenue   float64 `json:"totalRevenue"`
	TotalProducts  int     `json:"totalProducts"`
	TotalOrders    int     `json:"totalOrders"`
	PendingOrders  int     `json:"pendingOrders"`
	TotalCustomers int     `json:"totalCustomers"`
	CurrentSection string  `json:"currentSection"`

	RevenueChart LineChartData `json:"revenueChart"`
	OrdersChart  PieChartData  `json:"ordersChart"`
}
