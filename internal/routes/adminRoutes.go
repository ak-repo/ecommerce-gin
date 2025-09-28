package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"

	"github.com/ak-repo/ecommerce-gin/internal/common/middleware"
	adminhandler "github.com/ak-repo/ecommerce-gin/internal/handlers/adminHandler"
	producthandler "github.com/ak-repo/ecommerce-gin/internal/handlers/productHandler"
	adminrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/adminRepository"
	productrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/productRepository"
	adminservice "github.com/ak-repo/ecommerce-gin/internal/services/adminService"
	productservice "github.com/ak-repo/ecommerce-gin/internal/services/productService"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	adminRoute := r.Group("/admin")
	//auth
	authRepo := adminrepository.NewAdminAuthRepo(db.DB)
	authService := adminservice.NewAdminAuthService(authRepo, cfg)
	authHandler := adminhandler.NewAdminAuthHandler(authService)
	adminRoute.GET("/login", authHandler.AdminLoginFormHandler)
	adminRoute.POST("/login", authHandler.AdminLoginHandler)
	adminProtected := adminRoute.Group("/")
	adminProtected.Use(middleware.AuthMiddleware(cfg), middleware.RoleMiddleware("admin"))
	{

		//dash bord
		adminProtected.GET("/dashboard", authHandler.AdminDashboardForm)

		// products
		productRepo := productrepository.NewProductRepo(db.DB)
		productService := productservice.NewProductService(cfg, productRepo)
		productHandler := producthandler.NewAdminproductHandler(productService)

		adminProtected.GET("/products", productHandler.GetAllProductHandler)
		adminProtected.GET("/product/:id", productHandler.GetProductByIDHandler)
		adminProtected.GET("/product/edit/:id", productHandler.ShowProductEdit)
		adminProtected.POST("/product/update/:id", productHandler.UpdateProductHandler)
	}
}
