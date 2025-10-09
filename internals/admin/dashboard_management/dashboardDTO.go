package dashboardmanagement

type DashboardData struct {
	TotalRevenue   float64
	TotalProducts  int
	TotalOrders    int
	PendingOrders  int
	TotalCustomers int
	RecentOrders   []OrderSummary
	TopProducts    []ProductSummary
	CurrentSection string
	RevenueChart   RevenueChartData
}

type RevenueChartData struct {
	Dates   []string  `json:"dates"`
	Amounts []float64 `json:"amounts"`
}
type OrderSummary struct {
	OrderID uint

	CustomerName string
	ItemsCount   int
	TotalAmount  float64
	Status       string // e.g., "Completed", "Pending", "Cancelled"
}

type ProductSummary struct {
	Name     string
	Category string
	Price    float64
	InStock  bool
}
