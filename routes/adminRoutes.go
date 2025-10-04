package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	admindashhandler "github.com/ak-repo/ecommerce-gin/internals/admin/dashboard_management/admin_dash_handler"
	adminorderhandler "github.com/ak-repo/ecommerce-gin/internals/admin/orders_management/admin_order_handler"
	adminproducthandler "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_handler"
	adminproductrepo "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_repo"
	adminproductservice "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_service"
	adminuserhandler "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_handler"
	adminuserrepo "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_repo"
	adminuserservice "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_service"
	authhandler "github.com/ak-repo/ecommerce-gin/internals/auth/auth_handler"
	authrepo "github.com/ak-repo/ecommerce-gin/internals/auth/auth_repo"
	authservice "github.com/ak-repo/ecommerce-gin/internals/auth/auth_service"
	orderrepos "github.com/ak-repo/ecommerce-gin/internals/order/order_repos"
	orderservices "github.com/ak-repo/ecommerce-gin/internals/order/order_services"
	middleware "github.com/ak-repo/ecommerce-gin/middleware/auth"

	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	// admin auth
	authRepo := authrepo.NewAuthRepo(db.DB)
	authService := authservice.NewAuthService(authRepo, cfg)
	authHandler := authhandler.NewAuthHandler(authService)

	r.GET("/login", authHandler.AdminLoginForm)
	r.POST("/login", authHandler.AdminLogin)

	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.AuthMiddleware(cfg), middleware.RoleMiddleware("admin"))

	{

		// dashbord
		dashboardHandler := admindashhandler.NewAdminDashboardHandler()
		adminRoute.GET("/dashboard", dashboardHandler.AdminDashboardShow)

		// products management
		productRepo := adminproductrepo.NewAdminProductRepo(db.DB)
		productService := adminproductservice.NewAdminProductService(productRepo)
		productHandler := adminproducthandler.NewAdminProductHandler(productService)

		adminRoute.GET("/products", productHandler.AdminAllProducstListHandler)
		adminRoute.GET("/product/:id", productHandler.AdminProductListHandler)

		// orders management
		orderRepo := orderrepos.NewOrderRepo(db.DB)
		orderService := orderservices.NewOrderService(orderRepo)
		orderHandler := adminorderhandler.NewAdminOrderHandler(orderService)

		adminRoute.GET("/orders", orderHandler.ShowAllOrderHandler)
		adminRoute.GET("/orders/:id", orderHandler.ShowOrderByIDHandler)

		// user management
		userRepo := adminuserrepo.NewAdminUserRepo(db.DB)
		userService := adminuserservice.NewAdminUserService(userRepo)
		userHandler := adminuserhandler.NewAdminUserHandler(userService)

		adminRoute.GET("/users", userHandler.ListAllUsersHandler)
		adminRoute.GET("/users/:id", userHandler.ListUserByIDHandler)

	}

}
