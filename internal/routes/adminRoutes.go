package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"

	"github.com/ak-repo/ecommerce-gin/internal/common/middleware"
	adminhandler "github.com/ak-repo/ecommerce-gin/internal/handlers/adminHandler"
	orderhandler "github.com/ak-repo/ecommerce-gin/internal/handlers/orderHandler"
	producthandler "github.com/ak-repo/ecommerce-gin/internal/handlers/productHandler"
	userhandler "github.com/ak-repo/ecommerce-gin/internal/handlers/userHandler"
	adminrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/adminRepository"
	orderrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/orderRepository"
	productrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/productRepository"
	userrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/userRepository"
	adminservice "github.com/ak-repo/ecommerce-gin/internal/services/adminService"
	orderservice "github.com/ak-repo/ecommerce-gin/internal/services/orderService"
	productservice "github.com/ak-repo/ecommerce-gin/internal/services/productService"
	userservice "github.com/ak-repo/ecommerce-gin/internal/services/userService"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	adminRoute := r.Group("/admin")
	//auth
	adminRepo := adminrepository.NewAdminRepo(db.DB)
	adminService := adminservice.NewAdminService(adminRepo, cfg)
	adminHandler := adminhandler.NewAdminHandler(adminService)
	adminRoute.GET("/login", adminHandler.AdminLoginFormHandler)
	adminRoute.POST("/login", adminHandler.AdminLoginHandler)

	adminProtected := adminRoute.Group("/")
	adminProtected.Use(middleware.AuthMiddleware(cfg), middleware.RoleMiddleware("admin"))
	{

		//dash bord
		adminProtected.GET("/dashboard", adminHandler.AdminDashboardForm)

		//profile
		adminProtected.GET("/profile", adminHandler.AdminProfileHandler)
		adminProtected.GET("/address/:address_id", adminHandler.ShowAddressFormHandler)
		adminProtected.POST("/address/update/:address_id", adminHandler.UpdateAddressHandler)

		// products
		productRepo := productrepository.NewProductRepo(db.DB)
		productService := productservice.NewProductService(cfg, productRepo)
		productHandler := producthandler.NewAdminproductHandler(productService)

		adminProtected.GET("/products", productHandler.GetAllProductHandler)
		adminProtected.GET("/product/:id", productHandler.GetProductByIDHandler)
		adminProtected.GET("/product/edit/:id", productHandler.ShowProductEdit)
		adminProtected.POST("/product/update/:id", productHandler.UpdateProductHandler)

		// orders
		orderRepo := orderrepository.NewOrderRepo(db.DB)
		orderServcie := orderservice.NewOrderService(orderRepo, cfg)
		orderHandler := orderhandler.NewAdminOrderHandler(orderServcie)

		adminProtected.GET("/orders", orderHandler.ShowAllOrderHandler)
		adminProtected.GET("/order/show/:orderID", orderHandler.ShowOrderByIDHandler)

		//users
		userRepo := userrepository.NewAdminUserRepo(db.DB)
		userService := userservice.NewAdminUserService(userRepo)
		userHandler := userhandler.NewAdminUserHandler(userService)

		adminProtected.GET("/users", userHandler.ListAllUsersHandler)
		adminProtected.GET("/users/:userID", userHandler.ListUserByIDHandler)
	}
}
