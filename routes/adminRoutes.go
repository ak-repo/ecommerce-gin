package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	profilehandler "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile/profile_handler"
	profilerepo "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile/profile_repo"
	profileservice "github.com/ak-repo/ecommerce-gin/internals/admin/admin_profile/profile_service"
	admindashhandler "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_management/admin_dash_handler"
	adminorderhandler "github.com/ak-repo/ecommerce-gin/internals/admin/orders_management/admin_order_handler"
	adminproducthandler "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_handler"
	adminproductrepo "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_repo"
	adminproductservice "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_service"
	reviewmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/review_management"
	adminuserhandler "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_handler"
	adminuserrepo "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_repo"
	adminuserservice "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_service"
	authhandler "github.com/ak-repo/ecommerce-gin/internals/auth/auth_handler"
	authrepo "github.com/ak-repo/ecommerce-gin/internals/auth/auth_repo"
	authservice "github.com/ak-repo/ecommerce-gin/internals/auth/auth_service"
	orderrepos "github.com/ak-repo/ecommerce-gin/internals/order/order_repos"
	orderservices "github.com/ak-repo/ecommerce-gin/internals/order/order_services"
	reviewrepo "github.com/ak-repo/ecommerce-gin/internals/review/review_repo"
	reviewservice "github.com/ak-repo/ecommerce-gin/internals/review/review_service"
	middleware "github.com/ak-repo/ecommerce-gin/middleware/auth"

	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	// login
	authRepo := authrepo.NewAuthRepo(db.DB)
	authService := authservice.NewAuthService(authRepo, cfg)
	authHandler := authhandler.NewAuthHandler(authService)

	r.GET("/login", authHandler.AdminLoginForm)
	r.POST("/login", authHandler.AdminLogin)

	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.AuthMiddleware(cfg), middleware.RoleMiddleware("admin"))

	{
		//  auth
		adminRoute.GET("/password-change", authHandler.AdminPasswordChange)
		adminRoute.POST("/password-change", authHandler.AdminPasswordChange)

		// dashbord
		dashboardHandler := admindashhandler.NewAdminDashboardHandler()
		adminRoute.GET("/dashboard", dashboardHandler.AdminDashboardShow)

		// products management
		productRepo := adminproductrepo.NewAdminProductRepo(db.DB)
		productService := adminproductservice.NewAdminProductService(productRepo)
		productHandler := adminproducthandler.NewAdminProductHandler(productService)

		adminRoute.GET("/products", productHandler.AdminAllProducstListHandler)
		adminRoute.GET("/products/:id", productHandler.AdminProductListHandler)
		adminRoute.GET("/products/add", productHandler.AdminAddProductPageHandler)
		adminRoute.POST("/products/add", productHandler.AdminAddProductHandler)
		adminRoute.GET("/products/edit/:id", productHandler.AdminShowProductEditPageHandler)
		adminRoute.POST("/products/edit/:id", productHandler.AdminUpdateProductHandler)

		// orders management
		orderRepo := orderrepos.NewOrderRepo(db.DB)
		orderService := orderservices.NewOrderService(orderRepo)
		orderHandler := adminorderhandler.NewAdminOrderHandler(orderService)

		adminRoute.GET("/orders", orderHandler.ShowAllOrderHandler)
		adminRoute.GET("/orders/:id", orderHandler.ShowOrderByIDHandler)
		adminRoute.POST("/orders/status/:id", orderHandler.UpdateOrderStatusHandler)
		adminRoute.GET("/orders/cancels", orderHandler.ShowAllCancelRequestHandler)
		adminRoute.POST("/orders/cancels/accept/:id", orderHandler.OrderCancellationAcceptHandler)
		adminRoute.POST("/orders/cancels/reject/:id", orderHandler.OrderCancellationRejectHandler)

		// user management
		userRepo := adminuserrepo.NewAdminUserRepo(db.DB)
		userService := adminuserservice.NewAdminUserService(userRepo)
		userHandler := adminuserhandler.NewAdminUserHandler(userService)

		adminRoute.GET("/users", userHandler.ListAllUsersHandler)
		adminRoute.GET("/users/:id", userHandler.ListUserByIDHandler)
		adminRoute.POST("/users/role/:id", userHandler.AdminUserRoleChangeHandler)
		adminRoute.POST("/users/status/:id", userHandler.AdminUserBlockHandler)

		// admin profile
		profileRepo := profilerepo.NewAdminProfileRepo(db.DB)
		profileSevice := profileservice.NewAdminProfileService(authRepo, profileRepo)
		profileHandle := profilehandler.NewAdminProfileHandler(profileSevice)

		adminRoute.GET("/profile", profileHandle.AdminProfileHandler)
		adminRoute.GET("/address/:id", profileHandle.ShowAddressFormHandler)
		adminRoute.POST("/address/update/:id", profileHandle.UpdateAddressHandler)

		// reviews and rating management
		reviewRepo := reviewrepo.NewReviewRepo(db.DB)
		reviewService := reviewservice.NewReviewService(reviewRepo)
		reviewHandle := reviewmanagement.NewAdminReviewService(reviewService)

		adminRoute.GET("/reviews", reviewHandle.ListAllReviewsHandler)
		adminRoute.POST("/reviews/approve/:id", reviewHandle.ApproveReviewHandler)
		adminRoute.POST("/reviews/reject/:id", reviewHandle.RejectReviewHandler)
	}

}
