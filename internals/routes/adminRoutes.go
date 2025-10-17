package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"

	db "github.com/ak-repo/ecommerce-gin/config/database"
	bannerhandler "github.com/ak-repo/ecommerce-gin/internals/admin/banner_mg/handler"
	bannerrepo "github.com/ak-repo/ecommerce-gin/internals/admin/banner_mg/repo"
	bannerservice "github.com/ak-repo/ecommerce-gin/internals/admin/banner_mg/service"
	categoryhandler "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/handler"
	categoryrepository "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/repository"
	categoryservice "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/service"

	boardhandler "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_mg/handler"
	boardrepo "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_mg/repo"
	boardservice "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_mg/service"
	orderhandler "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/handler"
	orderrepo "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/repo"
	orderservice "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/service"
	producthandler "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/handler"
	productrepo "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/repository"
	productservice "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/service"
	profilehandler "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/handler"
	profilerepository "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/repository"
	profileservice "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/service"
	reviewhandler "github.com/ak-repo/ecommerce-gin/internals/admin/review_management/handler"
	reviewrepo "github.com/ak-repo/ecommerce-gin/internals/admin/review_management/repository"
	reviewsvc "github.com/ak-repo/ecommerce-gin/internals/admin/review_management/service"
	usershandler "github.com/ak-repo/ecommerce-gin/internals/admin/users_mg/handler"
	usersrepository "github.com/ak-repo/ecommerce-gin/internals/admin/users_mg/repository"
	usersservice "github.com/ak-repo/ecommerce-gin/internals/admin/users_mg/service"

	authhandler "github.com/ak-repo/ecommerce-gin/internals/auth/auth_handler"
	authrepo "github.com/ak-repo/ecommerce-gin/internals/auth/auth_repo"
	authservice "github.com/ak-repo/ecommerce-gin/internals/auth/auth_service"
	authmiddleware "github.com/ak-repo/ecommerce-gin/pkg/middleware/auth"

	"github.com/gin-gonic/gin"
)

// RegisterAdminRoutes defines all routes for admin functionalities
func RegisterAdminRoutes(r *gin.Engine, db *db.Database, cfg *config.Config) {
	// --------------------------
	// INIT CORE REPOSITORIES
	// --------------------------
	authRepo := authrepo.NewAuthRepo(db.DB)
	orderRepo := orderrepo.NewOrderRepositoryMG(db.DB)
	userRepo := usersrepository.NewrUsersRpository(db.DB)
	productRepo := productrepo.NewProductRepoMG(db.DB)
	categoryRepo := categoryrepository.NewCategoryRepo(db.DB)
	reviewRepo := reviewrepo.NewReviewRepoMG(db.DB)
	profileRepo := profilerepository.NewProfileRepoMG(db.DB)
	bannerRepo := bannerrepo.NewBannerRepoMG(db.DB)
	boardRepo := boardrepo.NewNewDashRepo(db.DB)

	// --------------------------
	// INIT SERVICES
	// --------------------------
	authService := authservice.NewAuthService(authRepo, cfg)
	orderService := orderservice.NewOrderServiceMG(orderRepo)
	userService := usersservice.NewUsersService(userRepo)
	productService := productservice.NewProductServiceMG(productRepo)
	categoryService := categoryservice.NewCategoryService(categoryRepo)
	reviewService := reviewsvc.NewReviewServiceMG(reviewRepo)
	profileService := profileservice.NewProfileServiceMG(profileRepo, authRepo)
	boardService := boardservice.NewDashboardService(boardRepo)
	bannerService := bannerservice.NewBannerServiceMG(bannerRepo)

	// --------------------------
	// INIT HANDLERS
	// --------------------------
	authHandler := authhandler.NewAuthHandler(authService)
	orderHandler := orderhandler.NewOrderHandlerMG(orderService)
	userHandler := usershandler.NewAdminUserHandler(userService)
	productHandler := producthandler.NewProductHandlerMG(productService, categoryService)
	categoryHandler := categoryhandler.NewCategoryHandler(categoryService)
	reviewHandler := reviewhandler.NewReviewHandlerMG(reviewService)
	profileHandler := profilehandler.NewProfileHandlerMG(profileService)
	dashboardHandler := boardhandler.NewAdminDashboardHandler(boardService)
	bannerHandler := bannerhandler.NewBannerHandlerMG(bannerService)

	// PUBLIC ADMIN ROUTES (no middleware)
	public := r.Group("/api/v1/admin")
	{
		public.GET("/login", authHandler.AdminLoginForm)
		public.POST("/login", authHandler.AdminLogin)
	}

	// PROTECTED ADMIN ROUTES
	protected := public.Group("")
	protected.Use(authmiddleware.AuthMiddleware(cfg), authmiddleware.RoleMiddleware("admin"))
	{
		// --------------------------
		// AUTH MANAGEMENT
		// --------------------------
		protected.POST("/password-change", authHandler.AdminPasswordChange)
		protected.GET("/password-change", authHandler.AdminPasswordChange)
		protected.GET("/logout", authHandler.AdminLogout)

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
		protected.GET("/products/delete/:id", productHandler.DeleteProduct)

		// --------------------------
		// ORDERS MANAGEMENT
		// --------------------------
		protected.GET("/orders", orderHandler.GetAllOrders)
		protected.GET("/orders/:id", orderHandler.GetOrderByID)
		protected.POST("/orders/status/:id", orderHandler.UpdateStatus)
		protected.GET("/orders/cancel-requests", orderHandler.GetAllCancels)
		protected.POST("/orders/cancel-requests/:id/accept", orderHandler.AcceptCancel)
		protected.POST("/orders/cancel-requests/:id/reject", orderHandler.RejectCancel)

		// --------------------------
		// USERS MANAGEMENT
		// --------------------------
		protected.GET("/users", userHandler.GetAllUsers)
		protected.GET("/users/:id", userHandler.GetUserByID)
		protected.GET("/users/add", userHandler.ShowUserAddForm)
		protected.POST("/users/add", userHandler.CreateUser)
		protected.POST("/users/:id/role", userHandler.ChangeUserRole)
		protected.POST("/users/:id/status", userHandler.BlockUser)
		protected.GET("/users/:id/orders", orderHandler.GetOrderByCustomerID)
		protected.GET("/users/delete/:id", userHandler.DeleteUser)

		// --------------------------
		// PROFILE MANAGEMENT
		// --------------------------
		protected.GET("/profile", profileHandler.GetProfile)
		protected.GET("/profile/address/:id", profileHandler.GetAddress)
		protected.POST("/profile/address/:id", profileHandler.UpdateAddress)
		protected.POST("/profile/profile_pic", profileHandler.UploadPicture)

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

		// --------------------------
		// BANNER MANAGEMENT
		// --------------------------
		protected.GET("/banners", bannerHandler.GetAllBanners)
		protected.GET("/banners/:id", bannerHandler.GetBannerByID)
		protected.GET("/banners/add", bannerHandler.CreateForm)
		protected.POST("/banners/add", bannerHandler.Create)
		protected.GET("/banners/:id/update", bannerHandler.UpdateForm)
		protected.POST("/banners/:id/update", bannerHandler.Update)
		protected.GET("/banners/:id/delete", bannerHandler.Delete)

	}
}
