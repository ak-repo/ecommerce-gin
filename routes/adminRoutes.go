package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	categoryhandler "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/handler"
	categoryrepository "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/repository"
	categoryservice "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/service"
	dashboardhandler "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_management/handler"
	dashboardservice "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_management/service"
	productmanagementhandler "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/handler"
	productmanagementrepository "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/repository"
	productmanagementservice "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/service"
	usershandler "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/handler"
	usersrepository "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/repository"
	usersservice "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/service"
	authhandler "github.com/ak-repo/ecommerce-gin/internals/auth/auth_handler"
	authrepo "github.com/ak-repo/ecommerce-gin/internals/auth/auth_repo"
	authservice "github.com/ak-repo/ecommerce-gin/internals/auth/auth_service"
	orderhandler "github.com/ak-repo/ecommerce-gin/internals/order/handler"
	orderrepository "github.com/ak-repo/ecommerce-gin/internals/order/order_repos"
	orderservices "github.com/ak-repo/ecommerce-gin/internals/order/order_services"
	profilehandler "github.com/ak-repo/ecommerce-gin/internals/profile/handler"
	profilerepository "github.com/ak-repo/ecommerce-gin/internals/profile/repository"
	profileservice "github.com/ak-repo/ecommerce-gin/internals/profile/service"
	reviewhandler "github.com/ak-repo/ecommerce-gin/internals/review/handler"
	reviewrepository "github.com/ak-repo/ecommerce-gin/internals/review/repository"
	reviewservice "github.com/ak-repo/ecommerce-gin/internals/review/service"
	middleware "github.com/ak-repo/ecommerce-gin/middleware/auth"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

// RegisterAdminRoutes defines all routes for admin functionalities
func RegisterAdminRoutes(r *gin.Engine, db *db.Database, cfg *config.Config) {
	// --------------------------
	// INIT CORE REPOSITORIES
	// --------------------------
	authRepo := authrepo.NewAuthRepo(db.DB)
	orderRepo := orderrepository.Newrepository(db.DB)
	userRepo := usersrepository.NewrUsersRpository(db.DB)
	productRepo := productmanagementrepository.NewProductRepo(db.DB)
	categoryRepo := categoryrepository.NewCategoryRepo(db.DB)
	reviewRepo := reviewrepository.NewReviewRepo(db.DB)
	profileRepo := profilerepository.NewProfileRepository(db.DB)

	// --------------------------
	// INIT SERVICES
	// --------------------------
	authService := authservice.NewAuthService(authRepo, cfg)
	orderService := orderservices.NewOrderService(orderRepo)
	userService := usersservice.NewUsersService(userRepo)
	productService := productmanagementservice.Newservice(productRepo)
	categoryService := categoryservice.NewCategoryService(categoryRepo)
	reviewService := reviewservice.NewReviewService(reviewRepo)
	profileService := profileservice.NewProfileService(profileRepo, authRepo)
	dashboardService := dashboardservice.NewDashboardService(orderRepo, productRepo, userRepo)

	// --------------------------
	// INIT HANDLERS
	// --------------------------
	authHandler := authhandler.NewAuthHandler(authService)
	orderHandler := orderhandler.NewOrderHandler(orderService)
	userHandler := usershandler.NewAdminUserHandler(userService)
	productHandler := productmanagementhandler.Newhandler(productService)
	categoryHandler := categoryhandler.NewCategoryHandler(categoryService)
	reviewHandler := reviewhandler.NewReviewHandler(reviewService)
	profileHandler := profilehandler.NewProfileHandler(profileService)
	dashboardHandler := dashboardhandler.NewAdminDashboardHandler(&dashboardService)

	// PUBLIC ADMIN ROUTES (no middleware)
	public := r.Group("/api/v1/admin")
	{
		public.GET("/login", authHandler.AdminLoginForm)
		public.POST("/login", authHandler.AdminLogin)
	}

	// PROTECTED ADMIN ROUTES
	protected := public.Group("")
	protected.Use(middleware.AuthMiddleware(cfg), middleware.RoleMiddleware("admin"))
	{
		// --------------------------
		// AUTH MANAGEMENT
		// --------------------------
		protected.POST("/password-change", authHandler.AdminPasswordChange)
		protected.GET("/password-change", authHandler.AdminPasswordChange)
		protected.POST("/logout", authHandler.AdminLogout)

		// --------------------------
		// DASHBOARD
		// --------------------------
		protected.GET("/dashboard", dashboardHandler.DashboardOverview)

		// --------------------------
		// PRODUCTS MANAGEMENT
		// --------------------------
		protected.GET("/products", productHandler.GetAllProducts)
		protected.GET("/products/:id", productHandler.GetProductByID)
		protected.GET("/products/add", productHandler.AddProductForm)
		protected.POST("/products/add", productHandler.AddProduct)
		protected.GET("/products/update/:id", productHandler.EditProductForm)
		protected.POST("/products/update/:id", productHandler.UpdateProduct)

		// --------------------------
		// ORDERS MANAGEMENT
		// --------------------------
		protected.GET("/orders", orderHandler.GetAllOrders)
		protected.GET("/orders/:id", orderHandler.GetOrderByIDForAdmin)
		protected.POST("/orders/status/:id", orderHandler.UpdateStatus)
		protected.GET("/orders/cancel-requests", orderHandler.GetAllCancels)
		protected.POST("/orders/cancel-requests/:id/accept", orderHandler.AcceptCancel)
		protected.POST("/orders/cancel-requests/:id/reject", orderHandler.RejectCancel)

		// need for change
		protected.GET("/users/:id/orders", orderHandler.GetOrderByCustomerIDForCustomer)

		// --------------------------
		// USERS MANAGEMENT
		// --------------------------
		protected.GET("/users", userHandler.GetAllUsers)
		protected.GET("/users/:id", userHandler.GetUserByID)
		protected.GET("/users/add", userHandler.ShowUserAddForm)
		protected.POST("/users/add", userHandler.CreateUser)
		protected.POST("/users/:id/role", userHandler.ChangeUserRole)
		protected.POST("/users/:id/status", userHandler.BlockUser)

		// --------------------------
		// PROFILE MANAGEMENT
		// --------------------------
		protected.GET("/profile", profileHandler.GetProfile)
		protected.GET("/profile/address/:id", profileHandler.GetAddress)
		protected.POST("/profile/address/:id", profileHandler.UpdateAddress)

		// --------------------------
		// REVIEWS MANAGEMENT
		// --------------------------
		protected.GET("/reviews", reviewHandler.GetAllReviews)
		protected.POST("/reviews/:id/approve", reviewHandler.ApporveReview)
		protected.POST("/reviews/:id/reject", reviewHandler.RejectReview)

		// --------------------------
		// CATEGORIES MANAGEMENT
		// --------------------------
		protected.GET("/categories", categoryHandler.GetAllCategories)
		protected.GET("/categories/:id", categoryHandler.GetCategoryDetails)
		protected.POST("/categories", categoryHandler.CreateCategory)
	}
}
